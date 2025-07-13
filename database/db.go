package database

import (
	"errors"

	"github.com/legzdev/OSM-Changesets-Bot/env"
	"github.com/legzdev/OSM-Changesets-Bot/types"
)

var ErrNotFound = errors.New("not found")

type ChangesetsRepo interface {
	GetLatest() (types.ChangesetID, error)
	SetLatest(id types.ChangesetID) error
}

func Init() (ChangesetsRepo, error) {
	return NewBolt(env.DataBaseURL)
}
