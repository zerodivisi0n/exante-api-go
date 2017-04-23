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

type Exchange struct {
	ID      string
	Name    string
	Country string
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

func (c *Client) Exchanges() ([]Exchange, error) {
	var exchanges []Exchange
	if err := c.apiCall("/exchanges", "symbols", nil, &exchanges); err != nil {
		return nil, err
	}
	return exchanges, nil
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
