package wager

import (
	"net/http"

	"project/common/failure"
	"project/project/lib/wager/sql"
)

var (
	MaxBuyRequestsPerSecond = 3
)

// Error codes...
var (
	selling_price_under_threshold   = "selling_price_under_threshold"
	buying_price_over_selling_price = "buying_price_over_selling_price"
	wager_out_of_item               = "wager_out_of_item"
	wager_not_found                 = "wager_not_found"
	buying_price_not_accepted       = "buying_price_not_accepted"
)

var ErrorCodeMap = map[string]int{
	selling_price_under_threshold:   http.StatusBadRequest,
	buying_price_over_selling_price: http.StatusBadRequest,
	wager_out_of_item:               http.StatusBadRequest,
	wager_not_found:                 http.StatusNotFound,
	buying_price_not_accepted:       http.StatusBadRequest,
}

var (
	SellingPriceUnderThreshold failure.ErrorWraper = failure.NewFailure(
		selling_price_under_threshold,
		ErrorCodeMap[selling_price_under_threshold],
	)

	BuyingPriceOverSellingPrice failure.ErrorWraper = failure.NewFailure(
		buying_price_over_selling_price,
		ErrorCodeMap[buying_price_over_selling_price],
	)

	WagerOutOfItem failure.ErrorWraper = failure.NewFailure(
		wager_out_of_item,
		ErrorCodeMap[wager_out_of_item],
	)

	WagerNotFound failure.ErrorWraper = failure.NewFailure(
		wager_not_found,
		ErrorCodeMap[wager_not_found],
	)

	BuyingPriceNotAccepted failure.ErrorWraper = failure.NewFailure(
		buying_price_not_accepted,
		ErrorCodeMap[buying_price_not_accepted],
	)
)

// Sql...
var (
	AddWagerSql                = sql.AddWager
	FetchWagerForUpdateSql     = sql.FetchWagerForUpdate
	AddWagerTxnLogSql          = sql.AddWagerTxnLog
	UpdateWagerSellingPriceSql = sql.UpdateWagerSellingPrice
	FetchWagersSql             = sql.FetchWagers
)
