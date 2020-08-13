package app

import (
	"fmt"
	"github.com/areknoster/table-driven-tests-gomock/pkg/names"
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
)

// namesLister is an interface let's us mock names.Service functions
type namesLister interface {
	ListNames(who names.Group) ([]string, error)
}

type PartyService struct {
	namesLister namesLister
	greeter     party.Greeter
}

func NewPartyService(namesService *names.Service, greeter party.Greeter) *PartyService {
	return &PartyService{
		namesLister: namesService,
		greeter:     greeter,
	}
}

func (s *PartyService) GreetVisitors(justNice bool) error {
	visitorNames, err := s.namesLister.ListNames(names.Nice)
	if err != nil {
		return fmt.Errorf("could get nice people names: %w", err)
	}
	if !justNice {
		notNice, err := s.namesLister.ListNames(names.NotNice)
		if err != nil {
			return fmt.Errorf("could not get not-nice people's names' ")
		}
		visitorNames = append(visitorNames, notNice...)
	}
	for _, visitorName := range visitorNames {
		fmt.Println(s.greeter.Hello(visitorName))
	}
	return nil
}
