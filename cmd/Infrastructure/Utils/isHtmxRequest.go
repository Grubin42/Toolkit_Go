package Utils

import (
    "net/http"
)

// IsHtmxRequest vérifie si la requête est faite via htmx
func IsHtmxRequest(r *http.Request) bool {
    return r.Header.Get("Hx-Request") == "true"
}