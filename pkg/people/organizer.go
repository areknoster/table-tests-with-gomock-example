package people

import (
	"fmt"
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
)

type Organizer struct{}

func (o *Organizer) ListVisitors(who party.VisitorGroup) ([]party.Visitor, error) {
	switch who {
	case party.NiceVisitor:
		return []party.Visitor{
			{"John", "Smith"},
		}, nil
	case party.NotNiceVisitor:
		return nil, fmt.Errorf("could not fetch people who are not nice")
	}
	return nil, fmt.Errorf("don't know who should be fetched")
}
