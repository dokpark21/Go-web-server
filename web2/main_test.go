package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	assert := assert.New(t)
	path := "/Users/parkssanghyeon/Downloads/아이묭묭.jpeg"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.Nil(err)

	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads" + filepath.Base(path)
	_, err := os.Stat()
	assert.NoError(err)

	uploadFileTest, err := os.Open(uploadFilePath)
	originFileTest, err := os.Open(path)
	defer uploadFileTest.Close()
	defer originFileTest.Close()

	uploadFile := []byte{}
	originFile := []byte{}
	uploadFile.Read(uploadFile)
	originFile.Read(originFile)

	assert.Equal(originFile, uploadFile)
}
