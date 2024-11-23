package handler

import (
	"math/rand"
	"net/http"
	"time"

	"cat-story/internal/repo/models"
	"cat-story/internal/repo/story"
	storypage "cat-story/internal/view/pages/story"
)

func Add(w http.ResponseWriter, r *http.Request) error {
	text := r.FormValue("text")
	name := r.FormValue("name")
	date := time.Now().String()

	stor := models.Story{
		Name: name,
		Date: date,
		Text: text,
	}

	return story.Add(r.Context(), stor)
}

func slicestring(str string) string {
	if len(str) > 10 {
        str = str[:len(str)-10]
    }

	return str
}

func torighttime(strs []models.Story) []models.Story{
	var s []models.Story

	for _, st := range strs{
		s = append(s, models.Story{Name: st.Name, Text: st.Text, Date: slicestring(st.Date)})
	}

	return s
}

func GetAll(w http.ResponseWriter, r *http.Request) error {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    
    stori, err := story.Get()
	stories := torighttime(stori)

    if err != nil {
        return err
    }
	st := stories[rand.Intn(len(stories))]

	story := storypage.StoryProps{
		Name: st.Name,
		Text: st.Text,
		Date: st.Date,
	}

	return storypage.RandStory(story).Render(r.Context(), w)
}