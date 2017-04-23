package exante

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSymbols(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/symbols").
		Reply(200).
		BodyString(`[
		{
			"id": "AAPL.NASDAQ",
			"name": "Apple",
			"ticker": "AAPL",
			"type": "STOCK",
			"description": "Apple",
			"exchange": "NASDAQ",
			"country": "US",
			"currency": "USD",
			"i18n": {},
			"mpi": 0.01
		},
		{
			"id": "GOOG.NASDAQ",
			"name": "Alphabet Class C",
			"ticker": "GOOG",
			"type": "STOCK",
			"description": "Alphabet Class C",
			"exchange": "NASDAQ",
			"country": "US",
			"currency": "USD",
			"i18n": {},
			"mpi": 0.01
		},
		{
			"id": "USD/RUB.EXANTE",
			"ticker": "USD/RUB",
			"type": "CURRENCY",
			"description": "USD/RUB",
			"currency": "RUB",
			"i18n": {},
			"mpi": 0.0001
		},
    {
        "id": "6R.CME.M2018",
        "name": "RUB/USD",
        "ticker": "6R",
        "type": "FUTURE",
        "description": "Futures On RUB/USD Jun 2018",
        "exchange": "CME",
        "country": "US",
        "currency": "USD",
        "i18n": {},
        "mpi": 5e-06,
        "group": "6R",
        "expiration": 1529028000000
    },
    {
        "id": "SPX.CBOE.16M2017.P2250",
        "name": "S&P 500 Index",
        "ticker": "SPX",
        "type": "OPTION",
        "description": "Options On S&P 500 Index 16 Jun 2017 PUT 2250",
        "exchange": "CBOE",
        "country": "US",
        "currency": "USD",
        "i18n": {},
        "mpi": 0.01,
        "group": "SPX.CBOE",
        "expiration": 1497626100000,
        "optionData": {
            "right": "PUT",
            "strikePrice": 2250
        }
    }
	]`)

	symbols, err := NewClient("", "", "").Symbols()

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 5, len(symbols), "Invalid symbols length")
	// AAPL Stock
	assert.Equal(t, "AAPL.NASDAQ", symbols[0].ID)
	assert.Equal(t, "Apple", symbols[0].Name)
	assert.Equal(t, "AAPL", symbols[0].Ticker)
	assert.Equal(t, "STOCK", symbols[0].Type)
	assert.Equal(t, "Apple", symbols[0].Description)
	assert.Equal(t, "NASDAQ", symbols[0].Exchange)
	assert.Equal(t, "US", symbols[0].Country)
	assert.Equal(t, "USD", symbols[0].Currency)
	assert.Equal(t, 0.01, symbols[0].MPI)
	// GOOG Stock
	assert.Equal(t, "GOOG.NASDAQ", symbols[1].ID)
	assert.Equal(t, "Alphabet Class C", symbols[1].Name)
	assert.Equal(t, "GOOG", symbols[1].Ticker)
	assert.Equal(t, "STOCK", symbols[1].Type)
	assert.Equal(t, "Alphabet Class C", symbols[1].Description)
	assert.Equal(t, "NASDAQ", symbols[1].Exchange)
	assert.Equal(t, "US", symbols[1].Country)
	assert.Equal(t, "USD", symbols[1].Currency)
	assert.Equal(t, 0.01, symbols[1].MPI)
	// RUB Currency
	assert.Equal(t, "USD/RUB.EXANTE", symbols[2].ID)
	assert.Equal(t, "USD/RUB", symbols[2].Ticker)
	assert.Equal(t, "CURRENCY", symbols[2].Type)
	assert.Equal(t, "USD/RUB", symbols[2].Description)
	assert.Equal(t, "RUB", symbols[2].Currency)
	assert.Equal(t, 0.0001, symbols[2].MPI)
	// RUB Future
	assert.Equal(t, "6R.CME.M2018", symbols[3].ID)
	assert.Equal(t, "RUB/USD", symbols[3].Name)
	assert.Equal(t, "6R", symbols[3].Ticker)
	assert.Equal(t, "FUTURE", symbols[3].Type)
	assert.Equal(t, "Futures On RUB/USD Jun 2018", symbols[3].Description)
	assert.Equal(t, "CME", symbols[3].Exchange)
	assert.Equal(t, "US", symbols[3].Country)
	assert.Equal(t, "USD", symbols[3].Currency)
	assert.Equal(t, 5e-6, symbols[3].MPI)
	assert.Equal(t, "6R", symbols[3].Group)
	assert.Equal(t, Timestamp{time.Unix(1529028000, 0)}, symbols[3].Expiration)
	// SPX Option
	assert.Equal(t, "SPX.CBOE.16M2017.P2250", symbols[4].ID)
	assert.Equal(t, "S&P 500 Index", symbols[4].Name)
	assert.Equal(t, "SPX", symbols[4].Ticker)
	assert.Equal(t, "OPTION", symbols[4].Type)
	assert.Equal(t, "Options On S&P 500 Index 16 Jun 2017 PUT 2250", symbols[4].Description)
	assert.Equal(t, "CBOE", symbols[4].Exchange)
	assert.Equal(t, "US", symbols[4].Country)
	assert.Equal(t, "USD", symbols[4].Currency)
	assert.Equal(t, 0.01, symbols[4].MPI)
	assert.Equal(t, "SPX.CBOE", symbols[4].Group)
	assert.Equal(t, Timestamp{time.Unix(1497626100, 0)}, symbols[4].Expiration)
	assert.Equal(t, "PUT", symbols[4].OptionData.Right)
	assert.Equal(t, 2250.0, symbols[4].OptionData.StrikePrice)
}

func TestExchanges(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/exchanges").
		Reply(200).
		BodyString(`[
		{"id":"EURONEXT LISBOA stocks","name":"EURONEXT: Euronext Lisboa","country":"FR"},
		{"id":"NYSE ARCA","name":"NYSE ARCA: Archipelago Exchange","country":"RU"},
    {"id":"NYSE","name":"NYSE: New York Stock Exchange","country":"US"}
	]`)

	exchanges, err := NewClient("", "", "").Exchanges()

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(exchanges), "Invalid exchanges length")
	assert.Equal(t, "EURONEXT LISBOA stocks", exchanges[0].ID)
	assert.Equal(t, "EURONEXT: Euronext Lisboa", exchanges[0].Name)
	assert.Equal(t, "FR", exchanges[0].Country)
	assert.Equal(t, "NYSE ARCA", exchanges[1].ID)
	assert.Equal(t, "NYSE ARCA: Archipelago Exchange", exchanges[1].Name)
	assert.Equal(t, "RU", exchanges[1].Country)
	assert.Equal(t, "NYSE", exchanges[2].ID)
	assert.Equal(t, "NYSE: New York Stock Exchange", exchanges[2].Name)
	assert.Equal(t, "US", exchanges[2].Country)
}
