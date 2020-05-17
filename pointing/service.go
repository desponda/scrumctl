package pointing

type App struct {
	sr SessionRepository
}

type SessionRepository interface {
	Create(Session) error
	Find(string) (Session, error)
	Update(Session) error
}

func (a *App) CreateSession(createdBy string) (Session, error) {
	ns, _ := newSession(createdBy)
	_ = a.sr.Create(*ns)
	return *ns, nil
}

func (a *App) CreateStory(sessionId string, storyName string) (Story, error) {
	s, _ := a.sr.Find(sessionId)
	story, _ := s.CreateStory(storyName)
	_ = a.sr.Update(s)
	return story, nil
}

func (a *App) FindSession(sessionId string) (Session, error) {
	s, _ := a.sr.Find(sessionId)
	return s, nil
}

func (a *App) StoryVote(sid string, story string, userId string, vote int) error {
	s, _ := a.sr.Find(sid)
	_ = s.StoryVote(story, userId, vote)
	_ = a.sr.Update(s)
	return nil
}

func (a *App) JoinSession(sid string, userName string) (Session, error) {
	s, _ := a.sr.Find(sid)
	_ = s.Join(userName)
	_ = a.sr.Update(s)
	return s, nil
}

func NewService(sr SessionRepository) *App {
	return &App{
		sr: sr,
	}
}
