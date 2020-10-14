package striven

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gopkg.in/resty.v1"
)

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

const TimestampFormat = time.RFC3339 // same as ISO8601

var (
	_ json.Unmarshaler = (*Timestamp)(nil)
	_ json.Marshaler   = (*Timestamp)(nil)
)

func NewTimestamp(t time.Time) Timestamp {
	return Timestamp(t)
}

func NowTimestamp() Timestamp {
	return NewTimestamp(time.Now())
}

// UnmarshalJSON parses a nullable RFC3339 string into time.
func (t *Timestamp) UnmarshalJSON(v []byte) error {
	str := fmt.Sprintf("%s%s", strings.Trim(string(v), "\""), "Z")
	if str == "null" {
		return nil
	}

	r, err := time.Parse(TimestampFormat, str)
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

	return []byte(t.Format(TimestampFormat)), nil
}

func (t Timestamp) IsValid() bool {
	return !t.Time().IsZero()
}

func (t Timestamp) Format(fmt string) string {
	return t.Time().Format(fmt)
}

func (t Timestamp) Time() time.Time {
	return time.Time(t)
}
