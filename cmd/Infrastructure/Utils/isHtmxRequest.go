package Utils

import (
    "net/http"
    "strings"
)

// IsHtmxRequest vérifie si la requête est faite via htmx
func IsHtmxRequest(r *http.Request) bool {
    return strings.ToLower(r.Header.Get("HX-Request")) == "true"
}