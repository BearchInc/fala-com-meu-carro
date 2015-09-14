package middleware
import (
	"appengine"
	"github.com/go-martini/martini"
	"net/http"
	"strings"
)

type RequestLocation struct {
	Country string
}

var (
	usLocationInfo = RequestLocation{Country: "us"}
	brLocationInfo = RequestLocation{Country: "br"}
)

func RequestLocationProvider(gaeContext appengine.Context, mContext martini.Context, request *http.Request) {
	locationCode := request.Header.Get("X-AppEngine-Country")

	location := locationFromCode(locationCode)

	gaeContext.Infof("Using location code: %v", location)
	mContext.Map(location)
}

func locationFromCode(code string) RequestLocation {
	if strings.ToLower(code) == "br" {
		return brLocationInfo
	}

	return usLocationInfo
}