package persistence

import (
	"github.com/kataras/iris/core/errors"
)

var (
	errNotExist = errors.New("not exist")
	errBadCast  = errors.New("bad cast")
)

type Competitor struct {
	ID     int64
	DiceID uint
	Avatar string
	Name   string
	UUID   string
	Win    bool
}
type Competitors []Competitor

func (c *Competitor) Create() (err error) {
	S.Store(c.UUID, *c)
	return
}

func (c *Competitor) Find(uuid string) (err error) {
	i, ok := S.Load(uuid)
	if !ok {
		return errNotExist
	}

	res, ok := i.(Competitor)
	if !ok {
		return errBadCast
	}

	*c = res
	return
}

func (l *Competitors) FillByStorage() Competitors {
	S.Range(func(_, value interface{}) bool {
		if add, ok := value.(Competitor); ok {
			*l = append(*l, add)
		}
		return true
	})
	return *l
}
