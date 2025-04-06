package main

import (
	"testing"
)

var num = 1000

func BenchmarkReflect(t *testing.B) {
	oldTxn := GetTransactionsOutput{Transactions: []TransactionOutput{
		{
			Channel:          "Infinity",
			Amount:           1000,
			OrderID:          "123456789",
			Status:           "SUCCESS",
			AccountNumbers:   []string{"21345667"},
			TxnTimeString:    "1111111",
			Method:           "PAYNOW",
			RecurringEnabled: false,
			PaidFor:          "auth0|123456789",
			PaidBy:           "auth0|123456789",
		},
	}}
	newTxn := GetTransactionsOutput{Transactions: []TransactionOutput{
		{
			Channel:          "Infinity",
			Amount:           1000,
			OrderID:          "123456789",
			Status:           "SUCCESS",
			AccountNumbers:   []string{"21345667"},
			TxnTimeString:    "1111111",
			Method:           "CARD",
			RecurringEnabled: true,
			PaidFor:          "auth0|123456789",
			PaidBy:           "auth0|123456789",
		},
	}}
	for i := 0; i < num; i++ {
		historyDeltaReflect(oldTxn, newTxn)
	}
}

func BenchmarkIf(t *testing.B) {
	oldTxn := GetTransactionsOutput{Transactions: []TransactionOutput{
		{
			Channel:          "Infinity",
			Amount:           1000,
			OrderID:          "123456789",
			Status:           "SUCCESS",
			AccountNumbers:   []string{"21345667"},
			TxnTimeString:    "1111111",
			Method:           "PAYNOW",
			RecurringEnabled: false,
			PaidFor:          "auth0|123456789",
			PaidBy:           "auth0|123456789",
		},
	}}
	newTxn := GetTransactionsOutput{Transactions: []TransactionOutput{
		{
			Channel:          "Infinity",
			Amount:           1000,
			OrderID:          "123456789",
			Status:           "SUCCESS",
			AccountNumbers:   []string{"21345667"},
			TxnTimeString:    "1111111",
			Method:           "CARD",
			RecurringEnabled: true,
			PaidFor:          "auth0|123456789",
			PaidBy:           "auth0|123456789",
		},
	}}
	for i := 0; i < num; i++ {
		historyDeltaIfAllFields(oldTxn, newTxn)
	}
}

func BenchmarkCMP(t *testing.B) {
	oldTxn := GetTransactionsOutput{Transactions: []TransactionOutput{
		{
			Channel:          "Infinity",
			Amount:           1000,
			OrderID:          "123456789",
			Status:           "SUCCESS",
			AccountNumbers:   []string{"21345667"},
			TxnTimeString:    "1111111",
			Method:           "PAYNOW",
			RecurringEnabled: false,
			PaidFor:          "auth0|123456789",
			PaidBy:           "auth0|123456789",
		},
	}}
	newTxn := GetTransactionsOutput{Transactions: []TransactionOutput{
		{
			Channel:          "Infinity",
			Amount:           1000,
			OrderID:          "123456789",
			Status:           "SUCCESS",
			AccountNumbers:   []string{"21345667"},
			TxnTimeString:    "1111111",
			Method:           "CARD",
			RecurringEnabled: true,
			PaidFor:          "auth0|123456789",
			PaidBy:           "auth0|123456789",
		},
	}}
	for i := 0; i < num; i++ {
		historyDeltaCMP(oldTxn, newTxn)
	}
}
