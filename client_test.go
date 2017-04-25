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

func TestSymbol(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/symbols/AAPL.NASDAQ").
		Reply(200).
		BodyString(`{
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
	}`)

	symbol, err := NewClient("", "", "").Symbol("AAPL.NASDAQ")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "AAPL.NASDAQ", symbol.ID)
	assert.Equal(t, "Apple", symbol.Name)
	assert.Equal(t, "AAPL", symbol.Ticker)
	assert.Equal(t, "STOCK", symbol.Type)
	assert.Equal(t, "Apple", symbol.Description)
	assert.Equal(t, "NASDAQ", symbol.Exchange)
	assert.Equal(t, "US", symbol.Country)
	assert.Equal(t, "USD", symbol.Currency)
	assert.Equal(t, 0.01, symbol.MPI)
}

func TestSymbolSpecification(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/symbols/AAPL.NASDAQ/specification").
		Reply(200).
		BodyString(`{
			"leverage": 0.2,
			"lotSize": 1.0,
			"contractMultiplier": 1.0,
			"priceUnit": 1.0,
			"units": "Shares"
	}`)

	spec, err := NewClient("", "", "").SymbolSpecification("AAPL.NASDAQ")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 0.2, spec.Leverage)
	assert.Equal(t, 1.0, spec.LotSize)
	assert.Equal(t, 1.0, spec.ContractMultiplier)
	assert.Equal(t, 1.0, spec.PriceUnit)
	assert.Equal(t, "Shares", spec.Units)
}

func TestSymbolSchedule(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/symbols/AAPL.NASDAQ/schedule").
		Reply(200).
		BodyString(`{"intervals":[
			{"name":"PreMarket","period":{"start":1493020800000,"end":1493040600000}},
			{"name":"MainSession","period":{"start":1493040600000,"end":1493064000000}},
			{"name":"AfterMarket","period":{"start":1493064000000,"end":1493078400000}}
	]}`)

	schedule, err := NewClient("", "", "").SymbolSchedule("AAPL.NASDAQ")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(schedule), "Invalid schedule length")
	assert.Equal(t, "PreMarket", schedule[0].Name)
	assert.Equal(t, Timestamp{time.Unix(1493020800, 0)}, schedule[0].Period.Start)
	assert.Equal(t, Timestamp{time.Unix(1493040600, 0)}, schedule[0].Period.End)
	assert.Equal(t, "MainSession", schedule[1].Name)
	assert.Equal(t, Timestamp{time.Unix(1493040600, 0)}, schedule[1].Period.Start)
	assert.Equal(t, Timestamp{time.Unix(1493064000, 0)}, schedule[1].Period.End)
	assert.Equal(t, "AfterMarket", schedule[2].Name)
	assert.Equal(t, Timestamp{time.Unix(1493064000, 0)}, schedule[2].Period.Start)
	assert.Equal(t, Timestamp{time.Unix(1493078400, 0)}, schedule[2].Period.End)
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

func TestExchangeSymbols(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/exchanges/NYSE").
		Reply(200).
		BodyString(`[
		{
			"id": "LEN.B.NYSE",
			"ticker": "LEN.B",
			"name": "Lennar Corporation",
			"description": "Lennar Corporation",
			"exchange": "NYSE",
			"country": "US",
			"i18n": {},
			"type": "STOCK",
			"mpi": 0.01,
			"currency": "USD"
		},
		{
			"id": "BK.NYSE",
			"ticker": "BK",
			"name": "Bank Of New York Mellon Corporation",
			"description": "Bank Of New York Mellon Corporation",
			"exchange": "NYSE",
			"country": "US",
			"i18n": {},
			"type": "STOCK",
			"mpi": 0.01,
			"currency": "USD"
		},
		{
			"id": "GJR.NYSE",
			"ticker": "GJR",
			"name": "Synthetic Fixed-Income Securities",
			"description": "Synthetic Fixed-Income Securities",
			"exchange": "NYSE",
			"country": "US",
			"i18n": {},
			"type": "STOCK",
			"mpi": 0.01,
			"currency": "USD"
		}
	]`)

	symbols, err := NewClient("", "", "").ExchangeSymbols("NYSE")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(symbols), "Invalid symbols length")
	// LEN.B.NYSE
	assert.Equal(t, "LEN.B.NYSE", symbols[0].ID)
	assert.Equal(t, "Lennar Corporation", symbols[0].Name)
	assert.Equal(t, "LEN.B", symbols[0].Ticker)
	assert.Equal(t, "STOCK", symbols[0].Type)
	assert.Equal(t, "Lennar Corporation", symbols[0].Description)
	assert.Equal(t, "NYSE", symbols[0].Exchange)
	assert.Equal(t, "US", symbols[0].Country)
	assert.Equal(t, "USD", symbols[0].Currency)
	assert.Equal(t, 0.01, symbols[0].MPI)
	// BK.NYSE
	assert.Equal(t, "BK.NYSE", symbols[1].ID)
	assert.Equal(t, "Bank Of New York Mellon Corporation", symbols[1].Name)
	assert.Equal(t, "BK", symbols[1].Ticker)
	assert.Equal(t, "STOCK", symbols[1].Type)
	assert.Equal(t, "Bank Of New York Mellon Corporation", symbols[1].Description)
	assert.Equal(t, "NYSE", symbols[1].Exchange)
	assert.Equal(t, "US", symbols[1].Country)
	assert.Equal(t, "USD", symbols[1].Currency)
	assert.Equal(t, 0.01, symbols[1].MPI)
	// GJR.NYSE
	assert.Equal(t, "GJR.NYSE", symbols[2].ID)
	assert.Equal(t, "Synthetic Fixed-Income Securities", symbols[2].Name)
	assert.Equal(t, "GJR", symbols[2].Ticker)
	assert.Equal(t, "STOCK", symbols[2].Type)
	assert.Equal(t, "Synthetic Fixed-Income Securities", symbols[2].Description)
	assert.Equal(t, "NYSE", symbols[2].Exchange)
	assert.Equal(t, "US", symbols[2].Country)
	assert.Equal(t, "USD", symbols[2].Currency)
	assert.Equal(t, 0.01, symbols[2].MPI)
}

