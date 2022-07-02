package providers

import (
	"github.com/shopspring/decimal"

	"converter"
)

type Client interface {
	Convert(amount decimal.Decimal, from, to converter.Currency) (decimal.Decimal, error)
}
