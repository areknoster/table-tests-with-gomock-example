package names

import "fmt"

type Service struct{}

func (s *Service) ListNames(justNice bool) ([]string, error) {
	if !justNice {
		return nil, fmt.Errorf("could not fetch non-nice people")
	}
	return []string{
		"Bob",
		"Maria",
		"Anthony",
		"Penelope",
	}, nil
}
