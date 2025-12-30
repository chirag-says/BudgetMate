package models

import "time"

// DashboardStats holds the key performance indicators for the dashboard
type DashboardStats struct {
	TotalBalance    float64 `json:"total_balance"`
	BalanceChange   float64 `json:"balance_change"`     // percentage change
	MonthlySpend    float64 `json:"monthly_spend"`
	SpendChange     float64 `json:"spend_change"`       // percentage change
	BudgetRemaining float64 `json:"budget_remaining"`
	BudgetTotal     float64 `json:"budget_total"`
	BudgetUsed      float64 `json:"budget_used"`        // percentage used
	TotalSavings    float64 `json:"total_savings"`
	SavingsChange   float64 `json:"savings_change"`     // percentage change
}

// Transaction represents a single financial transaction
type Transaction struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"` // "credit" or "debit"
	Status      string    `json:"status"` // "completed", "pending", "failed"
	Merchant    string    `json:"merchant"`
}

// ChartDataPoint represents a single point on the spending chart
type ChartDataPoint struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

// SpendingChart holds data for the spending trends chart
type SpendingChart struct {
	Labels []string  `json:"labels"`
	Income []float64 `json:"income"`
	Spend  []float64 `json:"spend"`
}

// DashboardData aggregates all data needed for the dashboard view
type DashboardData struct {
	Stats        DashboardStats  `json:"stats"`
	Transactions []Transaction   `json:"transactions"`
	Chart        SpendingChart   `json:"chart"`
	UserName     string          `json:"user_name"`
	UserEmail    string          `json:"user_email"`
	UserInitials string          `json:"user_initials"`
}

// NavItem represents a navigation menu item
type NavItem struct {
	Label    string `json:"label"`
	Href     string `json:"href"`
	Icon     string `json:"icon"`
	IsActive bool   `json:"is_active"`
}

// BudgetCategory represents a budget category with spending progress
type BudgetCategory struct {
	Name    string  `json:"name"`
	Spent   float64 `json:"spent"`
	Total   float64 `json:"total"`
	Percent float64 `json:"percent"`
}
