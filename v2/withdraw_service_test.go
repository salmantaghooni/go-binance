package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type withdrawServiceTestSuite struct {
	baseTestSuite
}

func TestWithdrawService(t *testing.T) {
	suite.Run(t, new(withdrawServiceTestSuite))
}

func (s *withdrawServiceTestSuite) TestCreateWithdraw() {
	data := []byte(`
	{
		"id":"7213fea8e94b4a5593d507237e5a555b"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "USDT"
	withdrawOrderID := "testID"
	network := "ETH"
	address := "myaddress"
	addressTag := "xyz"
	amount := "0.01"
	transactionFeeFlag := true
	name := "eth"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":               coin,
			"withdrawOrderId":    withdrawOrderID,
			"network":            network,
			"address":            address,
			"addressTag":         addressTag,
			"amount":             amount,
			"transactionFeeFlag": transactionFeeFlag,
			"name":               name,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCreateWithdrawService().
		Coin(coin).
		WithdrawOrderID(withdrawOrderID).
		Network(network).
		Address(address).
		AddressTag(addressTag).
		Amount(amount).
		TransactionFeeFlag(transactionFeeFlag).
		Name(name).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("7213fea8e94b4a5593d507237e5a555b", res.ID)
}

func (s *withdrawServiceTestSuite) TestListWithdraws() {
	data := []byte(`
	[
		{
			"id":"7213fea8e94b4a5593d507237e5a555b",
			"withdrawOrderID": "",    
			"amount": "0.99",
			"transactionFee": "0.01",
			"address": "0x6915f16f8791d0a1cc2bf47c13a6b2a92000504b",
			"coin": "USDT",
			"txId": "0xdf33b22bdb2b28b1f75ccd201a4a4m6e7g83jy5fc5d5a9d1340961598cfcb0a1",
			"applyTime": "2019-10-12 11:12:02",
			"network": "ETH",
			"status": 4
		},
		{
			"id":"7213fea8e94b4a5534ggsd237e5a555b",
			"withdrawOrderID": "withdrawtest", 
			"amount": "999.9999",
			"transactionFee": "0.0001",
			"address": "463tWEBn5XZJSxLU34r6g7h8jtxuNcDbjLSjkn3XAXHCbLrTTErJrBWYgHJQyrCwkNgYvyV3z8zctJLPCZy24jvb3NiTcTJ",
			"addressTag": "342341222",
			"txId": "b3c6219639c8ae3f9cf010cdc24fw7f7yt8j1e063f9b4bd1a05cb44c4b6e2509",
			"coin": "XMR",
			"applyTime": "2019-10-12 11:12:02",
			"status": 4
		}
	]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "ETH"
	status := 0
	startTime := int64(1508198532000)
	endTime := int64(1508198532001)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":      coin,
			"status":    status,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	withdraws, err := s.client.NewListWithdrawsService().
		Coin(coin).
		Status(status).
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())
	r := s.r()
	r.NoError(err)

	s.Len(withdraws, 2)
	s.assertWithdrawEqual(&Withdraw{
		ID:              "7213fea8e94b4a5593d507237e5a555b",
		WithdrawOrderID: "",
		Amount:          "0.99",
		TransactionFee:  "0.01",
		Address:         "0x6915f16f8791d0a1cc2bf47c13a6b2a92000504b",
		AddressTag:      "",
		Coin:            "USDT",
		TxID:            "0xdf33b22bdb2b28b1f75ccd201a4a4m6e7g83jy5fc5d5a9d1340961598cfcb0a1",
		ApplyTime:       "2019-10-12 11:12:02",
		Network:         "ETH",
		Status:          4,
	}, withdraws[0])
	s.assertWithdrawEqual(&Withdraw{
		ID:              "7213fea8e94b4a5534ggsd237e5a555b",
		WithdrawOrderID: "withdrawOrderID",
		Amount:          "999.9999",
		TransactionFee:  "0.0001",
		Address:         "463tWEBn5XZJSxLU34r6g7h8jtxuNcDbjLSjkn3XAXHCbLrTTErJrBWYgHJQyrCwkNgYvyV3z8zctJLPCZy24jvb3NiTcTJ",
		AddressTag:      "342341222",
		TxID:            "b3c6219639c8ae3f9cf010cdc24fw7f7yt8j1e063f9b4bd1a05cb44c4b6e2509",
		Coin:            "XMR",
		ApplyTime:       "2019-10-12 11:12:02",
		Status:          4,
	}, withdraws[1])
}

func (s *withdrawServiceTestSuite) assertWithdrawEqual(e, a *Withdraw) {
	r := s.r()
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Address, a.Address, "Address")
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.ApplyTime, a.ApplyTime, "ApplyTime")
	r.Equal(e.Status, a.Status, "Status")
}
