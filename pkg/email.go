package pkg

import (
	"crypto/tls"
	"github.com/arioprima/jobseekers_api/config"
	"gopkg.in/gomail.v2"
	"math/rand"
)

func GenerateOtp() string {
	number := "0123456789"

	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = number[rand.Intn(len(number))]
	}

	return string(otp)
}

func SendEmail(email, otp string) {
	configs, _ := config.LoadConfig(".")

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
	mailer.SetBody("text/html", "Your OTP is "+otp)

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
