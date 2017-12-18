package main

import (
	"fmt"
	"net/smtp"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
notifyForHost sends a notification that a particular host's open ports
do not match the expected ones specified in the configuration file.
*/
func notifyForHost(host string, desiredPorts, openPorts []int) {
	log.Warnf("Host %s port mismatch! Expected: %v. Got: %v.", host, desiredPorts, openPorts)

	if conf.SendEmail == true {
		auth := smtp.PlainAuth("", conf.SMTPUsername, conf.SMTPPassword, conf.SMTPServer)

		for _, mail := range conf.AlertEmail {
			err := smtp.SendMail(
				fmt.Sprintf("%s:%d", conf.SMTPServer, conf.SMTPPort),
				auth,
				conf.FromEmail,
				[]string{mail},
				[]byte(
					fmt.Sprintf(
						"From: PortSpec <%s>\r\nTo:%s\r\nSubject: PortSpec Alert [%s]\r\n\r\nHost: %s\r\nDesired Ports: %v\r\nOpen Ports: %v\r\nTime: %d %s %d %d:%d:%d\r\n",
						conf.FromEmail,
						mail,
						host,
						host,
						desiredPorts,
						openPorts,
						time.Now().Day(),
						time.Now().Month().String()[0:3],
						time.Now().Year(),
						time.Now().Hour(),
						time.Now().Minute(),
						time.Now().Second(),
					),
				),
			)

			if err != nil {
				log.Errorf("Failed to send e-mail to %s about host %s: %s", mail, host, err.Error())
			}
		}

	}

}
