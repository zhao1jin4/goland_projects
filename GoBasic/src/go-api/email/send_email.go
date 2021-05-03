package main

import (
	"flag"
	"fmt"
	"net/smtp"
	"strings"
)

//go run src/adv/email/send_email.go --from=lisi@sina.com --frompass=123 --fromhost=smtp.sina.com:25 --mailto=wang@163.com
func main() {

	from := flag.String("from", "", "邮件用户名")
	frompass := flag.String("frompass", "", "邮件密码")
	fromhost := flag.String("fromhost", "", "SMTP邮件服务器")
	mailto := flag.String("mailto", "", "收件人")
	flag.Parse()

	fmt.Println("from=", *from, "frompass=", *frompass, "fromhost=", *fromhost, "to=", *mailto)
	if *from == "" || *fromhost == "" || *mailto == "" {
		fmt.Println("from 或 fromhost 或 mailto 为空")
		return
	}

	user := *from
	password := *frompass
	host := *fromhost
	to := *mailto

	subject := "使用Golang发送邮件"
	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
