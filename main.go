package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"scrumctl.dev/m/pointing"
	"scrumctl.dev/m/repository/mem"
)

var pointingService pointing.App

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/session/{sessionId}", GetSession)
	r.Post("/session", CreateSession)
	r.Post("/session/{sessionId}/story", CreateStory)
	r.Put("/session/{sessionId}/story/{storyName}/vote", StoryVote)
	sr := mem.NewSessionRepository()
	pointingService = *pointing.NewService(sr)
	_ = http.ListenAndServe(":3000", r)
}

func StoryVote(_ http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "sessionId")
	story := chi.URLParam(r, "storyName")
	svr := &storyVoteRequest{}
	_ = render.Decode(r, svr)
	_ = pointingService.StoryVote(sid, story, svr.UserId, svr.Vote)

}

func GetSession(w http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "sessionId")
	s, _ := pointingService.FindSession(sid)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, s)
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	csr := &createSessionRequest{}
	_ = render.Decode(r, csr)
	s, _ := pointingService.CreateSession(csr.UserName)
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, s)
}

func CreateStory(w http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "sessionId")
	csr := &createStoryRequest{}
	_ = render.Decode(r, csr)
	s, _ := pointingService.CreateStory(sid, csr.StoryName)
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, s)
}

type createSessionRequest struct {
	UserName string `json:"userName"`
}

type createStoryRequest struct {
	StoryName string `json:"storyName"`
}

type storyVoteRequest struct {
	UserId string `json:"userId"`
	Vote   int    `json:"vote"`
}
