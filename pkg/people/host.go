package people

import (
	"fmt"
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
)

type Host struct{}

var _ party.Greeter = Host{}

func (h Host) Hello(name string) string {
	return fmt.Sprintf("Hello %s, nice to see you!", name)
}
