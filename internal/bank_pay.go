package internal

import "time"

type PurchaseReceivedEntity struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	Number    string `gorm:"unique;not null"`
	Note      string `gorm:"not null"`
	Amount    int
	Date      string `gorm:"type:date;not null"`
	CreatedBy int    `gorm:"column:createdBy;not null"`
}

func (PurchaseReceivedEntity) TableName() string {
	return "purchase_received"
}
