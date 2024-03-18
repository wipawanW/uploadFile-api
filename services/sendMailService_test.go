package services

import (
	"net/smtp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMailNotiFileUpload(t *testing.T) {

	t.Run("SendEmail_Success", func(t *testing.T) {
		//Arrange
		smtpHost := "smtp.gmail.com"
		smtpPort := "587"
		sender := "wipawan.won@gmail.com"
		password := "fyntdjxzifeawrhv"
		auth := smtp.PlainAuth("", sender, password, smtpHost)

		to := []string{"somder303@gmail.com"}
		subject := "Notification: File Upload Complete"
		body := "Dear [Every One],\n\n" +
			"The system would like to inform you that the file upload process has been successfully completed. \n\n" +
			"Your files have been received and processed accordingly.\n\n" +
			"Should you have any questions or require further assistance, please do not hesitate to contact us. \n\n" +

			"We are here to help. \n\n" +

			"Thank you for choosing our service.,\n[Wipawan Wonganan]\n"
		// สร้างข้อความอีเมล
		msg := []byte("To: " + to[0] + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n")

		//Act
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, to, msg)

		//Assert
		assert.Nil(t, err)
	})

}
