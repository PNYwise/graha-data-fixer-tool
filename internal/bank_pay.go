package internal

import (
	"time"

	"gorm.io/gorm"
)

type BankPayEntity struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	Number          string `gorm:"unique;not null"`
	Note            string `gorm:"not null"`
	Amount          float64
	Date            string `gorm:"type:date;not null"`
	CreatedBy       int
	UpdatedBy       int
	BankId          *uint
	PaymentMethodId *uint
	PaymentDetailId uint
}

func (BankPayEntity) TableName() string {
	return "bank_pays"
}

type IBankPayRepository interface {
	FindFilteredBankPays() (*[]BankPayEntity, error)
}

type bankPayRepository struct {
	db *gorm.DB
}

func NewBankPayRepository(db *gorm.DB) IBankPayRepository {
	return &bankPayRepository{
		db,
	}
}

// FindAll implements IMemberRepository.
func (b *bankPayRepository) FindFilteredBankPays() (*[]BankPayEntity, error) {
	bankPays := new([]BankPayEntity)
	query := b.db.
		Where("payment_detail_id is null and payment_id is null").
		Order("id ASC").
		Find(&bankPays)

	if err := query.Error; err != nil {
		return nil, err
	}
	return bankPays, nil
}
