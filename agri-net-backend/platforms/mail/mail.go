package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"os"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
)

var dpassword = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> {{ if .NewPassword }}  
		New Password!
		{{else}}
	Forgot Password {{end}}</title>
</head>
<body>
    <h1 style="background-color:#00FF00";color:#fff >  Agri-Net  </h1>
    {{if .NewPassword}}
		<p> Hi {{.Fullname}}, this is your new password for Agri-Net systems <br>
		</p>
	{{else }}
	 <p> Hi {{.Fullname }}  , According Your Request of forgot password this email is sent to you.</p>
	{{end}}
	<p>  Your New Password is <b>[{{.Password }}]</b> </p>
    <p> <i> Use this password to log in, and change the password with is 30 minutes.</i> </p>
	<p> <i> If this password doesn't work try to get a new password again using forgot password or Use the previous password if you still remember it.</i></p>
	<hr>
	<i> Agri-Net Systems <small>  Agri-Net ( Agricultural Products commercial network ) </small></i> 
	<hr>
	</body>
</html>`

var emailupdate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> Email Account Confirmation </title>
</head>
<body>
    <h1 style="background-color:#00FF00";color:#fff > Agri-net Systems </h1>
		<p> Hi {{.Fullname }}, According to your action in agri-net systems, this email is sent to confirm the email address.<br>
		<p><b> Please Make Sure You confirm your email with is 3- Minutes.</b></p>
		To Confirm your Email, click the link below
		<a href="http://{{.HOST}}/api/user/account/email/confirm?token={{.TOKEN}}" > HERE </a>
		</p>
	<hr>
	<i> Agri-Net Systems </i>
	<small> We Provide a reliable agricultural products exchange method </small>
	<hr>
	</body>
</html>`

func SendPasswordEmailSMTP(to []string, password string, newpassword bool, fullname, host string) bool {
	println("The Password is : ", password)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := os.Getenv("EMAIL_ADDRESS")
	auth := smtp.PlainAuth("", from, os.Getenv("EMAIL_PASSWORD"), smtpHost)
	t, _ := template.New("forgot-password").Parse(dpassword)
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	subject := " New Account Password "
	if !newpassword {
		subject = " Update Forgotten Password Request "
	}
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, struct {
		NewPassword bool
		Password    string
		Email       string
		HOST        string
		Fullname    string
	}{
		Fullname:    fullname,
		Email:       to[0],
		HOST:        state.HOST,
		NewPassword: newpassword,
		Password:    password,
	})
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Email Sent!")
	return true
}

func ConfirmUpdateEmailAccount(response http.ResponseWriter, to []string, token, fullname, host string) bool {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := os.Getenv("EMAIL_ADDRESS")
	auth := smtp.PlainAuth("", from, os.Getenv("EMAIL_PASSWORD"), smtpHost)
	t, _ := template.New("email-update").Parse(emailupdate)
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := " Email account Confirmation "
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	if host == "" {
		host = state.HOST
	}
	t.Execute(&body, struct {
		Email    string
		HOST     string
		Fullname string
		TOKEN    string
		PORT     string
	}{
		Fullname: fullname,
		Email:    to[0],
		HOST:     host + ":8080",
		TOKEN:    token,
	})
	response.Header().Set("Authorization", "Bearer "+token)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
