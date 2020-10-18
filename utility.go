package striven

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"
)

// IDNamePair is used pretty much everywhere in the API.
type IDNamePair struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//StrivenCurrency is the standard structure for displaying currency in Striven
type StrivenCurrency struct {
	CurrencyISOCode string  `json:"currencyISOCode"`
	ExchangeRate    float64 `json:"exchangeRate"`
}

func (s *Striven) apiGet(URI string) (*resty.Response, error) {
	err := s.validateAccessToken()
	if err != nil {
		return nil, err
	}
	client := resty.New()

	resp, err := client.R().
		SetAuthToken(s.Token.AccessToken).
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s%s", StrivenURL, URI))

	if resp.StatusCode() != 200 || err != nil {
		return nil, fmt.Errorf("%d %v", resp.StatusCode(), err)
	}

	return resp, nil
}

// Timestamp is a custom time field that can be unmarshalled direct from striven because it's timestamp is not RFC3339
type Timestamp time.Time

const timestampFormat = time.RFC3339 // same as ISO8601

var (
	_ json.Unmarshaler = (*Timestamp)(nil)
	_ json.Marshaler   = (*Timestamp)(nil)
)

// NewTimestamp returns a new Timestamp formatted time based on a real golang time
func NewTimestamp(t time.Time) Timestamp {
	return Timestamp(t)
}

// NowTimestamp returns a new timestamp formatted time with the current time.
func NowTimestamp() Timestamp {
	return NewTimestamp(time.Now())
}

// UnmarshalJSON parses a nullable RFC3339 string into time.
func (t *Timestamp) UnmarshalJSON(v []byte) error {
	str := strings.Trim(string(v), `"`)
	//str := fmt.Sprintf("%s%s", strings.Trim(string(v), "\""), "Z")
	if str == "null" {
		return nil
	}

	tz, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic("Cannot load Time America/New_York")
	}
	r, err := time.ParseInLocation("2006-01-02T15:04:05.999999", str, tz)
	if err != nil {
		return err
	}

	*t = Timestamp(r)
	return nil
}

// MarshalJSON returns null if Timestamp is not valid (zero). It returns the
// time formatted in RFC3339 otherwise.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if !t.IsValid() {
		return []byte("null"), nil
	}

	return []byte(t.Format(timestampFormat)), nil
}

// IsValid just returns whether or not a time is valid.
func (t Timestamp) IsValid() bool {
	return !t.Time().IsZero()
}

// Format is an Implementaion of the built in time lib's Format function
func (t Timestamp) Format(fmt string) string {
	return t.Time().Format(fmt)
}

// Time is an Implementation of the built in time lib's Time function
func (t Timestamp) Time() time.Time {
	return time.Time(t)
}
