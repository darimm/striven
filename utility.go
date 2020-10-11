package striven

import (
	"fmt"

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
		return nil, fmt.Errorf("%d %w", resp.StatusCode(), err)
	}

	return resp, nil
}