func TestTypes(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/types").
		Reply(200).
		BodyString(`[
			{"id": "CALENDAR_SPREAD"},
			{"id": "FUND"},
			{"id": "FX_SPOT"},
			{"id": "CURRENCY"},
			{"id": "BOND"},
			{"id": "FUTURE"},
			{"id": "STOCK"},
			{"id": "OPTION"}
	]`)

	types, err := NewClient("", "", "").Types()

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 8, len(types), "Invalid types length")
	assert.Equal(t,
		[]string{"CALENDAR_SPREAD", "FUND", "FX_SPOT", "CURRENCY",
			"BOND", "FUTURE", "STOCK", "OPTION"}, types)
}

func TestTypeSymbols(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/types/STOCK").
		Reply(200).
		BodyString(`[
			{
				"id":"MAXD.OTCMKTS",
				"ticker":"MAXD",
				"name":"Max Sound",
				"description":"Max Sound",
				"exchange":"OTCMKTS",
				"country":"US",
				"i18n":{},
				"type":"STOCK",
				"mpi":1.0E-4,
				"currency":"USD"
			},
			{
				"id":"QINC.NASDAQ",
				"ticker":"QINC",
				"name":"First Trust RBA Quality Income ETF",
				"description":"First Trust RBA Quality Income ETF",
				"exchange":"NASDAQ",
				"country":"US",
				"i18n":{},
				"type":"STOCK",
				"mpi":0.01,
				"currency":"USD"
			},
			{
				"id":"DGRE.NASDAQ",
				"ticker":"DGRE",
				"name":"WisdomTree Emerging Markets Quality Dividend Growth Fund",
				"description":"WisdomTree Emerging Markets Quality Dividend Growth Fund",
				"exchange":"NASDAQ",
				"country":"US",
				"i18n":{},
				"type":"STOCK",
				"mpi":0.01,
				"currency":"USD"
			}
	]`)

	symbols, err := NewClient("", "", "").TypeSymbols("STOCK")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(symbols), "Invalid types length")
	// MAXD.OTCMKTS
	assert.Equal(t, "MAXD.OTCMKTS", symbols[0].ID)
	assert.Equal(t, "Max Sound", symbols[0].Name)
	assert.Equal(t, "MAXD", symbols[0].Ticker)
	assert.Equal(t, "STOCK", symbols[0].Type)
	assert.Equal(t, "Max Sound", symbols[0].Description)
	assert.Equal(t, "OTCMKTS", symbols[0].Exchange)
	assert.Equal(t, "US", symbols[0].Country)
	assert.Equal(t, "USD", symbols[0].Currency)
	assert.Equal(t, 1.0E-4, symbols[0].MPI)
	// QINC.NASDAQ
	assert.Equal(t, "QINC.NASDAQ", symbols[1].ID)
	assert.Equal(t, "First Trust RBA Quality Income ETF", symbols[1].Name)
	assert.Equal(t, "QINC", symbols[1].Ticker)
	assert.Equal(t, "STOCK", symbols[1].Type)
	assert.Equal(t, "First Trust RBA Quality Income ETF", symbols[1].Description)
	assert.Equal(t, "NASDAQ", symbols[1].Exchange)
	assert.Equal(t, "US", symbols[1].Country)
	assert.Equal(t, "USD", symbols[1].Currency)
	assert.Equal(t, 0.01, symbols[1].MPI)
	// DGRE.NASDAQ
	assert.Equal(t, "DGRE.NASDAQ", symbols[2].ID)
	assert.Equal(t, "WisdomTree Emerging Markets Quality Dividend Growth Fund", symbols[2].Name)
	assert.Equal(t, "DGRE", symbols[2].Ticker)
	assert.Equal(t, "STOCK", symbols[2].Type)
	assert.Equal(t, "WisdomTree Emerging Markets Quality Dividend Growth Fund", symbols[2].Description)
	assert.Equal(t, "NASDAQ", symbols[2].Exchange)
	assert.Equal(t, "US", symbols[2].Country)
	assert.Equal(t, "USD", symbols[2].Currency)
	assert.Equal(t, 0.01, symbols[2].MPI)
}

