package utils

import "together/be8/delivery/view/transaction"

type ConfigMidtrans interface {
	CreateTransaction(OrderID string, GrossAmt int64) map[string]interface{}
	FinishPayment(order string) transaction.ResponsePayment
}
