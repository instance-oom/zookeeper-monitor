package common

import (
	"bytes"
	"net/http"

	"github.com/astaxie/beego"
)

//SendMail : Send email
func SendMail(body string, to string, serverIP string) {
	client := &http.Client{}
	from := beego.AppConfig.String("mail.from")
	subject := beego.AppConfig.String("mail.subject")
	msg := `{
				"From":"` + from + `",
				"To":"` + to + `",
				"Subject":"` + subject + " - " + serverIP + `",
				"Body":"` + body + `"
				"ContentType":"Html",
				"MailType":"Smtp",
				"SmtpSetting":{}
			}`
	client.Post("http://apis.newegg.org/framework/v1/mail", "application/json", bytes.NewBuffer([]byte(msg)))
}
