package main

import (
	"log"

	"github.com/missionMeteora/fileInformant/internal/config"
	"github.com/missionMeteora/fileInformant/internal/file"
	"github.com/missionMeteora/fileInformant/internal/notifiers"
)

func main() {
	cfg, err := config.New("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	n := notifiers.New(cfg.Notifiers)

	for _, f := range cfg.Files {
		file.New(cfg.Name, f.Location, f.Interval, cfg.Subscribers, n)
	}

	select {}
}
