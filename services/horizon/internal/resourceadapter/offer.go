package resourceadapter

import (
	"context"

	"github.com/stivens13/go/amount"
	"github.com/stivens13/go/services/horizon/internal/assets"
	"github.com/stivens13/go/services/horizon/internal/db2/core"
	"github.com/stivens13/go/services/horizon/internal/httpx"
	. "github.com/stivens13/go/protocols/horizon"
	"github.com/stivens13/go/support/render/hal"
	"github.com/stivens13/go/services/horizon/internal/db2/history"
)

func PopulateOffer(ctx context.Context, dest *Offer, row core.Offer, ledger history.Ledger) {
	dest.ID = row.OfferID
	dest.PT = row.PagingToken()
	dest.Seller = row.SellerID
	dest.Amount = amount.String(row.Amount)
	dest.PriceR.N = row.Pricen
	dest.PriceR.D = row.Priced
	dest.Price = row.PriceAsString()
	dest.Buying = Asset{
		Type:   assets.MustString(row.BuyingAssetType),
		Code:   row.BuyingAssetCode.String,
		Issuer: row.BuyingIssuer.String,
	}
	dest.Selling = Asset{
		Type:   assets.MustString(row.SellingAssetType),
		Code:   row.SellingAssetCode.String,
		Issuer: row.SellingIssuer.String,
	}
	dest.LastModifiedLedger = row.Lastmodified
	dest.LastModifiedTime = ledger.ClosedAt
	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	dest.Links.Self = lb.Linkf("/offers/%d", row.OfferID)
	dest.Links.OfferMaker = lb.Linkf("/accounts/%s", row.SellerID)
	return
}