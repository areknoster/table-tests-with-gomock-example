package app

import (
	"fmt"
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
)

type PartyService struct {
	visitorLister party.VisitorLister
	greeter       party.Greeter
}

func NewPartyService(namesService party.VisitorLister, greeter party.Greeter) *PartyService {
	return &PartyService{
		visitorLister: namesService,
		greeter:       greeter,
	}
}

func (s *PartyService) GreetVisitors(justNice bool) error {
	visitors, err := s.visitorLister.ListVisitors(party.NiceVisitor)
	if err != nil {
		return fmt.Errorf("could get nice people names: %w", err)
	}
	if !justNice {
		notNice, err := s.visitorLister.ListVisitors(party.NotNiceVisitor)
		if err != nil {
			return fmt.Errorf("could not get not-nice people's names' ")
		}
		visitors = append(visitors, notNice...)
	}
	for _, visitor := range visitors {
		fmt.Println(s.greeter.Hello(visitor.String()))
	}
	return nil
}
