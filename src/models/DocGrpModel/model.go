package DocGrpModel

import "time"

type DocGrpImpl struct {
	ID         int
	GroupName  string
	KbID       int
	CreateTime time.Time
	DocCount   int
	CreatorId  int
}
