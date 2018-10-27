package resourceadapter

import (
	"context"

	"github.com/BonexIO/go/services/horizon/internal/txsub"
	. "github.com/BonexIO/go/protocols/horizon"

)

// Populate fills out the details
func PopulateTransactionResultCodes(ctx context.Context,
	dest *TransactionResultCodes,
	fail *txsub.FailedTransactionError,
) (err error) {

	dest.TransactionCode, err = fail.TransactionResultCode()
	if err != nil {
		return
	}

	dest.OperationCodes, err = fail.OperationResultCodes()
	if err != nil {
		return
	}

	return
}
