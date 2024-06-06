package swagger

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = CORS(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func Docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	htmlFile, err := ioutil.ReadFile("docs/index.html")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		http.Error(w, "Could not read the documentation file", http.StatusInternalServerError)
		return
	}
	w.Write(htmlFile)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"LoginPost",
		strings.ToUpper("Post"),
		"/login",
		LoginPost,
	},

	Route{
		"RegisterPost",
		strings.ToUpper("Post"),
		"/register",
		RegisterPost,
	},

	Route{
		"MessageIdGet",
		strings.ToUpper("Get"),
		"/message/{id}",
		MessageIdGet,
	},

	Route{
		"MessagePost",
		strings.ToUpper("Post"),
		"/message",
		MessagePost,
	},

	Route{
		"Docs",
		strings.ToUpper("Get"),
		"/docs",
		Docs,
	},
}
