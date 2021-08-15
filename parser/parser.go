package parser

import (
	"bufio"
	"errors"
	"github.com/Azuki-bar/parsender/template"
	"log"
	"net/mail"
	"os"
	"strings"
)

type mailContents struct {
	From    string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}
type filePath string

func isFrom(s string) bool {
	suf := template.Mo.From + template.Mo.ContentsDelimiter
	return strings.Contains(s, suf)
}
func isTo(s string) bool {
	suf := template.Mo.To + template.Mo.ContentsDelimiter
	return strings.Contains(s, suf)
}
func isSubject(s string) bool {
	suf := template.Mo.Subject + template.Mo.ContentsDelimiter
	return strings.Contains(s, suf)
}
func isCc(s string) bool {
	suf := template.Mo.Cc + template.Mo.ContentsDelimiter
	return strings.Contains(s, suf)
}
func isBcc(s string) bool {
	suf := template.Mo.Bcc + template.Mo.ContentsDelimiter
	return strings.Contains(s, suf)
}
func isDelimiter(s string) bool {
	return strings.Contains(s, template.Mo.BodyDelimiter)
}
func (m *mailContents) getFrom(s string) error {
	prf := template.Mo.From + template.Mo.ContentsDelimiter
	if strings.Contains(s, prf) {

		rawAddress := strings.TrimPrefix(s, prf)
		addr, err := mail.ParseAddress(rawAddress)
		if err != nil {
			log.Fatal(err)
		}
		m.From = addr.Address
		return nil
	}
	return errors.New("From address parse error")
}
func (m *mailContents) getTo(s string) error {
	prf := template.Mo.To + template.Mo.ContentsDelimiter
	if strings.Contains(s, prf) {
		rawAddress := strings.TrimPrefix(s, prf)
		addr, err := mail.ParseAddress(rawAddress)
		if err != nil {
			log.Fatal(err)
		}
		m.To = append(m.To, addr.Address)
		return nil
	}
	return errors.New("To address parse error")
}
func (m *mailContents) getCc(s string) error {
	prf := template.Mo.Cc + template.Mo.ContentsDelimiter
	if strings.Contains(s, prf) {

		rawAddress := strings.TrimPrefix(s, prf)
		addr, err := mail.ParseAddress(rawAddress)
		if err != nil {
			log.Fatal(err)
		}
		m.Cc = append(m.Cc, addr.Address)
		return nil
	}
	return errors.New("Cc address parse error")
}
func (m *mailContents) getBcc(s string) error {
	prf := template.Mo.Bcc + template.Mo.ContentsDelimiter
	if strings.Contains(s, prf) {

		rawAddress := strings.TrimPrefix(s, prf)
		addr, err := mail.ParseAddress(rawAddress)
		if err != nil {
			log.Fatal(err)
		}
		m.Bcc = append(m.Bcc, addr.Address)
		return nil
	}
	return errors.New("Cc address parse error")
}
func (m *mailContents) getSubject(s string) error {
	prf := template.Mo.Subject + template.Mo.ContentsDelimiter
	if strings.Contains(s, prf) {
		m.Subject = strings.TrimPrefix(s, prf)
		return nil
	}
	return errors.New("Subject parse error")
}
func (m *mailContents) getBody(s []string) {
	m.Body = strings.Join(s, "\r\n")
}
func (m *mailContents) GetMailText() string {
	exists := struct {
		to  bool
		cc  bool
		bcc bool
	}{
		to:  m.To != nil,
		cc:  m.Cc != nil,
		bcc: m.Bcc != nil,
	}
	if (!exists.to && !exists.cc && !exists.bcc) && (m.From == "" || m.Subject == "" || m.Body == "") {
		log.Fatal(errors.New("mail Text failed. please check your input"))
	}
	var text []string

	text = append(text, template.Mo.From+template.Mo.ContentsDelimiter+m.From)
	if exists.to {
		to := template.Mo.To + template.Mo.ContentsDelimiter
		for _, t := range m.To {
			to = to + t + ","
		}
		text = append(text, to)
	}
	if exists.cc {
		cc := template.Mo.Cc + template.Mo.ContentsDelimiter
		for _, c := range m.Cc {
			cc = cc + c + ","
		}
		text = append(text, cc)
	}
	text = append(text, template.Mo.Subject+template.Mo.ContentsDelimiter+m.Subject)
	text = append(text, "")
	text = append(text, m.Body)
	text = append(text, "")
	return strings.Join(text, "\r\n")
}
func ReadTextFile(f filePath) (*mailContents, error) {
	fp, err := os.Open(string(f))
	if err != nil {
		log.Fatal(err)
	}
	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(fp)
	return ReadText(fp)
}
func ReadText(fp *os.File) (*mailContents, error) {
	mc := mailContents{
		From:    "me",
		To:      nil,
		Cc:      nil,
		Bcc:     nil,
		Subject: "",
		Body:    "",
	}
	scanner := bufio.NewScanner(fp)
	isBody := false
	var body []string
	for scanner.Scan() {
		text := scanner.Text()
		if isBody {
			body = append(body, text)
		} else if isFrom(text) {
			err := mc.getFrom(text)
			if err != nil {
				log.Fatal(err)
			}
		} else if isTo(text) {
			err := mc.getTo(text)
			if err != nil {
				log.Fatal(err)
			}
		} else if isCc(text) {
			err := mc.getCc(text)
			if err != nil {
				log.Fatal(err)
			}
		} else if isBcc(text) {
			err := mc.getBcc(text)
			if err != nil {
				log.Fatal(err)
			}
		} else if isSubject(text) {
			err := mc.getSubject(text)
			if err != nil {
				log.Fatal(err)
			}
		} else if isDelimiter(text) {
			isBody = true
		} else {
			return nil, errors.New("text parse error")
		}
	}
	if isBody {
		mc.getBody(body)
	} else {
		log.Fatal(errors.New("get Body failed"))
	}
	return &mc, nil
}
