package mailer

import (
	"bytes"
	"crypto/tls"
	"github.com/lenistwo/model"
	"github.com/lenistwo/util"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"strconv"
)

const (
	mailTitle = "Sale Alert!"
)

var dialer *gomail.Dialer

type TemplateData struct {
	Name  string
	Time  string
	Price float64
}

func Send(promotion model.Promotion) {
	message := gomail.NewMessage()
	message.SetHeader("To", os.Getenv("RECEIVER_MAIL"))
	message.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	message.SetHeader("Subject", mailTitle)
	message.SetBody("text/html", buildTemplate(promotion))
	util.CheckError(dialer.DialAndSend(message))
}

func Setup() {
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	util.CheckError(err)
	dial := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))
	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	dialer = dial
}

func buildTemplate(promotion model.Promotion) string {
	t, err := template.ParseFiles("template/template.html")
	util.CheckError(err)
	var buffer bytes.Buffer
	util.CheckError(t.Execute(&buffer, TemplateData{
		Name:  promotion.PromotionName,
		Time:  promotion.PromotionEnd,
		Price: promotion.Price,
	}))
	return buffer.String()
}
