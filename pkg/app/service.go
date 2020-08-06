package app

import (
	"fmt"
	"github.com/areknoster/table-driven-tests-gomock/pkg/names"
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
)

// namesLister is a facade which let's us mock some external dependency's behaviour
type namesLister interface {
	ListNames(who names.Group) ([]string, error)
}

type PartyService struct {
	namesLister namesLister
	helloer     party.Helloer
}

func NewService(service *names.Service, helloer party.Helloer) *PartyService {
	return &PartyService{
		namesLister: service,
		helloer:     helloer,
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
		fmt.Println(s.helloer.Hello(visitorName))
	}
	return nil
}
