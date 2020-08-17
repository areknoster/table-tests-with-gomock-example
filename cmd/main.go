package main

import (
	"github.com/areknoster/table-driven-tests-gomock/pkg/app"
	"github.com/areknoster/table-driven-tests-gomock/pkg/people"
	"log"
)

func main() {
	partyService := app.NewPartyService(&people.Organizer{}, people.Host{})
	if err := partyService.GreetVisitors(true); err != nil {
		log.Fatalf("could not greet visitors: %w", err)
	}
}
