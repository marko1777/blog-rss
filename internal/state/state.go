package state

import (
	"github.com/marko1777/blog-rss/internal/config"
	"github.com/marko1777/blog-rss/internal/database"
)

type State struct {
	Cfg       *config.Config
	DBQueries *database.Queries
}
