# BudgetMate Enterprise Dashboard

A production-grade, high-contrast financial dashboard built with Go, Templ, HTMX, and Alpine.js.

## Tech Stack

- **Go 1.23+** - Backend server with `net/http`
- **Templ** - Type-safe HTML templating
- **HTMX** - Dynamic HTML interactions without full page reloads
- **Alpine.js** - Lightweight JavaScript for micro-interactions
- **Chart.js** - Clean, flat data visualizations
- **Tailwind CSS** - Utility-first styling (via CDN)

## Design Philosophy

**"Institutional Finance"** - Crisp, dense, and trustworthy:
- Deep Navy Blue sidebar (`bg-slate-900`)
- Pure White and Slate backgrounds (`bg-white`, `bg-slate-50`)
- Emerald Green for positive financials (`text-emerald-600`)
- Thin, subtle borders (`border-slate-200`)
- Sharp, minimal shadows (`shadow-sm`)
- Standard border radius (4px-6px)

## Getting Started

### Prerequisites

- Go 1.23 or later
- Templ CLI (`go install github.com/a-h/templ/cmd/templ@latest`)

### Installation

1. Clone/enter the project directory:
   ```bash
   cd BudgetMate
   ```

2. Download dependencies:
   ```bash
   go mod tidy
   ```

3. Generate Templ files:
   ```bash
   templ generate
   ```

4. Run the server:
   ```bash
   go run .
   ```

5. Open your browser to [http://localhost:8080](http://localhost:8080)

## Project Structure

```
BudgetMate/
├── main.go                 # HTTP server & routes
├── go.mod                  # Go module definition
├── handlers/
│   └── dashboard.go        # Request handlers & mock data
├── models/
│   └── models.go           # Data structures
├── templates/
│   ├── base.templ          # Layout with sidebar & header
│   └── dashboard.templ     # Dashboard page components
└── README.md
```

## Features

### KPI Cards
4 high-density metric cards displaying:
- Total Balance (with change indicator)
- Monthly Spend (with trend)
- Budget Used (percentage progress)
- Total Savings (with growth)

### Cash Flow Chart
Clean line chart showing income vs expenses over time using Chart.js.

### Budget Status Panel
Visual progress bars for budget category utilization with color-coded thresholds:
- Green: Under 75%
- Amber: 75-90%
- Red: Over 90%

### Transactions Table
Data-dense table with:
- HTMX-powered live filtering (no page refresh)
- Status badges (completed, pending, failed)
- Credit/debit color coding
- Sortable columns

### Responsive Layout
- Collapsible sidebar on mobile
- Alpine.js-powered mobile menu toggle
- Responsive grid layouts

## HTMX Integration

The transactions table uses HTMX for seamless filtering:

```html
<input 
    hx-get="/api/transactions/filter"
    hx-trigger="keyup changed delay:300ms"
    hx-target="#transaction-rows"
    hx-swap="innerHTML"
/>
```

## Alpine.js Usage

Mobile menu toggle and dropdown menus:

```html
<div x-data="{ sidebarOpen: false }">
    <button @click="sidebarOpen = true">Menu</button>
    <aside x-show="sidebarOpen">...</aside>
</div>
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/` | Main dashboard |
| GET | `/api/transactions/filter?search=query` | Filter transactions |

## Development

### Hot Reload with Templ

For development with hot reload:

```bash
# Terminal 1: Watch templ files
templ generate --watch

# Terminal 2: Run with air (optional)
air
```

### Adding New Pages

1. Create a new handler in `handlers/`
2. Create a new template in `templates/`
3. Register the route in `main.go`

## License

MIT License
