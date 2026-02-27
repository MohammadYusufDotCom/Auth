package mailer

import (
	"bytes"
	"embed"
	"html/template"
)

type OTPData struct {
	OTP  string
	NAME string
}

var templateFS embed.FS

func ParseOTPTemplate(otp string, name *string) (string, error) {

	tmpl, err := template.ParseFiles("internal/mailer/templates/otp_email.html")
	if err != nil {
		return "", err
	}

	// data := struct {
	// 	OTP string
	//  NAME string
	// }{
	// 	OTP: otp,
	// 	NAME: *name,
	// }
	data := OTPData{
		OTP:  otp,
		NAME: *name,
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return "", err
	}

	// fmt.Println("template: ", body.String())
	return body.String(), nil
}
