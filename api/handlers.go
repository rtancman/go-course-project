package api

import (
	"fmt"
    "log"
	"net/http"
    "github.com/rtancman/go-course-project/db"
    "github.com/rtancman/go-course-project/music"
)

type GetWelcomecHandler struct{}

func (h *GetWelcomecHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println("log Sua Api de musicas", nil)
    fmt.Fprintf(w, "Sua Api de musicas!")
}

type PostMusicHandler struct{
    Repository *db.MusicRepository
}

func (h *PostMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	m := &music.Music{Id: "1231", Name: "Juliana"}
	err := h.Repository.Create(m)

	if err == db.ErrDuplicatedPerson {
		log.Printf("%s is already created\n", m.Name)
	} else if err != nil {
		log.Println("Failed to create a person: ", err)
	}
    fmt.Fprintf(w, "OK!")
}
