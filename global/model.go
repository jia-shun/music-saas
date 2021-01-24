package global

import (
	"time"
)

type MODEL struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:主键"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
