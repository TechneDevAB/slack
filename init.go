// +build !appengine

package slack

import (
	"net/http"
)

func initHttpClient(req *http.Request) {
	httpClient = &http.Client{}
}
