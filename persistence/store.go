package persistence

import (
	"github.com/d7561985/1pt_lottery/dto"
	"github.com/rs/zerolog/log"
	"sync"
)

var (
	// storage representation
	S = Store{}
)

type Store struct {
	sync.Map

	// online users only
	Online sync.Map
}

func (s *Store) Clean() {
	s.Range(func(key, _ interface{}) bool {
		s.Delete(key)
		return true
	})
}

func (s *Store) GetOnline() (num int, res []dto.UserRequest) {
	foreach(s.Online, func(w Competitor) bool {
		num++
		res = append(res, dto.UserRequest{Name: w.Name, Avatar: w.Avatar})
		return true
	})
	return
}

func (s *Store) FindUser(name string) (res *dto.UserRequest) {
	foreach(s.Map, func(w Competitor) bool {
		if w.Name != name {
			return true
		}

		res = &dto.UserRequest{Name: w.Name, Avatar: w.Avatar}
		return false
	})
	return
}

func foreach(store sync.Map, fn func(Competitor) bool) {
	store.Range(func(key, value interface{}) bool {
		w, ok := value.(Competitor)
		if !ok {
			log.Fatal().Interface("value", value).Msg("cast fail")
		}
		return fn(w)
	})
}
