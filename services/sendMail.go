package services

import (
	"net/smtp"
)

type SendMailService interface {
	sendEmailSetup() error
	// sendEmail(smtpHost, smtpPort, sender, password string, to []string, subject, body string) error
}

type Mailservice struct {
}

func NewMailService() UploadFileService {
	return &service{}
}

func (s *service) sendEmailSetup() error {

	// เซ็ตข้อมูลการเชื่อมต่อ SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	sender := "wipawan.won@gmail.com"
	password := "fyntdjxzifeawrhv"

	// ข้อความอีเมล
	to := []string{"somder303@gmail.com"}
	subject := "Notification: File Upload Complete"
	body := "Dear [Every One],\n\n" +
		"The system would like to inform you that the file upload process has been successfully completed. \n\n" +
		"Your files have been received and processed accordingly.\n\n" +
		"Should you have any questions or require further assistance, please do not hesitate to contact us. \n\n" +

		"We are here to help. \n\n" +

		"Thank you for choosing our service.,\n[Wipawan Wonganan]\n"

	// กำหนดการเชื่อมต่อ SMTP ผ่าน TLS
	auth := smtp.PlainAuth("", sender, password, smtpHost)

	// สร้างข้อความอีเมล
	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// ส่งอีเมลผ่านการเชื่อมต่อ TLS
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, to, msg)
	if err != nil {
		return err
	}
	return nil
}

// func (s *service) sendEmail(smtpHost, smtpPort, sender, password string, to []string, subject, body string) error {
// 	// กำหนดการเชื่อมต่อ SMTP
// 	auth := smtp.PlainAuth("", sender, password, smtpHost)

// 	// สร้างข้อความอีเมล
// 	msg := []byte("To: " + to[0] + "\r\n" +
// 		"Subject: " + subject + "\r\n" +
// 		"\r\n" +
// 		body + "\r\n")

// 	// ส่งอีเมล
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, to, msg)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
