package names

import "fmt"

type Service struct{}

type Group string

const (
	Nice    Group = "nice"
	NotNice Group = "not-nice"
)

func (s *Service) ListNames(who Group) ([]string, error) {
	switch who {
	case Nice:
		return []string{
			"Maria",
		}, nil
	case NotNice:
		return nil, fmt.Errorf("could not fetch people who are not nice")
	}
	return nil, fmt.Errorf("don't know who should be fetched")
}
