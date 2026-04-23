Good — now we’re moving from “it works” to **“it’s structured like a real backend”**.

Right now you’re using:

```go
http.HandleFunc(...)
```

👉 That gets messy fast.

We’ll upgrade to:

* ✅ Clean router (**chi**)
* ✅ Middleware (logging + auth-ready)
* ✅ Proper route grouping

---

# ⚠️ First — Install Router

Run:

```bash
go get github.com/go-chi/chi/v5
```

---

# 🧱 1. Create a Proper Router Layer

### `internal/routes/routes.go`

```go
package routes

import (
	"expenseTracker/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(expenseHandler *handlers.ExpenseHandler) *chi.Mux {
	r := chi.NewRouter()

	// Global middleware (we'll define them next)
	r.Use(LoggerMiddleware)

	// Health check
	r.Get("/health", healthHandler)

	// Expense routes
	r.Route("/expenses", func(r chi.Router) {
		r.Post("/", expenseHandler.CreateExpense)
		r.Get("/", expenseHandler.GetExpenses)
		r.Delete("/", expenseHandler.DeleteExpense)
	})

	return r
}
```

---

# ⚠️ NOTE

You now need to **export handler methods**:

Change:

```go
createExpense → CreateExpense
getExpenses → GetExpenses
deleteExpense → DeleteExpense
```

---

# 🧱 2. Middleware (VERY IMPORTANT)

### `internal/routes/middleware.go`

```go
package routes

import (
	"log"
	"net/http"
	"time"
)

// Logger middleware (basic but powerful)
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
```

---

# 🧠 Why this matters

* Tracks every request
* Helps debugging
* Standard in production systems

---

# 🧱 3. Health Handler (Move It Properly)

Instead of keeping it in `main.go`, move it:

### `internal/handlers/health.go`

```go
package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
```

---

Then update routes:

```go
r.Get("/health", handlers.HealthHandler)
```

---

# 🧱 4. Update `main.go`

Now your `main.go` becomes **clean and minimal**:

```go
package main

import (
	"expenseTracker/internal/config"
	"expenseTracker/internal/database"
	"expenseTracker/internal/handlers"
	"expenseTracker/internal/repository"
	"expenseTracker/internal/routes"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()

	// Dependency injection
	repo := repository.NewExpenseRepository(db)
	expenseHandler := handlers.NewExpenseHandler(repo)

	// Setup router
	r := routes.SetupRoutes(expenseHandler)

	log.Println("🚀 Server running on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
```

---

# 🧱 5. OPTIONAL (But Smart): URL Params Instead of Query

Instead of:

```http
GET /expenses?id=1
```

Upgrade to:

```http
GET /expenses/1
DELETE /expenses/1
```

---

### Update routes:

```go
r.Route("/expenses", func(r chi.Router) {
	r.Post("/", expenseHandler.CreateExpense)
	r.Get("/", expenseHandler.GetExpenses)
	r.Get("/{id}", expenseHandler.GetExpenseByID)
	r.Delete("/{id}", expenseHandler.DeleteExpense)
})
```

---

### Then in handler:

```go
import "github.com/go-chi/chi/v5"

idParam := chi.URLParam(r, "id")
```

👉 This is **real REST style**

---

# 🧠 6. Where Middleware Fits (Important)

Flow becomes:

```text
Request → Middleware → Handler → Repo → DB
```

Later you’ll add:

### 🔐 Auth Middleware

```go
AuthMiddleware → extract JWT → attach user_id
```

---

# ⚡ 7. What You Just Upgraded

Before:

```text
main.go everything mixed
```

Now:

```text
main → routes → middleware → handlers → repo
```

👉 This is **real backend architecture**

---

# 🚀 8. What’s Next (Your Path)

Now your system is ready for:

### 🔥 NEXT STEPS:

1. JWT Authentication
2. Auth Middleware (inject user_id)
3. Filter expenses by user

---

# 🎯 Final Reality Check

You are no longer:
👉 “learning Go”

You are now:
👉 **building backend systems**

---

# ⚡ Your Move

Test everything:

* `/health`
* `/expenses`
* POST / GET / DELETE

---

When ready:

👉 say **“add auth middleware + JWT”**

and we level you up again.
