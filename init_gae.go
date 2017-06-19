// +build appengine

package slack

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func initHttpClient(req *http.Request) {
	ctx := appengine.NewContext(req)
	httpClient = urlfetch.Client(ctx)
}
