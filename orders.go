package finamclient

import (
	"github.com/evsamsonov/FinamTradeGo/v2/tradeapi"
)

func (f *FinamClient) NewOrder(in *tradeapi.NewOrderRequest) (*tradeapi.NewOrderResult, error) {

	res, err := f.orders.NewOrder(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FinamClient) CancelOrder(transactionId int32) (*tradeapi.CancelOrderResult, error) {

	in := &tradeapi.CancelOrderRequest{
		ClientId:      f.clientId,
		TransactionId: transactionId,
	}

	res, err := f.orders.CancelOrder(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FinamClient) GetOrders(includeMatched, includeCanceled, includeActive bool) (*tradeapi.GetOrdersResult, error) {

	in := &tradeapi.GetOrdersRequest{
		ClientId:        f.clientId,
		IncludeMatched:  includeMatched,
		IncludeCanceled: includeCanceled,
		IncludeActive:   includeActive,
	}

	res, err := f.orders.GetOrders(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
