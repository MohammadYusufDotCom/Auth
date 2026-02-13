package mailer

import (
	"auth/internal/config"
	"fmt"
	"net/smtp"
)

const subject = "Your One-Time Password (OTP) for YUSUF"

// const body = "Your 6 digit verification code is: "

var Cnfig = config.LoadConfig()

func SendMail(to string, otp string, name string) error {

	from := Cnfig.Mail.User
	password := Cnfig.Mail.Password
	host := Cnfig.Mail.Host
	port := Cnfig.Mail.Port

	template, err := ParseOTPTemplate(otp, &name)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}

	// message := []byte(
	// 	"Subject: " + subject + "\r\n" +
	// 		"\r\n" +
	// 		body + otp + "\r\n" +
	// 		"It will expire in 5 minutes.\r\n",
	// )

	message := []byte(
		"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
			template,
	)

	auth := smtp.PlainAuth(
		"",
		from,
		password,
		Cnfig.Mail.Host,
	)

	addr := fmt.Sprintf("%s:%d", host, port)

	mailErr := smtp.SendMail(
		addr,
		auth,
		from,
		[]string{to},
		message,
	)

	if mailErr != nil {
		fmt.Println("Error sending email:", mailErr)
		return mailErr
	}

	fmt.Println("OTP Email sent successfully!")
	return nil
}
