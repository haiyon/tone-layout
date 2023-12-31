package email

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// Mailgun - mailgun config
type Mailgun struct {
	Key    string `json:"key"`
	Domain string `json:"domain"`
	Title  string `json:"title"`
	Email  string `json:"email"`
}

// AuthTemplate - auth template
type AuthTemplate struct {
	Subject  string `json:"subject"`
	Template string `json:"template"`
	Keyword  string `json:"keyword"`
	URL      string `json:"url"`
}

// SendMailgun - use mailgun send mail
func SendMailgun(m *Mailgun, template AuthTemplate) (string, error) {
	domain := m.Domain

	mg := mailgun.NewMailgun(domain, m.Key)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// give time for template to show up in the system.
	time.Sleep(time.Second * 1)

	// Create a new message with template
	from := fmt.Sprintf("%s <no-reply@%s>", m.Title, domain)
	mail := mg.NewMessage(from, template.Subject, "")

	// set template
	mail.SetTemplate(template.Template)
	// recipient
	_ = mail.AddRecipient(m.Email)
	// template params
	_ = mail.AddVariable("keyword", template.Keyword)
	_ = mail.AddVariable("url", template.URL)

	// send
	_, id, err := mg.Send(ctx, mail)
	if err != nil {
		return "", err
	}

	log.Printf("Queued::" + id)
	return id, err
}
