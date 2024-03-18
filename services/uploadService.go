package services

import (
	"errors"

	// "fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type UploadFileService interface {
	ReceiptFileUpload(writer http.ResponseWriter, req *http.Request) error
}

type service struct {
	s string
}

func NewService(s string) UploadFileService {
	return &service{s: s}
}

const (
	APIKeyHeader = "X-API-Key"
	ValidAPIKey  = "N6RJJuPaOMI1hAwUHeSw"
)

func (s *service) ReceiptFileUpload(writer http.ResponseWriter, req *http.Request) error {
	if req.Method != "POST" {
		return errors.New("เฉพาะเมธอด POST เท่านั้นที่ได้รับอนุญาต")
	}
	apiKey := req.Header.Get(APIKeyHeader)
	if apiKey != ValidAPIKey {
		return errors.New("InvalidAPIKey")
	}

	file, header, err := req.FormFile("fileUpload")
	if err != nil {
		return errors.New("ไม่สามารถรับไฟล์ได้: " + err.Error())
	}
	defer file.Close()

	//create file
	dateStr := time.Now().Format(string("20060102"))

	f, err := os.Create("./uploaded/" + "uploaded_" + header.Filename + dateStr)
	if err != nil {
		return errors.New("ไม่สามารถสร้างไฟล์ได้: " + err.Error())
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return errors.New("ไม่สามารถเขียนข้อมูลลงในไฟล์ได้: " + err.Error())
	}

	err = s.sendEmailSetup()
	if err != nil {
		return errors.New("ไม่สามารถสส่งเมลล์ได้: " + err.Error())
	}

	return nil
}
