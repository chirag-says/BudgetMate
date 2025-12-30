package handlers

import (
	"budgetmate/models"
	"budgetmate/templates"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// getAllTransactions returns a comprehensive list of mock transactions for India
func getAllTransactions() []models.Transaction {
	return []models.Transaction{
		{
			ID:          "TXN-001",
			Date:        time.Now().AddDate(0, 0, -1),
			Description: "Monthly Groceries",
			Category:    "Groceries",
			Amount:      8542.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "BigBasket",
		},
		{
			ID:          "TXN-002",
			Date:        time.Now().AddDate(0, 0, -1),
			Description: "Salary Credit - December",
			Category:    "Income",
			Amount:      185000.00,
			Type:        "credit",
			Status:      "completed",
			Merchant:    "TechMahindra Ltd",
		},
		{
			ID:          "TXN-003",
			Date:        time.Now().AddDate(0, 0, -2),
			Description: "Electricity Bill - December",
			Category:    "Utilities",
			Amount:      3420.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Tata Power",
		},
		{
			ID:          "TXN-004",
			Date:        time.Now().AddDate(0, 0, -3),
			Description: "House Rent - January",
			Category:    "Rent",
			Amount:      35000.00,
			Type:        "debit",
			Status:      "pending",
			Merchant:    "Landlord Transfer",
		},
		{
			ID:          "TXN-005",
			Date:        time.Now().AddDate(0, 0, -3),
			Description: "Food Delivery - Dinner",
			Category:    "Food & Dining",
			Amount:      1250.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Zomato",
		},
		{
			ID:          "TXN-006",
			Date:        time.Now().AddDate(0, 0, -4),
			Description: "Mobile Recharge",
			Category:    "Utilities",
			Amount:      599.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Amazon Pay",
		},
		{
			ID:          "TXN-007",
			Date:        time.Now().AddDate(0, 0, -5),
			Description: "School Tuition Fee - Q4",
			Category:    "Tuition",
			Amount:      45000.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "DPS Bangalore",
		},
		{
			ID:          "TXN-008",
			Date:        time.Now().AddDate(0, 0, -5),
			Description: "SIP Investment - Monthly",
			Category:    "Investment",
			Amount:      25000.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Zerodha",
		},
		{
			ID:          "TXN-009",
			Date:        time.Now().AddDate(0, 0, -6),
			Description: "Petrol Refill",
			Category:    "Transport",
			Amount:      4500.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "HP Petrol Pump",
		},
		{
			ID:          "TXN-010",
			Date:        time.Now().AddDate(0, 0, -7),
			Description: "Freelance Payment Received",
			Category:    "Income",
			Amount:      50000.00,
			Type:        "credit",
			Status:      "completed",
			Merchant:    "Upwork",
		},
		{
			ID:          "TXN-011",
			Date:        time.Now().AddDate(0, 0, -8),
			Description: "Home Loan EMI - December",
			Category:    "EMI",
			Amount:      42500.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "HDFC Bank",
		},
		{
			ID:          "TXN-012",
			Date:        time.Now().AddDate(0, 0, -9),
			Description: "DTH Recharge - Annual",
			Category:    "Utilities",
			Amount:      4999.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Tata Play",
		},
		{
			ID:          "TXN-013",
			Date:        time.Now().AddDate(0, 0, -10),
			Description: "Online Shopping - Electronics",
			Category:    "Shopping",
			Amount:      15999.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Amazon India",
		},
		{
			ID:          "TXN-014",
			Date:        time.Now().AddDate(0, 0, -10),
			Description: "Netflix Subscription",
			Category:    "Entertainment",
			Amount:      649.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Netflix",
		},
		{
			ID:          "TXN-015",
			Date:        time.Now().AddDate(0, 0, -11),
			Description: "LIC Premium - Annual",
			Category:    "Insurance",
			Amount:      28000.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "LIC India",
		},
		{
			ID:          "TXN-016",
			Date:        time.Now().AddDate(0, 0, -12),
			Description: "Cab Ride - Airport",
			Category:    "Transport",
			Amount:      1850.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Uber India",
		},
		{
			ID:          "TXN-017",
			Date:        time.Now().AddDate(0, 0, -13),
			Description: "Medicine Purchase",
			Category:    "Healthcare",
			Amount:      2340.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Apollo Pharmacy",
		},
		{
			ID:          "TXN-018",
			Date:        time.Now().AddDate(0, 0, -14),
			Description: "Dividend Credit - Stocks",
			Category:    "Income",
			Amount:      8500.00,
			Type:        "credit",
			Status:      "completed",
			Merchant:    "Zerodha",
		},
		{
			ID:          "TXN-019",
			Date:        time.Now().AddDate(0, 0, -15),
			Description: "Gas Cylinder Refill",
			Category:    "Utilities",
			Amount:      1103.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Indane Gas",
		},
		{
			ID:          "TXN-020",
			Date:        time.Now().AddDate(0, 0, -16),
			Description: "Swiggy Order - Lunch",
			Category:    "Food & Dining",
			Amount:      450.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Swiggy",
		},
		{
			ID:          "TXN-021",
			Date:        time.Now().AddDate(0, 0, -17),
			Description: "Credit Card Payment",
			Category:    "EMI",
			Amount:      15000.00,
			Type:        "debit",
			Status:      "pending",
			Merchant:    "ICICI Bank",
		},
		{
			ID:          "TXN-022",
			Date:        time.Now().AddDate(0, 0, -18),
			Description: "Gym Membership - Quarterly",
			Category:    "Healthcare",
			Amount:      7500.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Cult.fit",
		},
		{
			ID:          "TXN-023",
			Date:        time.Now().AddDate(0, 0, -19),
			Description: "Broadband Bill - Monthly",
			Category:    "Utilities",
			Amount:      1499.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Jio Fiber",
		},
		{
			ID:          "TXN-024",
			Date:        time.Now().AddDate(0, 0, -20),
			Description: "Birthday Gift Purchase",
			Category:    "Shopping",
			Amount:      5200.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "Flipkart",
		},
		{
			ID:          "TXN-025",
			Date:        time.Now().AddDate(0, 0, -21),
			Description: "Water Bill - Quarterly",
			Category:    "Utilities",
			Amount:      850.00,
			Type:        "debit",
			Status:      "completed",
			Merchant:    "BWSSB",
		},
	}
}

// getCategories returns all available transaction categories
func getCategories() []string {
	return []string{
		"All Categories",
		"Income",
		"Groceries",
		"Utilities",
		"Rent",
		"Food & Dining",
		"Transport",
		"Tuition",
		"Investment",
		"EMI",
		"Shopping",
		"Entertainment",
		"Insurance",
		"Healthcare",
	}
}

// HandleTransactions handles the transactions page request with filtering and pagination
func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	searchQuery := strings.ToLower(r.URL.Query().Get("search"))
	categoryFilter := r.URL.Query().Get("category")
	pageStr := r.URL.Query().Get("page")
	
	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// Get all transactions and filter
	allTransactions := getAllTransactions()
	filtered := filterTransactions(allTransactions, searchQuery, categoryFilter)

	// Pagination
	perPage := 10
	totalItems := len(filtered)
	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	if totalPages == 0 {
		totalPages = 1
	}
	if page > totalPages {
		page = totalPages
	}

	startIdx := (page - 1) * perPage
	endIdx := startIdx + perPage
	if endIdx > totalItems {
		endIdx = totalItems
	}

	var paginatedTransactions []models.Transaction
	if startIdx < totalItems {
		paginatedTransactions = filtered[startIdx:endIdx]
	}

	// Prepare page data
	pageData := models.TransactionsPageData{
		Transactions:   paginatedTransactions,
		Categories:     getCategories(),
		CurrentPage:    page,
		TotalPages:     totalPages,
		TotalItems:     totalItems,
		PerPage:        perPage,
		SearchQuery:    searchQuery,
		CategoryFilter: categoryFilter,
	}

	navItems := getNavItems("/transactions")
	userData := models.DashboardData{
		UserName:     "Priya Sharma",
		UserEmail:    "priya.sharma@company.in",
		UserInitials: "PS",
	}

	component := templates.TransactionsPage(pageData, navItems, userData)
	component.Render(r.Context(), w)
}

