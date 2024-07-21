package pkg

import (
	"crypto/tls"
	"fmt"
	"github.com/arioprima/jobseekers_api/config"
	"github.com/arioprima/jobseekers_api/helpers"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

var logger = config.NewLogger()

func GenerateOtp() string {
	number := "0123456789"

	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = number[rand.Intn(len(number))]
	}

	return string(otp)
}

func SendEmail(name, email, otp, fileName string) {
	configs, _ := config.LoadConfig(".")
	date := time.Now().Format("2 Jan, 2006")
	year := time.Now().Year()

	//send email
	smtpEmail := configs.Email.SMTPEmail
	to := email
	smtpPass := configs.Email.SMTPPass
	smtpHost := configs.Email.SMTPHost
	smtpSender := configs.Email.SMTPSender
	smtpPort := configs.Email.SMTPPort

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", smtpSender)
	mailer.SetHeader("To", to)
	mailer.SetAddressHeader("Cc", smtpEmail, smtpSender)
	mailer.SetHeaders(map[string][]string{
		"From":    {mailer.FormatAddress(smtpEmail, smtpSender)},
		"To":      {to},
		"Subject": {"OTP Verification"},
	})
	mailer.SetBody("text/html", helpers.ParseHtml(fileName, map[string]string{
		"to":   name,
		"otp":  otp,
		"date": date,
		"year": fmt.Sprintf("%d", year),
	}))
	//
	//logger.WithFields(logrus.Fields{
	//	"smtp_email":  smtpEmail,
	//	"smtp_host":   smtpHost,
	//	"smtp_sender": smtpSender,
	//	"smtp_port":   smtpPort,
	//	"smtp_pass":   smtpPass,
	//	"to":          to,
	//}).Debug("Preparing to send email")

	dialer := gomail.NewDialer(
		smtpHost,
		smtpPort,
		smtpEmail,
		smtpPass,
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(mailer); err != nil {
		panic(err)
	}

	return
}
