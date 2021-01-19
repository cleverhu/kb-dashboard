package KbModel

import "time"

type KbImpl struct {
	ID         string
	Name       string
	Desc       string
	Kind       int
	CreatorId  int
	Private    bool
	CreateTime time.Time
	State      string
}
