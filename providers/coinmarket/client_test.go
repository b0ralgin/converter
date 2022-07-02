package coinmarket

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	cli := NewClient("https://sandbox-api.coinmarketcap.com", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")
	price, err := cli.Convert(decimal.NewFromInt(1), "BTC", "ETH")
	require.NoError(t, err)
	require.True(t, price.IsPositive())
}
