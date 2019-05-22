package persistence

import (
	"fmt"
	"github.com/d7561985/heroku_boilerplate/pkg/database"
)

func InitDB() (err error) {
	fmt.Println("Init database")

	_, err = database.D.Query(`CREATE TABLE IF NOT EXISTS competitors
(
	id SERIAL,
	dice_id integer,
	name character varying(254) NOT NULL,
	uuid character varying(254) NOT NULL,
	win boolean,
	CONSTRAINT competitors_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_competitors_uuid ON competitors(uuid);`)
	if err != nil {
		return
	}

	return
}

func Clean() (err error) {
	fmt.Println("Clean database")
	_, err = database.D.Query(`DELETE FROM competitors;`)
	return
}
