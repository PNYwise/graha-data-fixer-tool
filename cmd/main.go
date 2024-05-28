package main

import (
	"fmt"
	"log"

	"github.com/PNYwise/graha-data-fixer-tool/internal"
)

func main() {
	/**
	 * Open DB connection
	 *
	**/
	internal.ConnectDb()
	defer func() {
		if err := internal.CloseDb(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	if err := internal.Ping(); err != nil {
		log.Fatalf("Error ping database connection: %v", err)
	}

	bankPayRepo := internal.NewBankPayRepository(internal.DB.Db)
	paymentDetailRepo := internal.NewPaymentDetailRepository(internal.DB.Db)
	_ = paymentDetailRepo
	bankPays, err := bankPayRepo.FindFilteredBankPays()
	if err != nil {
		panic(err)
	}

	fmt.Printf("bankPay len: %v\n", len(*bankPays))

	for _, bankPay := range *bankPays {
		fmt.Printf("bankPay: %v\n", bankPay)
	}

	var paymentDetails []internal.PaymentDetailEntity
	for _, bankPay := range *bankPays {
		paymentDetail := internal.PaymentDetailEntity{
			SubTotal:         bankPay.Amount,
			Paid:             bankPay.Amount,
			IsConfirmed:      true,
			CreatedBy:        bankPay.CreatedBy,
			BankId:           bankPay.BankId,
			PaymentMethodId:  bankPay.PaymentMethodId,
			Date:             bankPay.Date,
			ConfirmationDate: bankPay.Date,
			Note:             bankPay.Note,
			BankPayId:        bankPay.ID,
		}
		paymentDetails = append(paymentDetails, paymentDetail)
	}

	if err := paymentDetailRepo.CreatePaymentDetailAndUpdateBankPay(paymentDetails); err != nil {
		panic(err)
	}
}
