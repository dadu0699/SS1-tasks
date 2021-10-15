package app

import "net/http"

// IndexHandler handler that works like healthcheck.
func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendResponse(w, r, "API running smoothly", http.StatusOK)
	}
}
