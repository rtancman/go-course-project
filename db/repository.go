package db

import (
    "errors"

	"gopkg.in/mgo.v2"

    "github.com/rtancman/go-course-project/music"
)

const MusicCollection = "music"

var ErrDuplicatedPerson = errors.New("Duplicated person")

type MusicRepository struct {
	session *mgo.Session
}

func NewMusicRepository(session *mgo.Session) *MusicRepository {
	return &MusicRepository{session}
}

func (r *MusicRepository) Create(m *music.Music) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	err := collection.Insert(m)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrDuplicatedPerson
		}
	}
	return err
}
