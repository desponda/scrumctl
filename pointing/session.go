package pointing

import "github.com/google/uuid"

type Session struct {
	SessionId   string            `json:"sessionId"`
	Stories     map[string]*Story `json:"stories"`
	Users       map[string]*User  `json:"users"`
	LatestStory string            `json:"latestStory"`
}

func (s *Session) CreateStory(name string) (Story, error) {
	story, _ := newStory(name)
	s.Stories[name] = story
	s.LatestStory = name
	return *story, nil
}

func (s *Session) StoryVote(story string, userId string, vote int) error {
	_ = s.Stories[story].Vote(userId, vote)
	return nil
}

func (s *Session) Join(name string) error {
	u, _ := newUser(name)
	s.Users[name] = u
	return nil
}

func newSession(createdBy string) (*Session, error) {
	u, _ := newUser(createdBy)
	users := make(map[string]*User)
	users[createdBy] = u
	return &Session{
		uuid.New().String(),
		make(map[string]*Story),
		users,
		"",
	}, nil
}
