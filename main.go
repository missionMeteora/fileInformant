package main

import (
	"log"

	"github.com/missionMeteora/mandrill"
	"github.com/missionMeteora/twilio"

	"github.com/missionMeteora/fileInformant/internal/config"
	"github.com/missionMeteora/fileInformant/internal/file"
)

func main() {
	cfg, err := config.New("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	ec := getEmailClient(cfg.ApiInfo.Mandrill)
	tc := getTwilioClient(cfg.ApiInfo.Twilio)

	for _, f := range cfg.Files {
		file.New(cfg.Name, f.Location, f.Interval, cfg.Subscribers, ec, tc)
	}

	select {}
}

func getEmailClient(m config.Mandrill) *mandrill.Client {
	return mandrill.New(m.Key, m.SubAccount, m.FromEmail, m.FromName)
}

func getTwilioClient(t config.Twilio) *twilio.Client {
	return twilio.New(t.Key, t.Token, t.FromPhone)
}
