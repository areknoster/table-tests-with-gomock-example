package people

import (
	"fmt"
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
)

type Organizer struct{

}

func (o *Organizer) ListVisitors(who party.VisitorGroup) ([]party.Visitor, error) {

}
