package myapp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIndexPathHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	indexHandler(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Hello World", string(data))
}

func TestBarPathHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Hello World!", string(data))
}

func TestBarPathHandlerWithName(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/bar?name=kakao", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Hello kakao!", string(data))
}

func TestFooHandlerWithoutJson(t *testing.T) {
	asst := assert.New(t)

	s := struct {
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}{
		FirstName: "kakao",
		LastName:  "daum",
		Email:     "hell@daum.net",
		CreatedAt: time.Now(),
	}
	jsonString, _ := json.Marshal(s)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/foo", strings.NewReader(string(jsonString)))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	asst.Equal(http.StatusOK, res.Code)

	u := &User{}
	err := json.NewDecoder(res.Body).Decode(u)
	asst.Nil(err)
	asst.Equal("kakao", u.FirstName)
	asst.Equal("daum", u.LastName)
	asst.Equal("hell@daum.net", u.Email)

}
