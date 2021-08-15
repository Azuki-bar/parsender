package main

import (
	"fmt"
	"github.com/Azuki-bar/parsender/parser"
	"github.com/Azuki-bar/parsender/setting"
	"github.com/Azuki-bar/parsender/template"
	"log"
	"net/smtp"
)

func main() {
	fmt.Printf(string(template.GetMailTemplate()))
	//m, err := parser.ReadText(os.Stdin)
	m, err := parser.ReadTextFile("./testMessage.txt")
	if err != nil {
		log.Fatal(err)
	}
	txt := m.GetMailText()
	smtpConf, err := setting.GetSmtpConf("")
	if err != nil {
		log.Fatal(err)
	}
	auth := smtp.PlainAuth(smtpConf.Auth.Identity, smtpConf.Auth.Username, smtpConf.Auth.Password, smtpConf.Auth.Host)
	var receivers []string
	receivers = append(append(append(receivers, m.Bcc...), m.Cc...), m.To...)

	err = smtp.SendMail(smtpConf.Server.Address, auth, smtpConf.Server.From, receivers, []byte(txt))
	if err != nil {
		log.Fatal(err)
	}
}
