package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gopkg.in/alecthomas/kingpin.v2"

	"converter"
	"converter/providers/coinmarket"
)

type config struct {
	APIKey string
	Host   string
}

func main() {
	cfg := config{}
	a := kingpin.New("coinconv", "program for converting currencies")
	a.Flag("apikey", "Allowed origins for CORS, splitted by `,`").
		Envar("COINCONV_API_KEY").
		Default("b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c").
		StringVar(&cfg.APIKey)
	a.Flag("host", "Allowed origins for CORS, splitted by `,`").
		Envar("COINCONV_HOST").
		Default("https://sandbox-api.coinmarketcap.com").
		StringVar(&cfg.Host)
	amountArg := a.Arg("amount", "amount to convert").Required().String()
	var fromCur string
	a.Arg("from", "convert from currency").Required().StringVar(&fromCur)
	var toCur string
	a.Arg("to", "convert to currency").Required().StringVar(&toCur)
	_, err := a.Parse(os.Args[1:])
	if err != nil {
		a.Usage(os.Args[1:])
		log.Fatal(errors.Wrap(err, "Error parsing command line arguments"))
	}

	if amountArg == nil {
		log.Fatal("amount isn't set")
	}
	amount, err := decimal.NewFromString(*amountArg)
	if err != nil {
		log.Fatal("wrong amount value")
	}
	cli := coinmarket.NewClient(cfg.Host, cfg.APIKey)
	result, err := cli.Convert(amount, converter.Currency(fromCur), converter.Currency(toCur))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to convert"))
	}
	fmt.Println(result.String())
}
