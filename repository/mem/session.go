package mem

import (
	"scrumctl.dev/m/pointing"
)

type SessionRepository struct {
	sessions map[string]pointing.Session
}

func (sr *SessionRepository) Create(session pointing.Session) error {
	sr.sessions[session.SessionId] = session
	return nil
}

func (sr *SessionRepository) Find(sessionId string) (pointing.Session, error) {
	return sr.sessions[sessionId], nil
}

func (sr *SessionRepository) Update(session pointing.Session) error {
	sr.sessions[session.SessionId] = session
	return nil
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{
		sessions: make(map[string]pointing.Session),
	}
}
