package resourceadapter

import (
	"context"

	"github.com/BonexIO/go/services/horizon/internal/db2/history"
	. "github.com/BonexIO/go/protocols/horizon"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
