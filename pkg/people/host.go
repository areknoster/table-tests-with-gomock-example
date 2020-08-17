package people

import (
	"fmt"
)

type Host struct{}

func (h Host) Hello(name string) string {
	return fmt.Sprintf("Hello %s, nice to see you!", name)
}
