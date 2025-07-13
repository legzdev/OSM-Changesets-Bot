package main

import (
	"errors"
	"time"

	"github.com/legzdev/OSM-Changesets-Bot/database"
	"github.com/legzdev/OSM-Changesets-Bot/env"
	"github.com/legzdev/OSM-Changesets-Bot/internal"

	"github.com/charmbracelet/log"
)

func main() {
	log.Info("bot started!")

	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Init()
	if err != nil {
		log.Fatal("db init:", err)
	}

	for {
		latest, err := db.GetLatest()
		if err != nil {
			if !errors.Is(err, database.ErrNotFound) {
				log.Error("error getting latest changeset", "err", err)
			}
		}

		changesets, err := internal.NewChangesets(latest)
		if err != nil {
			log.Error("error getting new changesets", "err", err)
		}

		for _, changeset := range changesets {
			err := internal.SendToTelegram(changeset)
			if err != nil {
				log.Error("error sending changeset to telegram", "id", changeset.ID, "err", err)
				break
			}

			err = db.SetLatest(changeset.ID)
			if err != nil {
				log.Error("error setting latest changeset", "id", changeset.ID, "err", err)
				break
			}

			// avoid flood
			time.Sleep(5 * time.Second)
		}

		time.Sleep(env.ChecksInterval)
	}
}
