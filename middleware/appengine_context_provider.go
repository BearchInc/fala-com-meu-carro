package middleware
import (
	"appengine"
	"fmt"
	"net/http"
	"github.com/go-martini/martini"
)

func AppengineContextProvider(c martini.Context, request *http.Request) {
	gae := appengine.NewContext(request)
	namespace := appengine.ModuleName(gae)

	context, err := appengine.Namespace(gae, namespace)

	if err != nil {
		panic(fmt.Sprintf("Could not create GAE context: %v", err))
	}

	c.MapTo(context, (*appengine.Context)(nil))
}
