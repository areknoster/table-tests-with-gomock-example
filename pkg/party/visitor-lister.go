package party

import "fmt"

var _ fmt.Stringer = Visitor{}

func (v Visitor) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Surname)
}

type VisitorGroup string

const (
	NiceVisitor    VisitorGroup = "nice"
	NotNiceVisitor VisitorGroup = "not-nice"
)

type VisitorLister interface {
	ListVisitors(who VisitorGroup) ([]Visitor, error)
}
