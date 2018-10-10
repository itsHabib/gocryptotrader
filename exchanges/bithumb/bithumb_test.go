package bithumb

import (
	"testing"

	"github.com/thrasher-/gocryptotrader/config"
	"github.com/thrasher-/gocryptotrader/currency/symbol"
	exchange "github.com/thrasher-/gocryptotrader/exchanges"
)

// Please supply your own keys here for due diligence testing
const (
	testAPIKey    = ""
	testAPISecret = ""
)

var b Bithumb

func TestSetDefaults(t *testing.T) {
	b.SetDefaults()
}

func TestSetup(t *testing.T) {
	cfg := config.GetConfig()
	cfg.LoadConfig("../../testdata/configtest.json")
	bitConfig, err := cfg.GetExchangeConfig("Bithumb")
	if err != nil {
		t.Error("Test Failed - Bithumb Setup() init error")
	}

	bitConfig.AuthenticatedAPISupport = true
	bitConfig.APIKey = testAPIKey
	bitConfig.APISecret = testAPISecret

	b.Setup(bitConfig)
}

func TestGetTradablePairs(t *testing.T) {
	t.Parallel()
	_, err := b.GetTradablePairs()
	if err != nil {
		t.Error("test failed - Bithumb GetTradablePairs() error", err)
	}
}

func TestGetTicker(t *testing.T) {
	t.Parallel()
	_, err := b.GetTicker("btc")
	if err != nil {
		t.Error("test failed - Bithumb GetTicker() error", err)
	}
}

func TestGetAllTickers(t *testing.T) {
	t.Parallel()
	_, err := b.GetAllTickers()
	if err != nil {
		t.Error("test failed - Bithumb GetAllTickers() error", err)
	}
}

func TestGetOrderBook(t *testing.T) {
	t.Parallel()
	_, err := b.GetOrderBook("btc")
	if err != nil {
		t.Error("test failed - Bithumb GetOrderBook() error", err)
	}
}

func TestGetTransactionHistory(t *testing.T) {
	t.Parallel()
	_, err := b.GetTransactionHistory("btc")
	if err != nil {
		t.Error("test failed - Bithumb GetTransactionHistory() error", err)
	}
}

func TestGetAccountInfo(t *testing.T) {
	t.Parallel()
	_, err := b.GetAccountInfo()
	if err == nil {
		t.Error("test failed - Bithumb GetAccountInfo() error", err)
	}
}

func TestGetAccountBalance(t *testing.T) {
	t.Parallel()
	_, err := b.GetAccountBalance()
	if err == nil {
		t.Error("test failed - Bithumb GetAccountBalance() error", err)
	}
}

func TestGetWalletAddress(t *testing.T) {
	t.Parallel()
	_, err := b.GetWalletAddress("")
	if err == nil {
		t.Error("test failed - Bithumb GetWalletAddress() error", err)
	}
}

func TestGetLastTransaction(t *testing.T) {
	t.Parallel()
	_, err := b.GetLastTransaction()
	if err == nil {
		t.Error("test failed - Bithumb GetLastTransaction() error", err)
	}
}

func TestGetOrders(t *testing.T) {
	t.Parallel()
	_, err := b.GetOrders("1337", "bid", "100", "", "BTC")
	if err == nil {
		t.Error("test failed - Bithumb GetOrders() error", err)
	}
}

func TestGetUserTransactions(t *testing.T) {
	t.Parallel()
	_, err := b.GetUserTransactions()
	if err == nil {
		t.Error("test failed - Bithumb GetUserTransactions() error", err)
	}
}

func TestPlaceTrade(t *testing.T) {
	t.Parallel()
	_, err := b.PlaceTrade("btc", "bid", 0, 0)
	if err == nil {
		t.Error("test failed - Bithumb PlaceTrade() error", err)
	}
}

func TestGetOrderDetails(t *testing.T) {
	t.Parallel()
	_, err := b.GetOrderDetails("1337", "bid", "btc")
	if err == nil {
		t.Error("test failed - Bithumb GetOrderDetails() error", err)
	}
}

func TestCancelTrade(t *testing.T) {
	t.Parallel()
	_, err := b.CancelTrade("", "", "")
	if err == nil {
		t.Error("test failed - Bithumb CancelTrade() error", err)
	}
}

func TestWithdrawCrypto(t *testing.T) {
	t.Parallel()
	_, err := b.WithdrawCrypto("LQxiDhKU7idKiWQhx4ALKYkBx8xKEQVxJR", "", "ltc", 0)
	if err == nil {
		t.Error("test failed - Bithumb WithdrawCrypto() error", err)
	}
}

func TestRequestKRWDepositDetails(t *testing.T) {
	t.Parallel()
	_, err := b.RequestKRWDepositDetails()
	if err == nil {
		t.Error("test failed - Bithumb RequestKRWDepositDetails() error", err)
	}
}

func TestRequestKRWWithdraw(t *testing.T) {
	t.Parallel()
	_, err := b.RequestKRWWithdraw("102_bank", "1337", 1000)
	if err == nil {
		t.Error("test failed - Bithumb RequestKRWWithdraw() error", err)
	}
}

func TestMarketBuyOrder(t *testing.T) {
	t.Parallel()
	_, err := b.MarketBuyOrder("btc", 0)
	if err == nil {
		t.Error("test failed - Bithumb MarketBuyOrder() error", err)
	}
}

func TestMarketSellOrder(t *testing.T) {
	t.Parallel()
	_, err := b.MarketSellOrder("btc", 0)
	if err == nil {
		t.Error("test failed - Bithumb MarketSellOrder() error", err)
	}
}

func TestGetFee(t *testing.T) {
	t.Parallel()
	b.SetDefaults()
	TestSetup(t)

	if resp, err := b.GetFee(exchange.CryptocurrencyTradeFee, symbol.BTC+symbol.LTC, 1, 1, false, false); resp != float64(0.0015) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0.0015), resp)
	}

	if resp, err := b.GetFee(exchange.CryptocurrencyTradeFee, symbol.BTC+symbol.LTC, 10000000000, -1000000000, true, true); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
	}

	if resp, err := b.GetFee(exchange.CryptocurrencyTradeFee, symbol.BTC+symbol.LTC, 1, 1, true, false); resp != float64(0.0015) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0.0015), resp)
	}

	if resp, err := b.GetFee(exchange.CryptocurrencyTradeFee, symbol.BTC+symbol.LTC, 10000000000, -1000000000, false, true); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
	}

	if resp, err := b.GetFee(exchange.CryptocurrencyWithdrawalFee, symbol.BTC, 1, 5, false, false); resp != float64(0.001) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
	}

	if resp, err := b.GetFee(exchange.CyptocurrencyDepositFee, symbol.BTC, 1, 0.001, false, false); resp != float64(0.001) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0.001), resp)
	}

	if resp, err := b.GetFee(exchange.CyptocurrencyDepositFee, symbol.BTC, 1, 555, false, false); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
	}

	if resp, err := b.GetFee(exchange.InternationalBankDepositFee, symbol.BTC, 1, 1, false, false); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
	}

	if resp, err := b.GetFee(exchange.InternationalBankWithdrawalFee, symbol.BTC, 1, 1, false, false); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Recieved: %f", float64(0), resp)
	}

}
