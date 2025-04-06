package main

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

func main() {
}

// GetTransactionsOutput is a response model
type GetTransactionsOutput struct {
	Transactions []TransactionOutput
}

// TransactionOutput is the transaction struct for InternalHistory
type TransactionOutput struct {
	Channel          Channel
	Amount           int64
	AmountArrears    int64
	AmountPrepaid    int64
	OrderID          string
	Status           PaymentStatus
	AccountNumbers   []string
	TxnTimeString    string // TODO: should be time.Time
	Method           string
	RecurringEnabled bool
	PaidFor          string
	PaidBy           string
}

// Channel is the enum for QR code generated by channel
type Channel string

// ChannelPrefix is the enum for QR code prefix value generated by channel
type ChannelPrefix string

const (
	ChannelInfinityPrefix ChannelPrefix = "MQR"
	ChannelUPortalPrefix  ChannelPrefix = "UQR"

	// ChannelInfinity represents Channel/Service
	ChannelInfinity Channel = "INFINITY"
	// ChannelUPortal represents Channel/Service
	ChannelUPortal Channel = "UPORTAL"
	// ChannelPaperBill represents Channel/Service
	ChannelPaperBill Channel = "PAPERBILL"
	// ChannelExternal represents Channel/Service
	ChannelExternal Channel = "EXTERNAL"
	// ChannelSAP represents Channel/Service, use in our atomic job only
	ChannelSAP Channel = "SAP"
	// ChannelSystem represents Channel/Service use in termination card recurring
	ChannelSystem Channel = "SYSTEM"
)

// IsValid checks if the value is in the enum list
func (c Channel) IsValid() bool {
	switch c {
	// no need to check ChannelSAP, ChannelSAP only valid in our atomic job
	case ChannelInfinity,
		ChannelUPortal,
		ChannelPaperBill,
		ChannelExternal,
		ChannelSAP,
		ChannelSystem:
		return true
	}
	return false
}

// PaymentStatus is the enum for Transaction/Notification PaymentStatus
type PaymentStatus string

const (
	// PaymentStatusSuccess represents payment has been successfully processed.
	PaymentStatusSuccess PaymentStatus = "SUCCESS"
	// PaymentStatusFailed represents payment failed to be processed.
	PaymentStatusFailed PaymentStatus = "FAILED"
	// PaymentStatusPending represents payment is pending
	PaymentStatusPending PaymentStatus = "PENDING"
	// PaymentStatusRejected represents payment has been rejected
	PaymentStatusRejected PaymentStatus = "REJECTED"
	// PaymentStatusExpired represents billing did not get result from payment after cutoff
	PaymentStatusExpired PaymentStatus = "EXPIRED"
	// PaymentStatusCancelled represents payment cancelled to be processed.
	// just using for unidollar payment
	PaymentStatusCancelled PaymentStatus = "CANCELLED"
	// PaymentStatusBoltRefunded represents refunded payment in Bolt1 side. It was migrated as R to Njord.
	PaymentStatusBoltRefunded PaymentStatus = "R"
)

// IsValid checks if the value is in the enum list
func (s PaymentStatus) IsValid() bool {
	switch s {
	case PaymentStatusSuccess,
		PaymentStatusFailed,
		PaymentStatusPending,
		PaymentStatusRejected,
		PaymentStatusCancelled,
		PaymentStatusExpired,
		PaymentStatusBoltRefunded:
		return true
	}
	return false
}

// String returns enum in string
func (s PaymentStatus) String() string {
	return string(s)
}

func historyDeltaReflect(oldTxns GetTransactionsOutput, newTxns GetTransactionsOutput) {
	// arbitrary number of up till 10 as some users may have years of transactions; too many
	tCountLimit := 10

	for tCount := 0; tCount < tCountLimit; tCount++ {
		if !reflect.DeepEqual(newTxns.Transactions[0], oldTxns.Transactions[0]) {
			// _ = 1 + 1
		}
	}
}

func historyDeltaCMP(oldTxns GetTransactionsOutput, newTxns GetTransactionsOutput) {
	// arbitrary number of up till 10 as some users may have years of transactions; too many
	tCountLimit := 10

	for tCount := 0; tCount < tCountLimit; tCount++ {
		if diff := cmp.Diff(oldTxns.Transactions, newTxns.Transactions); diff != "" {
			// _ = 1 + 1
		}
	}
}

func historyDeltaIfAllFields(oldTxns GetTransactionsOutput, newTxns GetTransactionsOutput) {
	// arbitrary number of up till 10 as some users may have years of transactions; too many
	tCountLimit := 10
	isPrintAll := false

	for tCount := 0; tCount < tCountLimit; tCount++ {
		if oldTxns.Transactions[0].RecurringEnabled != newTxns.Transactions[0].RecurringEnabled {
			isPrintAll = true
		}
		if oldTxns.Transactions[0].OrderID != newTxns.Transactions[0].OrderID {
			isPrintAll = true
		}
		if oldTxns.Transactions[0].TxnTimeString != newTxns.Transactions[0].TxnTimeString {
			isPrintAll = true
		}
		if oldTxns.Transactions[0].AccountNumbers[0] != newTxns.Transactions[0].AccountNumbers[0] {
			isPrintAll = true
		}

		if oldTxns.Transactions[0].Channel != newTxns.Transactions[0].Channel {
			isPrintAll = true
		}

		if oldTxns.Transactions[0].Amount != newTxns.Transactions[0].Amount {
			isPrintAll = true
		}

		if oldTxns.Transactions[0].AmountArrears != newTxns.Transactions[0].AmountArrears {
			isPrintAll = true
		}

		if oldTxns.Transactions[0].Status != newTxns.Transactions[0].Status {
			isPrintAll = true
		}

		if oldTxns.Transactions[0].AmountPrepaid != newTxns.Transactions[0].AmountPrepaid {
			isPrintAll = true
		}

		if oldTxns.Transactions[0].Method != newTxns.Transactions[0].Method {
			isPrintAll = true
		}

		if oldTxns.Transactions[0].PaidBy != newTxns.Transactions[0].PaidBy {
			isPrintAll = true
		}

		if oldTxns.Transactions[0].PaidFor != newTxns.Transactions[0].PaidFor {
			isPrintAll = true
		}

		if isPrintAll {
			// _ = 1 + 1
		}
	}
}
