package framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/textproto"
	"sync"
)

type MyContext struct {
	w          http.ResponseWriter
	r          *http.Request
	params     map[string]string
	keys       map[string]any
	mux        sync.RWMutex
	hasTimeout bool
}

func NewMyContext(w http.ResponseWriter, r *http.Request) *MyContext {
	return &MyContext{
		w:      w,
		r:      r,
		params: make(map[string]string),
		mux:    sync.RWMutex{},
	}
}

func (ctx *MyContext) Get(key string, defaultValue any) any {
	ctx.mux.RLock()
	defer ctx.mux.RUnlock()
	if ctx.keys == nil {
		return defaultValue
	}

	if res, ok := ctx.keys[key]; ok {
		return res
	}
	return defaultValue
}

func (ctx *MyContext) Set(key string, value any) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	if ctx.keys == nil {
		ctx.keys = make(map[string]any)
	}
	ctx.keys[key] = value
}

func (ctx *MyContext) SetHasTimeout(timeout bool) {
	ctx.hasTimeout = timeout
}
func (ctx *MyContext) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *MyContext) BindJson(data any) error {
	byteData, err := io.ReadAll(ctx.r.Body)
	if err != nil {
		return err
	}

	ctx.r.Body = io.NopCloser(bytes.NewBuffer(byteData))

	return json.Unmarshal(byteData, data)
}

func (ctx *MyContext) Json(data any) {
	if ctx.hasTimeout {
		return
	}

	responseData, err := json.Marshal(data)
	if err != nil {
		ctx.w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.w.Header().Set("Content-Type", "application/json")
	ctx.w.WriteHeader(http.StatusOK)
	ctx.w.Write(responseData)
}

func (ctx *MyContext) JsonP(callback string, parameter any) error {
	if ctx.hasTimeout {
		return nil
	}
	ctx.w.Header().Add("Content-Type", "application/javascript")
	callback = template.JSEscapeString(callback)
	_, err := ctx.w.Write([]byte(callback))
	_, err = ctx.w.Write([]byte("("))
	parameterData, err := json.Marshal(parameter)
	if err != nil {
		return err
	}
	_, err = ctx.w.Write(parameterData)
	_, err = ctx.w.Write([]byte(")"))
	if err != nil {
		return err
	}

	return nil
}

func (ctx *MyContext) WriteString(data string) {
	if ctx.hasTimeout {
		return
	}
	ctx.w.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.w, data)
}

func (ctx *MyContext) QueryAll() map[string][]string {
	return ctx.r.URL.Query()
}

func (ctx *MyContext) QueryKey(key string, defaultValue string) string {
	values := ctx.QueryAll()
	if target, ok := values[key]; ok {
		if len(target) == 0 {
			return defaultValue
		}
		return target[len(target)-1]
	}
	return defaultValue
}

func (ctx *MyContext) SetParams(dicts map[string]string) {
	ctx.params = dicts
}

func (ctx *MyContext) GetParam(key string, defaultValue string) string {
	params := ctx.params

	if v, ok := params[key]; ok {
		return v
	}
	return defaultValue
}

func (ctx *MyContext) FormKey(key, defaultName string) string {
	if ctx.r.Form == nil {
		ctx.r.ParseMultipartForm(32 << 20)
	}
	if vs := ctx.r.Form[key]; len(vs) > 0 {
		return vs[0]
	}
	return defaultName
}

type FormFileInfo struct {
	Data     []byte
	Filename string
	Header   textproto.MIMEHeader
	Size     int64
}

func (ctx *MyContext) FormFile(key string) (*FormFileInfo, error) {
	file, fileHeader, err := ctx.r.FormFile(key)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &FormFileInfo{
		Data:     data,
		Filename: fileHeader.Filename,
		Header:   fileHeader.Header,
		Size:     fileHeader.Size,
	}, nil
}

func (ctx *MyContext) WriteHeader(httpStatusCode int) {
	ctx.w.WriteHeader(httpStatusCode)
}

func (ctx *MyContext) RenderHtml(filepath string, data any) error {
	if ctx.hasTimeout {
		return nil
	}
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}

	return t.Execute(ctx.w, data)
}
