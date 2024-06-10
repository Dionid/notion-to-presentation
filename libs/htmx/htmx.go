package htmx

import "net/http"

func isHTMX(r *http.Request) bool {
	return r.Header.Get("HX-Request") != ""
}
