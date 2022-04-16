package email

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

type EmailSender struct {
	UserName string
	Password string
	Addr     string
	Name     string
}
type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

func (sender EmailSender) SendEmail(toAddress string, subject string, body string, use_plain bool) {
	from := mail.Address{Name: sender.Name, Address: sender.UserName}
	to := mail.Address{Name: "", Address: toAddress}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["MIME-version"] = ": 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := sender.Addr

	host, _, _ := net.SplitHostPort(servername)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	var auth smtp.Auth
	var c *smtp.Client
	if use_plain {
		auth = smtp.PlainAuth("", sender.UserName, sender.Password, host)

		conn, err := tls.Dial("tcp", servername, tlsconfig)
		if err != nil {
			panic(err)
		}

		c, err = smtp.NewClient(conn, host)
		if err != nil {
			panic(err)
		}
	} else {
		auth = LoginAuth(sender.UserName, sender.Password)

		conn, err := net.Dial("tcp", servername)
		if err != nil {
			panic(err)
		}

		c, err = smtp.NewClient(conn, host)
		if err != nil {
			panic(err)
		}
	}

	if err := c.StartTLS(tlsconfig); err != nil {
		panic(err)
	}

	// Auth
	if err := c.Auth(auth); err != nil {
		panic(err)
	}

	// To && From
	if err := c.Mail(from.Address); err != nil {
		panic(err)
	}

	if err := c.Rcpt(to.Address); err != nil {
		panic(err)
	}

	// Data
	if w, err := c.Data(); err != nil {
		panic(err)
	} else {
		if _, err = w.Write([]byte(message)); err != nil {
			panic(err)
		}

		if err = w.Close(); err != nil {
			panic(err)
		}
	}

	c.Quit()
}
