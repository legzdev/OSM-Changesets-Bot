package types

import "time"

type ChangesetID = int64

type Changeset struct {
	ID          ChangesetID
	Title       string
	Description string
	Create      string
	Modify      string
	Delete      string
	Username    string
	Date        time.Time
}
