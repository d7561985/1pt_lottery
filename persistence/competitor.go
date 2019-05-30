package persistence

import (
	"fmt"
	"github.com/d7561985/heroku_boilerplate/pkg/database"
	"github.com/kataras/iris/core/errors"
	"github.com/lib/pq"
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
	quoted := pq.QuoteIdentifier("competitors")
	sql := fmt.Sprintf("INSERT INTO %s(dice_id,name, uuid,win) VALUES ($1,$2,$3,$4)", quoted)
	result, err := database.D.Exec(sql, c.DiceID, c.Name, c.UUID, c.Win)
	if err != nil {
		return
	}

	n, err := result.RowsAffected()
	fmt.Printf("[ROW:%d]%s [%d,%q,%q,%t]\n", n, sql, c.DiceID, c.Name, c.UUID, c.Win)

	// add to storage
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
	/*sql := `SELECT id, dice_id,name,uuid,win FROM "competitors" WHERE uuid LIKE $1 LIMIT 1`
	fmt.Println(sql, uuid)

	rows, err := database.D.Query(sql, uuid)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&c.ID, &c.DiceID, &c.Name, &c.UUID, &c.Win); err != nil {
			return
		}
	}

	return*/
}

func (l *Competitors) All() (err error) {
	sql := `SELECT id, dice_id,name,uuid,win FROM "competitors"`
	fmt.Println(sql)

	rows, err := database.D.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		c := Competitor{}
		if err = rows.Scan(&c.ID, &c.DiceID, &c.Name, &c.UUID, &c.Win); err != nil {
			return
		}
		*l = append(*l, c)
	}

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
