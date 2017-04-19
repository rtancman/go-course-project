package main

import (
	"log"
	"net/http"
	"github.com/dimfeld/httptreemux"
	"github.com/rtancman/go-course-project/api"
    "github.com/rtancman/go-course-project/db"
    "gopkg.in/mgo.v2"
)

func main() {

    session, err := mgo.Dial("localhost:27017/go-course-project")
	if err != nil {
		log.Fatal(err)
	}

	Repository := db.NewMusicRepository(session)

    addr := "127.0.0.1:8081"
    router := httptreemux.NewContextMux()
    router.Handler(http.MethodGet, "/", &api.GetWelcomecHandler{})
	router.Handler(http.MethodPost, "/music", &api.PostMusicHandler{Repository: Repository})
    log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
