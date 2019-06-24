package model

import "time"

type Comment struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
	OrgName   string     `json:"orgName"`
	Comment   string     `json:"comment"`
}

type CommentRequest struct {
	Comment string `json:"comment" valid:"required~comment cannot empty"`
}