// HandleTransactionsTable handles HTMX requests to filter/paginate just the table
func HandleTransactionsTable(w http.ResponseWriter, r *http.Request) {
	searchQuery := strings.ToLower(r.URL.Query().Get("search"))
	categoryFilter := r.URL.Query().Get("category")
	pageStr := r.URL.Query().Get("page")

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	allTransactions := getAllTransactions()
	filtered := filterTransactions(allTransactions, searchQuery, categoryFilter)

	perPage := 10
	totalItems := len(filtered)
	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	if totalPages == 0 {
		totalPages = 1
	}
	if page > totalPages {
		page = totalPages
	}

	startIdx := (page - 1) * perPage
	endIdx := startIdx + perPage
	if endIdx > totalItems {
		endIdx = totalItems
	}

	var paginatedTransactions []models.Transaction
	if startIdx < totalItems {
		paginatedTransactions = filtered[startIdx:endIdx]
	}

	pageData := models.TransactionsPageData{
		Transactions:   paginatedTransactions,
		CurrentPage:    page,
		TotalPages:     totalPages,
		TotalItems:     totalItems,
		PerPage:        perPage,
		SearchQuery:    searchQuery,
		CategoryFilter: categoryFilter,
	}

	component := templates.TransactionsTableContent(pageData)
	component.Render(r.Context(), w)
}

// filterTransactions filters transactions based on search query and category
func filterTransactions(transactions []models.Transaction, search, category string) []models.Transaction {
	var filtered []models.Transaction

	for _, txn := range transactions {
		// Category filter
		if category != "" && category != "All Categories" {
			if !strings.EqualFold(txn.Category, category) {
				continue
			}
		}

		// Search filter
		if search != "" {
			searchLower := strings.ToLower(search)
			if !strings.Contains(strings.ToLower(txn.Description), searchLower) &&
				!strings.Contains(strings.ToLower(txn.Category), searchLower) &&
				!strings.Contains(strings.ToLower(txn.Merchant), searchLower) &&
				!strings.Contains(strings.ToLower(txn.ID), searchLower) {
				continue
			}
		}

		filtered = append(filtered, txn)
	}

	return filtered
}

// buildFilterURL builds URL with current filter params
func buildFilterURL(search, category string, page int) string {
	params := []string{}
	if search != "" {
		params = append(params, fmt.Sprintf("search=%s", search))
	}
	if category != "" && category != "All Categories" {
		params = append(params, fmt.Sprintf("category=%s", category))
	}
	params = append(params, fmt.Sprintf("page=%d", page))
	
	return "/api/transactions/table?" + strings.Join(params, "&")
}
