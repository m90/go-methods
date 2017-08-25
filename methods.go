package methods

import (
	"net/http"
)

func contains(needle string, haystack []string) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}
	return false
}

func makeMiddleware(predicate func(string, []string) bool) func(...string) func(http.Handler) http.Handler {
	return func(methods ...string) func(http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if ok := predicate(r.Method, methods); !ok {
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
					return
				}
				next.ServeHTTP(w, r)
			})
		}
	}
}

// Allow will only pass through requests to the given handler
// if the request's method is contained in the given list of
// HTTP methods
var Allow = makeMiddleware(contains)

// Disallow will not pass through requests to the given handler
// if the request's method is contained in the given list of
// HTTP methods
var Disallow = makeMiddleware(func(needle string, haystack []string) bool {
	return !contains(needle, haystack)
})
