package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiptFileUpload(t *testing.T) {
	// Arrange

	const (
		sAPIKeyHeader = "X-API-Key"
		sValidAPIKey  = "N6RJJuPaOMI1hAwUHeSw"
	)

	w := httptest.NewRecorder()
	fmt.Println(w)
	fileContent := []byte("fileUpload")
	file := bytes.NewReader(fileContent)

	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("fileUpload", "test.txt")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-API-Key", sValidAPIKey)

	httpReq, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	httpReq.Header = req.Header

	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("UploadFile_Success", func(t *testing.T) {
		s := NewService("example string")
		err = s.ReceiptFileUpload(w, httpReq)

		if err == nil {
			t.Error("Expected an error but got nil")
		}
		//Assert
		assert.NotNil(t, err)

	})
	t.Run("UploadCheckAPIKEY_Success", func(t *testing.T) {

		//Act
		apiKey := req.Header.Get(APIKeyHeader)
		//Assert
		assert.Equal(t, apiKey, sValidAPIKey)
	})

	t.Run("UploadCheckAPIKEY_Fail", func(t *testing.T) {
		//Arrange
		const (
			sAPIKeyHeader = "X-API-Key"
			sValidAPIKey  = "N6RJJuPaOMI1hAwdddddUHeSw"
		)

		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("X-API-Key", ValidAPIKey)
		//Act
		apiKey := req.Header.Get(APIKeyHeader)
		//Assert
		assert.NotEqual(t, apiKey, sValidAPIKey)
	})
	writer.Close()
}

func Tfff(t *testing.T) {

}
