package slack

import (
	"net/http"
	"net/url"
)

type Slack struct {
	token      string
	HttpClient *http.Client
}

// Create a slack client with an API token.
func New(token string) *Slack {
	return &Slack{
		token: token,
	}
}

func (sl *Slack) urlValues() *url.Values {
	uv := url.Values{}
	uv.Add("token", sl.token)
	return &uv
}
