package party

import "fmt"

type Visitor struct {
	Name    string
	Surname string
}

var _ fmt.Stringer = Visitor{}

func (v Visitor) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Surname)
}
