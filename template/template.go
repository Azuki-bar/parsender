package template

import (
	"bytes"
	"html/template"
	"log"
)

type mailTemplate string

type mailOption struct {
	From              string
	ReplyTo           string
	To                string
	Cc                string
	Bcc               string
	Subject           string
	ContentsDelimiter string
	BodyDelimiter     string
}

var Mo = mailOption{
	From:              "From",
	ReplyTo:           "reply-to",
	To:                "To",
	Cc:                "Cc",
	Bcc:               "Bcc",
	Subject:           "Subject",
	ContentsDelimiter: ": ",
	BodyDelimiter:     "==write body below==",
}

func GetMailTemplate() mailTemplate {
	t := `
{{.From}}{{.ContentsDelimiter}}
{{.To}}{{.ContentsDelimiter}}
{{.Cc}}{{.ContentsDelimiter}}
{{.Bcc}}{{.ContentsDelimiter}}
{{.Subject}}{{.ContentsDelimiter}}
{{.BodyDelimiter}}

`
	tpl, err := template.New("mailString").Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, Mo)
	if err != nil {
		log.Fatal(err)
	}
	return mailTemplate(buf.String())
}
