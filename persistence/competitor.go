package persistence

import (
	"fmt"
	"github.com/d7561985/heroku_boilerplate/pkg/database"
	"github.com/lib/pq"
)

type Competitor struct {
	ID     int64
	DiceID uint
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
	fmt.Printf("[ROW:%d]%s\n", n, sql)
	return
}

func (c *Competitor) Find(uuid string) (err error) {
	rows, err := database.D.Query(`SELECT id, dice_id,name,uuid,win FROM "competitors" WHERE uuid LIKE $1 LIMIT 1`, uuid)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&c.ID, &c.DiceID, &c.Name, &c.UUID, &c.Win); err != nil {
			return
		}
	}

	return
}

func (l *Competitors) All() (err error) {
	rows, err := database.D.Query(`SELECT id, dice_id,name,uuid,win FROM "competitors"`)
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
