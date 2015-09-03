package middleware

import (
	"github.com/drborges/appx"
	"github.com/go-martini/martini"
	"appengine"
)

func AppxProvider(c martini.Context, context appengine.Context) {
	c.Map(appx.NewDatastore(context))
}