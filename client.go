package exante

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	baseUrl  = "https://api-demo.exante.eu/md/1.0"
	tokenTTL = 10 * time.Second
)

type Duration int

const (
	Duration1Minute   Duration = 60
	Duration5Minutes           = 300
	Duration10Minutes          = 600
	Duration15Minutes          = 900
	Duration1Hour              = 3600
	Duration6Hours             = 21600
	Duration1Day               = 86400
)

type Client struct {
	conn *http.Client

	// Auth info
	clientID      string
	applicationID string
	sharedKey     string
}

type Timestamp struct {
	// Embedding time allows to use Timestamp as original time
	time.Time
}

type Symbol struct {
	ID          string
	Name        string
	Description string
	Ticker      string
	Type        string
	Exchange    string
	Country     string
	Currency    string
	MPI         float64
	Group       string
	Expiration  Timestamp
	OptionData  struct {
		Right       string
		StrikePrice float64
	}
}

type SymbolSpecification struct {
	Leverage           float64
	LotSize            float64
	ContractMultiplier float64
	PriceUnit          float64
	Units              string
}

type SymbolScheduleInterval struct {
	Name   string
	Period struct {
		Start Timestamp
		End   Timestamp
	}
}

type Exchange struct {
	ID      string
	Name    string
	Country string
}

type Group struct {
	Group    string
	Name     string
	Types    []string
	Exchange string
}

type OHLC struct {
	Timestamp Timestamp
	Open      float64
	High      float64
	Low       float64
	Close     float64
}

func NewClient(clientID, applicationID, sharedKey string) *Client {
	return &Client{
		conn: &http.Client{
			Timeout: 30 * time.Second,
		},
		clientID:      clientID,
		applicationID: applicationID,
		sharedKey:     sharedKey,
	}
}

func (c *Client) Symbols() ([]Symbol, error) {
	var symbols []Symbol
	if err := c.apiCall("/symbols", "symbols", nil, &symbols); err != nil {
		return nil, err
	}
	return symbols, nil
}

func (c *Client) Symbol(id string) (*Symbol, error) {
	var symbol Symbol
	if err := c.apiCall("/symbols/"+id, "symbols", nil, &symbol); err != nil {
		return nil, err
	}
	return &symbol, nil
}

func (c *Client) SymbolSpecification(id string) (*SymbolSpecification, error) {
	var spec SymbolSpecification
	if err := c.apiCall("/symbols/"+id+"/specification", "symbols", nil, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}

func (c *Client) SymbolSchedule(id string) ([]SymbolScheduleInterval, error) {
	var schedule struct{ Intervals []SymbolScheduleInterval }
	if err := c.apiCall("/symbols/"+id+"/schedule", "symbols", nil, &schedule); err != nil {
		return nil, err
	}
	return schedule.Intervals, nil
}

func (c *Client) Exchanges() ([]Exchange, error) {
	var exchanges []Exchange
	if err := c.apiCall("/exchanges", "symbols", nil, &exchanges); err != nil {
		return nil, err
	}
	return exchanges, nil
}

func (c *Client) ExchangeSymbols(id string) ([]Symbol, error) {
	var symbols []Symbol
	if err := c.apiCall("/exchanges/"+id, "symbols", nil, &symbols); err != nil {
		return nil, err
	}
	return symbols, nil
}

func (c *Client) Types() ([]string, error) {
	var res []struct{ ID string }
	if err := c.apiCall("/types", "symbols", nil, &res); err != nil {
		return nil, err
	}
	types := make([]string, len(res))
	for i, v := range res {
		types[i] = v.ID
	}
	return types, nil
}

func (c *Client) TypeSymbols(id string) ([]Symbol, error) {
	var symbols []Symbol
	if err := c.apiCall("/types/"+id, "symbols", nil, &symbols); err != nil {
		return nil, err
	}
	return symbols, nil
}

func (c *Client) Groups() ([]Group, error) {
	var groups []Group
	if err := c.apiCall("/groups", "symbols", nil, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

func (c *Client) GroupSymbols(id string) ([]Symbol, error) {
	var symbols []Symbol
	if err := c.apiCall("/groups/"+id, "symbols", nil, &symbols); err != nil {
		return nil, err
	}
	return symbols, nil
}

func (c *Client) GroupNearestSymbol(id string) (*Symbol, error) {
	var symbol Symbol
	if err := c.apiCall("/groups/"+id+"/nearest", "symbols", nil, &symbol); err != nil {
		return nil, err
	}
	return &symbol, nil
}

func (c *Client) OHLC(symbolId string, duration Duration, from time.Time, to time.Time, size int) ([]OHLC, error) {
	var candles []OHLC
	durationStr := strconv.Itoa(int(duration))
	params := map[string]string{
		"from": strconv.FormatInt(from.Unix()*1000, 10),
		"to":   strconv.FormatInt(to.Unix()*1000, 10),
		"size": strconv.Itoa(size),
	}
	if err := c.apiCall("/ohlc/"+symbolId+"/"+durationStr, "ohlc", params, &candles); err != nil {
		return nil, err
	}
	return candles, nil
}

func (c *Client) apiCall(endpoint string, scope string, params map[string]string, result interface{}) error {
	req, err := http.NewRequest("GET", baseUrl+endpoint, nil)
	if err != nil {
		return err
	}
	token, err := c.signRequest(scope)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := c.conn.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return parseError(res, body)
	}
	return json.Unmarshal(body, result)
}

func (c *Client) signRequest(scope string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": c.clientID,
		"sub": c.applicationID,
		"aud": []string{scope},
		"iat": now.Unix(),
		"exp": now.Add(tokenTTL).Unix(),
	})

	return token.SignedString([]byte(c.sharedKey))
}

func parseError(res *http.Response, body []byte) error {
	// TODO: construct different errors
	return errors.New(string(body))
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	ts := t.Time.Unix()
	return []byte(strconv.FormatInt(ts*1000, 10)), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 0)
	if err != nil {
		return err
	}

	t.Time = time.Unix(ts/1000, 0)

	return nil
}
