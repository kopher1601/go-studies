package app

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo by")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "User ID :89")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	dto := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}{
		FirstName: "Kakao",
		LastName:  "Daum",
		Email:     "line@line.com",
	}
	jsonBytes, _ := json.Marshal(dto)
	resp, err := http.Post(
		ts.URL+"/users",
		"application/json",
		strings.NewReader(string(jsonBytes)),
	)
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	u := &User{}
	err = json.NewDecoder(resp.Body).Decode(u)
	assert.NoError(err)
	assert.NotEqual(0, u.ID)

	id := u.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	u2 := &User{}
	err = json.NewDecoder(resp.Body).Decode(u2)
	assert.NoError(err)
	assert.Equal(id, u2.ID)
	assert.Equal(u.FirstName, u2.FirstName)
}
