package DocModel

import "time"

type DocImpl struct {
	ID           int
	KbID         int
	Title        string
	TitleUrl     string
	Content      string
	CreatorId    int
	LastEditorId int
	UpdatedAt    time.Time
	Deleted      bool
	GroupID      int
}
