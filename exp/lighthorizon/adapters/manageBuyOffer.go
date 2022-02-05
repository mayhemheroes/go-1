package adapters

import (
	"github.com/stellar/go/amount"
	"github.com/stellar/go/protocols/horizon/base"
	"github.com/stellar/go/protocols/horizon/operations"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

func populateManageBuyOfferOperation(
	op *xdr.Operation,
	baseOp operations.Base,
) (operations.ManageBuyOffer, error) {
	manageBuyOffer := op.Body.ManageBuyOfferOp
	baseOp.Type = "manage_buy_offer"

	var (
		buyingAssetType string
		buyingCode      string
		buyingIssuer    string
	)
	err := manageBuyOffer.Buying.Extract(&buyingAssetType, &buyingCode, &buyingIssuer)
	if err != nil {
		return operations.ManageBuyOffer{}, errors.Wrap(err, "xdr.Asset.Extract error")
	}

	var (
		sellingAssetType string
		sellingCode      string
		sellingIssuer    string
	)
	err = manageBuyOffer.Selling.Extract(&sellingAssetType, &sellingCode, &sellingIssuer)
	if err != nil {
		return operations.ManageBuyOffer{}, errors.Wrap(err, "xdr.Asset.Extract error")
	}

	return operations.ManageBuyOffer{
		Offer: operations.Offer{
			Base:   baseOp,
			Amount: amount.String(manageBuyOffer.BuyAmount),
			Price:  manageBuyOffer.Price.String(),
			PriceR: base.Price{
				N: int32(manageBuyOffer.Price.N),
				D: int32(manageBuyOffer.Price.D),
			},
			BuyingAssetType:    buyingAssetType,
			BuyingAssetCode:    buyingCode,
			BuyingAssetIssuer:  buyingIssuer,
			SellingAssetType:   sellingAssetType,
			SellingAssetCode:   sellingCode,
			SellingAssetIssuer: sellingIssuer,
		},
		OfferID: int64(manageBuyOffer.OfferId),
	}, nil
}
