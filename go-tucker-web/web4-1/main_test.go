package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadTest(t *testing.T) {
	// given
	assert := assert.New(t)
	path := "/Users/yun/Downloads/converted_image.png"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	_, err = io.Copy(multi, file)
	writer.Close()

	// when
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/uploads", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// then
	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath) // return fileInfo, 파일이 있으면 반환 없으면 에러가 나기 때문에 파일이 있는지 없는지 이걸로 체크
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(uploadData, originData)
}
