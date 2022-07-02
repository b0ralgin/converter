package coinmarket

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/shopspring/decimal"
	"gopkg.in/resty.v1"

	"converter"
	"converter/providers"
)

const path = "/v2/tools/price-conversion"

type client struct {
	conn *resty.Client
}

func NewClient(host string, key string) providers.Client {
	return &client{
		conn: resty.New().SetHostURL(host).SetHeader("X-CMC_PRO_API_KEY", key),
	}
}

func (c client) Convert(amount decimal.Decimal, from, to converter.Currency) (decimal.Decimal, error) {
	var r response
	resp, err := c.conn.R().SetQueryParams(map[string]string{
		"amount":  amount.String(),
		"symbol":  string(from),
		"convert": string(to),
	}).SetResult(&r).Get(path)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "failed to get data from coinmarket")
	}
	if resp.StatusCode() != http.StatusOK {
		return decimal.Zero, fmt.Errorf("failed to convert: %s", string(resp.Body()))
	}
	quota, ok := r.Data[from]
	if !ok {
		return decimal.Zero, errors.New("can't find from currency in response")
	}
	price, ok := quota.Quote[to]
	if !ok {
		return decimal.Zero, errors.New("can't find to currency in response")
	}
	return price.Price, nil
}

type response struct {
	Data map[converter.Currency]struct {
		Quote map[converter.Currency]struct {
			Price decimal.Decimal
		}
	}
}
