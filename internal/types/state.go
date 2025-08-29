package types

import (
	"github.com/NachoGz/blog-aggregator/internal/config"
	"github.com/NachoGz/blog-aggregator/internal/database"
)

type State struct {
	DB  *database.Queries
	Cfg *config.Config
}

func NewState(db *database.Queries, cfg *config.Config) *State {
	return &State{
		DB:  db,
		Cfg: cfg,
	}
}
