package party

type VisitorGroup string

const (
	NiceVisitor    VisitorGroup = "nice"
	NotNiceVisitor VisitorGroup = "not-nice"
)

type VisitorLister interface {
	ListVisitors(who VisitorGroup) ([]Visitor, error)
}
