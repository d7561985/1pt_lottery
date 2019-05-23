package persistence

import (
	"github.com/d7561985/heroku_boilerplate/pkg/database"
	"github.com/rs/zerolog/log"
)

func InitDB() (err error) {
	log.Info().Msg("init database")
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

	loadStorage()
	return
}

func Clean() (err error) {
	log.Info().Msg("clean database")
	_, err = database.D.Query(`DELETE FROM competitors;`)
	return
}

func loadStorage() {
	log.Info().Msg("load storage")

	// competitors
	com := make(Competitors, 0)
	if err := com.All(); err != nil {
		panic(err)
	}

	for _, c := range com {
		S.Store(c.UUID, c)
	}
}
