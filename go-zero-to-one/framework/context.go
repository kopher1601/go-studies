package framework

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
)

type MyContext struct {
	w      http.ResponseWriter
	r      *http.Request
	params map[string]string
}

func NewMyContext(w http.ResponseWriter, r *http.Request) *MyContext {
	return &MyContext{
		w:      w,
		r:      r,
		params: make(map[string]string),
	}
}

func (ctx *MyContext) Json(data any) {
	responseData, err := json.Marshal(data)
	if err != nil {
		ctx.w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.w.Header().Set("Content-Type", "application/json")
	ctx.w.WriteHeader(http.StatusOK)
	ctx.w.Write(responseData)
}

func (ctx *MyContext) WriteString(data string) {
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
