package handlers

import (
	"budgetmate/models"
	"budgetmate/templates"
	"net/http"
	"strings"
	"time"
)

// getMockDashboardData returns mock data for the dashboard
func getMockDashboardData() models.DashboardData {
	return models.DashboardData{
		UserName:     "Sarah Mitchell",
		UserEmail:    "sarah.mitchell@company.com",
		UserInitials: "SM",
		Stats: models.DashboardStats{
			TotalBalance:    284750.42,
			BalanceChange:   12.5,
			MonthlySpend:    18432.15,
			SpendChange:     -8.3,
			BudgetRemaining: 6567.85,
			BudgetTotal:     25000.00,
			BudgetUsed:      73.7,
			TotalSavings:    45280.00,
			SavingsChange:   15.2,
		},
		Transactions: []models.Transaction{
			{
				ID:          "TXN-001",
				Date:        time.Now().AddDate(0, 0, -1),
				Description: "Office Supplies - Q4 Restock",
				Category:    "Operations",
				Amount:      1240.50,
				Type:        "debit",
				Status:      "completed",
				Merchant:    "Staples Business",
			},
			{
				ID:          "TXN-002",
				Date:        time.Now().AddDate(0, 0, -1),
				Description: "Client Payment - Invoice #4521",
				Category:    "Revenue",
				Amount:      15750.00,
				Type:        "credit",
				Status:      "completed",
				Merchant:    "Acme Corporation",
			},
			{
				ID:          "TXN-003",
				Date:        time.Now().AddDate(0, 0, -2),
				Description: "Software Subscription - Annual",
				Category:    "Technology",
				Amount:      2999.00,
				Type:        "debit",
				Status:      "completed",
				Merchant:    "Atlassian",
			},
			{
				ID:          "TXN-004",
				Date:        time.Now().AddDate(0, 0, -3),
				Description: "Travel Expenses - NYC Conference",
				Category:    "Travel",
				Amount:      3420.85,
				Type:        "debit",
				Status:      "pending",
				Merchant:    "Delta Airlines",
			},
			{
				ID:          "TXN-005",
				Date:        time.Now().AddDate(0, 0, -3),
				Description: "Consulting Fee - December",
				Category:    "Revenue",
				Amount:      8500.00,
				Type:        "credit",
				Status:      "completed",
				Merchant:    "Tech Solutions LLC",
			},
			{
				ID:          "TXN-006",
				Date:        time.Now().AddDate(0, 0, -4),
				Description: "Marketing Campaign - Social Ads",
				Category:    "Marketing",
				Amount:      4250.00,
				Type:        "debit",
				Status:      "completed",
				Merchant:    "Meta Platforms",
			},
			{
				ID:          "TXN-007",
				Date:        time.Now().AddDate(0, 0, -5),
				Description: "Equipment Lease Payment",
				Category:    "Operations",
				Amount:      1875.00,
				Type:        "debit",
				Status:      "completed",
				Merchant:    "Dell Financial",
			},
			{
				ID:          "TXN-008",
				Date:        time.Now().AddDate(0, 0, -5),
				Description: "Wire Transfer - Partner Payout",
				Category:    "Payroll",
				Amount:      12000.00,
				Type:        "debit",
				Status:      "completed",
				Merchant:    "Internal Transfer",
			},
		},
		Chart: models.SpendingChart{
			Labels: []string{"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
			Income: []float64{42000, 38500, 45200, 52100, 48750, 55420},
			Spend:  []float64{28500, 31200, 27800, 35400, 29100, 32150},
		},
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

	component := templates.Dashboard(data, navItems)
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
