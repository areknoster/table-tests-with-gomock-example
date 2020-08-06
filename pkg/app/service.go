package app

import (
	"fmt"
	"github.com/areknoster/table-driven-tests-gomock/pkg/names"
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
	"github.com/areknoster/table-driven-tests-gomock/pkg/people"
)

// namesLister is a facade which let's us mock some external dependency's behaviour
type namesLister interface {
	ListNames(justNice bool) ([]string, error)
}

type PartyService struct {
	namesLister namesLister
	helloer     party.Helloer
}

func NewService() *PartyService {
	return &PartyService{
		namesLister: &names.Service{},
		helloer:     people.Host{},
	}
}

func (s *PartyService) WelcomeEveryone() error {
	gusetNames, err := s.namesLister.ListNames(false)
	if err != nil {
		return fmt.Errorf("could not list gusetNames: %w", err)
	}
	for _, name := range gusetNames {
		fmt.Println(s.helloer.Hello(name))
	}
	return nil
}
