package DocModel

import (
	"fmt"
	"time"
)

type DocImpl struct {
	ID           int       `gorm:"column:doc_id;primary_key"`
	KbID         int       `gorm:"column:kb_id"`
	Title        string    `gorm:"column:doc_title"`
	TitleUrl     string    `gorm:"column:doc_title_url"`
	Content      string    `gorm:"column:doc_content"`
	CreatorID    int       `gorm:"column:creator_id"`
	LastEditorID int       `gorm:"-"`
	UpdatedAt    time.Time `gorm:"-"`
	Removed      string    `gorm:"-"`
	GroupID      int       `gorm:"column:group_id"`
}

func (this *DocImpl) String() string {
	str := fmt.Sprintf(`{docId:%d,kbId:%d,title:%s,titleUrl:%s,content:%s,creatorID:%d,lastEditorID:%d,updatedAt:%s,removed:%s,groupID:%d}`,
		this.ID, this.KbID, this.Title, this.TitleUrl, this.Content, this.CreatorID, this.LastEditorID, this.UpdatedAt.Format("2006-01-02 15:04:05"), this.Removed, this.GroupID)
	return str
}

func New(attrs ...DocModelAttrFunc) *DocImpl {
	d := &DocImpl{}
	DocModelAttrFuncs(attrs).Apply(d)
	return d
}

func (this *DocImpl) Mutate(attrs ...DocModelAttrFunc) *DocImpl {
	DocModelAttrFuncs(attrs).Apply(this)
	return this
}


