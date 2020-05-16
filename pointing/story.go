package pointing

type Story struct {
	Name  string         `json:"name"`
	Votes map[string]int `json:"votes"`
}

func newStory(name string) (*Story, error) {
	return &Story{
		name,
		make(map[string]int),
	}, nil
}

func (s *Story) Vote(userId string, vote int) error {
	s.Votes[userId] = vote
	return nil
}
