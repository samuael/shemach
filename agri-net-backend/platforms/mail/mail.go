package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	gomail "gopkg.in/mail.v2"
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
    <h1 style="background-color:#006699";color:#fff >  Agri-Net  </h1>
    {{if .NewPassword}}
		<p> Hi, According to your action in the Agri-info Systems this email is used to register a User named {{.Fullname }} in the system .<br>
		we have sent you a password to be used when login with <br>
		</p>
	{{else }}
	 <p> Hi {{.Fullname }}  , According Your Request of forgot password this email is sent to you.</p>
	{{end}}
	<p>  Your New Password is <b>[{{.Password }}]</b> </p>
    <p> <i> Use this password to log in and change the password ASAP.</i> </p>
	<p>  If this password doesn't work try to get a new password again using forgot password or Use the previous password if you still remember it.</p>
	<hr>
	<i> Agri-Info Systems <small>  Agri-Net ( Agricultural Products commecrial netword systems ) </small> </i> 
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
    <h1 style="background-color:#006699";color:#fff > Agri-net Systems </h1>
		<p> Hi {{.Fullname }}, According to your action in agri-net systems, this email is sent to confirm the email address.<br>
		<p><b> Please Make Sure You confirm your email with is 3- Minutes.</b></p>
		To Confirm your Email, click the link below
		<a href="{{.HOST}}/api/account/confirm/?token={{.TOKEN}}" > Not Mine </a>
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

func SendEmail(from, to, password string) bool {
	templ, er := template.New("forgot").Parse(dpassword)
	println(password)
	if er != nil || templ == nil {
		return false
	}
	file, er := os.Create("forgot-password.html")
	if er != nil || file == nil {
		return false
	}
	if era := templ.Execute(file, password); era != nil {
		return false
	}
	bytes := []byte{}
	if _, er = file.Read(bytes); er != nil {
		return false
	}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "New Password Update!")
	m.SetBody("text/html", string(bytes))
	d := gomail.NewDialer("smtp.gmail.com", 587, from, "0774samuael")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		println(err.Error())
		return false
	}
	return true
}

func SendEmailChangeSMTP(to []string, password, fullname, host string) bool {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := os.Getenv("EMAIL_ADDRESS")
	auth := smtp.PlainAuth("", from, os.Getenv("EMAIL_PASSWORD"), smtpHost)
	t, _ := template.New("email-update").Parse(emailupdate)
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := " Email account update "
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, struct {
		Email    string
		HOST     string
		Fullname string
		Password string
	}{
		Fullname: fullname,
		Email:    to[0],
		HOST:     state.HOST,
		Password: password,
	})
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func ConfirmUpdateEmailAccount(to []string, token, fullname, host string) bool {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := os.Getenv("EMAIL_ADDRESS")
	auth := smtp.PlainAuth("", from, os.Getenv("EMAIL_PASSWORD"), smtpHost)
	t, _ := template.New("email-update").Parse(emailupdate)
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := " Email account Confirmation "
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, struct {
		Email    string
		HOST     string
		Fullname string
		Password string
	}{
		Fullname: fullname,
		Email:    to[0],
		HOST:     state.HOST,
		Password: token,
	})
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
