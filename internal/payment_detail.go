package internal

import (
	"time"

	"gorm.io/gorm"
)

type PaymentDetailEntity struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	SubTotal         float64
	Paid             float64
	IsConfirmed      bool
	Date             string `gorm:"type:date;not null"`
	CreatedBy        int
	BankId           *uint
	PaymentMethodId  *uint
	ConfirmationDate string `gorm:"type:date"`
	Note             string
	// helper
	BankPayId uint `gorm:"-"`
}

func (PaymentDetailEntity) TableName() string {
	return "payment_details"
}

type IPaymentDetailRepository interface {
	CreatePaymentDetailAndUpdateBankPay(paymentDetails []PaymentDetailEntity) error
}

type paymentDetailRepository struct {
	db *gorm.DB
}

func NewPaymentDetailRepository(db *gorm.DB) IPaymentDetailRepository {
	return &paymentDetailRepository{
		db,
	}
}

// CreatePaymentDetailAndUpdateBankPay implements IPaymentDetailRepository.
func (p *paymentDetailRepository) CreatePaymentDetailAndUpdateBankPay(paymentDetails []PaymentDetailEntity) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(paymentDetails, 500).Error; err != nil {
			return err
		}
		for _, paymentDetail := range paymentDetails {
			err := tx.Model(&BankPayEntity{}).Where("id = ?", paymentDetail.BankPayId).Updates(
				BankPayEntity{
					PaymentDetailId: paymentDetail.ID,
					UpdatedBy:       1,
				},
			).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
