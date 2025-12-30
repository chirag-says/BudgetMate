package main

import (
	"budgetmate/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Public landing page
	mux.HandleFunc("GET /", handlers.HandleLandingPage)

	// Dashboard (App routes)
	mux.HandleFunc("GET /app", handlers.HandleDashboard)
	mux.HandleFunc("GET /app/", handlers.HandleDashboard)

	// Transactions page
	mux.HandleFunc("GET /transactions", handlers.HandleTransactions)

	// HTMX API routes
	mux.HandleFunc("GET /api/transactions/filter", handlers.HandleTransactionsFilter)
	mux.HandleFunc("GET /api/transactions/table", handlers.HandleTransactionsTable)

	// Static file serving (if needed)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Server configuration
	addr := ":8080"
	log.Printf("🚀 BudgetMate Enterprise Dashboard starting on http://localhost%s", addr)
	log.Printf("📊 Access the dashboard at http://localhost%s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
