package handlers

import (
	"budgetmate/templates"
	"net/http"
)

// HandleLandingPage handles the public landing page
func HandleLandingPage(w http.ResponseWriter, r *http.Request) {
	component := templates.LandingPage()
	component.Render(r.Context(), w)
}
