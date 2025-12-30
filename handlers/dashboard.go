package handlers

import (
	"budgetmate/models"
	"budgetmate/templates"
	"net/http"
	"strings"
	"time"
)

// getMockDashboardData returns mock data for the dashboard (Indian market)
func getMockDashboardData() models.DashboardData {
	return models.DashboardData{
		UserName:     "Priya Sharma",
		UserEmail:    "priya.sharma@company.in",
		UserInitials: "PS",
		Stats: models.DashboardStats{
			TotalBalance:    2847504.25,    // ₹28.47 Lakh
			BalanceChange:   12.5,
			MonthlySpend:    184321.50,     // ₹1.84 Lakh
			SpendChange:     -8.3,
			BudgetRemaining: 65678.50,      // ₹65,678.50
			BudgetTotal:     250000.00,     // ₹2.5 Lakh
			BudgetUsed:      73.7,
			TotalSavings:    452800.00,     // ₹4.52 Lakh
			SavingsChange:   15.2,
		},
		Transactions: []models.Transaction{
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
				Description: "Food Delivery",
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
				Description: "School Tuition - Q4",
				Category:    "Tuition",
				Amount:      45000.00,
				Type:        "debit",
				Status:      "completed",
				Merchant:    "DPS Bangalore",
			},
			{
				ID:          "TXN-008",
				Date:        time.Now().AddDate(0, 0, -5),
				Description: "SIP Investment",
				Category:    "Investment",
				Amount:      25000.00,
				Type:        "debit",
				Status:      "completed",
				Merchant:    "Zerodha",
			},
			{
				ID:          "TXN-009",
				Date:        time.Now().AddDate(0, 0, -6),
				Description: "Petrol",
				Category:    "Transport",
				Amount:      4500.00,
				Type:        "debit",
				Status:      "completed",
				Merchant:    "HP Petrol Pump",
			},
			{
				ID:          "TXN-010",
				Date:        time.Now().AddDate(0, 0, -7),
				Description: "Freelance Payment",
				Category:    "Income",
				Amount:      50000.00,
				Type:        "credit",
				Status:      "completed",
				Merchant:    "Upwork",
			},
		},
		Chart: models.SpendingChart{
			Labels: []string{"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
			Income: []float64{185000, 192000, 185000, 235000, 190000, 235000},
			Spend:  []float64{142000, 156000, 138000, 175000, 148000, 165000},
		},
	}
}

// getBudgetCategories returns budget categories with spending progress
func getBudgetCategories() []models.BudgetCategory {
	return []models.BudgetCategory{
		{Name: "Groceries & Household", Spent: 15200, Total: 20000, Percent: 76},
		{Name: "Utilities & Bills", Spent: 8500, Total: 15000, Percent: 56.7},
		{Name: "Food & Dining", Spent: 12400, Total: 15000, Percent: 82.7},
		{Name: "Transport", Spent: 9200, Total: 12000, Percent: 76.7},
		{Name: "Entertainment", Spent: 4800, Total: 8000, Percent: 60},
	}
}

// getNavItems returns the navigation items for the sidebar
func getNavItems(activePath string) []models.NavItem {
	items := []models.NavItem{
		{Label: "Dashboard", Href: "/", Icon: "dashboard"},
		{Label: "Transactions", Href: "/transactions", Icon: "transactions"},
		{Label: "Reports", Href: "/reports", Icon: "reports"},
		{Label: "Budgets", Href: "/budgets", Icon: "budgets"},
		{Label: "Settings", Href: "/settings", Icon: "settings"},
	}

	for i := range items {
		if items[i].Href == activePath {
			items[i].IsActive = true
		}
	}

	return items
}

// HandleDashboard handles the main dashboard page request
func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	data := getMockDashboardData()
	navItems := getNavItems("/")
	budgets := getBudgetCategories()

	component := templates.Dashboard(data, navItems, budgets)
	component.Render(r.Context(), w)
}

// HandleTransactionsFilter handles HTMX requests to filter transactions
func HandleTransactionsFilter(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("search"))
	data := getMockDashboardData()

	// Filter transactions based on search query
	var filtered []models.Transaction
	for _, txn := range data.Transactions {
		if query == "" ||
			strings.Contains(strings.ToLower(txn.Description), query) ||
			strings.Contains(strings.ToLower(txn.Category), query) ||
			strings.Contains(strings.ToLower(txn.Merchant), query) {
			filtered = append(filtered, txn)
		}
	}

	component := templates.TransactionRows(filtered)
	component.Render(r.Context(), w)
}
