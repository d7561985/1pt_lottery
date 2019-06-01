package persistence

import (
	"github.com/d7561985/1pt_lottery/dto"
	"github.com/rs/zerolog/log"
	"sync"
)

var (
	S      = Store{}
	Online sync.Map
)

type Store struct {
	sync.Map
}

func (s *Store) Clean() {
	s.Range(func(key, _ interface{}) bool {
		s.Delete(key)
		return true
	})
}

func (s *Store) Online() (num int, res []dto.UserRequest) {
	Online.Range(func(key, value interface{}) bool {
		w, ok := value.(Competitor)
		if !ok {
			log.Fatal().Interface("value", value).Msg("cast fail")
		}
		num++
		res = append(res, dto.UserRequest{Name: w.Name, Avatar: w.Avatar})
		return true
	})
	return

}