func TestGroups(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/groups").
		Reply(200).
		BodyString(`[
		{"group":"ABX","name":"Barrick Gold","types":["OPTION"],"exchange":"CBOE"},
		{"group":"MA","name":"Mastercard","types":["OPTION"],"exchange":"CBOE"},
		{"group":"ATLN","name":"Actelion","types":["OPTION"],"exchange":"EUREX"}
	]`)

	groups, err := NewClient("", "", "").Groups()

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(groups), "Invalid groups length")
	// ABX
	assert.Equal(t, "ABX", groups[0].Group)
	assert.Equal(t, "Barrick Gold", groups[0].Name)
	assert.Equal(t, []string{"OPTION"}, groups[0].Types)
	assert.Equal(t, "CBOE", groups[0].Exchange)
	// MA
	assert.Equal(t, "MA", groups[1].Group)
	assert.Equal(t, "Mastercard", groups[1].Name)
	assert.Equal(t, []string{"OPTION"}, groups[1].Types)
	assert.Equal(t, "CBOE", groups[1].Exchange)
	// ATLN
	assert.Equal(t, "ATLN", groups[2].Group)
	assert.Equal(t, "Actelion", groups[2].Name)
	assert.Equal(t, []string{"OPTION"}, groups[2].Types)
	assert.Equal(t, "EUREX", groups[2].Exchange)
}

func TestGroupSymbols(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/groups/MA").
		Reply(200).
		BodyString(`[
			{
				"id":"MA.CBOE.15U2017.P140",
				"ticker":"MA",
				"name":"Mastercard",
				"description":"Mastercard 15 Sep 2017 PUT 140",
				"exchange":"CBOE",
				"country":"US",
				"i18n":{},
				"type":"OPTION",
				"mpi":0.01,
				"currency":"USD",
				"group":"MA",
				"expiration":1505487600000,
				"optionData":{
					"right":"PUT",
					"strikePrice":140
				}
			}
	]`)

	symbols, err := NewClient("", "", "").GroupSymbols("MA")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(symbols), "Invalid symbols length")
	// MA
	assert.Equal(t, "MA.CBOE.15U2017.P140", symbols[0].ID)
	assert.Equal(t, "Mastercard", symbols[0].Name)
	assert.Equal(t, "MA", symbols[0].Ticker)
	assert.Equal(t, "OPTION", symbols[0].Type)
	assert.Equal(t, "Mastercard 15 Sep 2017 PUT 140", symbols[0].Description)
	assert.Equal(t, "CBOE", symbols[0].Exchange)
	assert.Equal(t, "US", symbols[0].Country)
	assert.Equal(t, "USD", symbols[0].Currency)
	assert.Equal(t, 0.01, symbols[0].MPI)
	assert.Equal(t, "MA", symbols[0].Group)
	assert.Equal(t, Timestamp{time.Unix(1505487600, 0)}, symbols[0].Expiration)
	assert.Equal(t, "PUT", symbols[0].OptionData.Right)
	assert.Equal(t, 140.0, symbols[0].OptionData.StrikePrice)
}

func TestGroupNearestSymbol(t *testing.T) {
	defer gock.Off()

	gock.New(baseUrl).
		Get("/groups/6R/nearest").
		Reply(200).
		BodyString(`{
			"id":"6R.CME.K2017",
			"ticker":"6R",
			"name":"RUB/USD",
			"description":"Futures On RUB/USD May 2017",
			"exchange":"CME",
			"country":"US",
			"i18n":{},
			"type":"FUTURE",
			"mpi":5.0E-6,
			"currency":"USD",
			"group":"6R",
			"expiration":1494813600000
	}`)

	symbol, err := NewClient("", "", "").GroupNearestSymbol("6R")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "6R.CME.K2017", symbol.ID)
	assert.Equal(t, "RUB/USD", symbol.Name)
	assert.Equal(t, "6R", symbol.Ticker)
	assert.Equal(t, "FUTURE", symbol.Type)
	assert.Equal(t, "Futures On RUB/USD May 2017", symbol.Description)
	assert.Equal(t, "CME", symbol.Exchange)
	assert.Equal(t, "US", symbol.Country)
	assert.Equal(t, "USD", symbol.Currency)
	assert.Equal(t, 5e-6, symbol.MPI)
	assert.Equal(t, "6R", symbol.Group)
	assert.Equal(t, Timestamp{time.Unix(1494813600, 0)}, symbol.Expiration)
}
