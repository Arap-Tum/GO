Good — now we connect everything **properly**, not in a messy way.

You already have:

* ✅ Repository (clean, with context)
* ✅ Basic handler structure

Now we’ll wire them together **like a real backend**.

---

# 🧠 1. Goal (What We’re Doing)

We want this flow:

```text
HTTP Request → Handler → Repository → Database
```

And we’ll make sure:

* Handlers use `context`
* Proper HTTP methods are respected
* Clean routing (no confusion between GET/POST/DELETE)

---

# ⚠️ First — Fix Your Handler Structure

Right now you had separate functions like:

* `CreateExpense`
* `GetExpenses`

👉 That’s fine, but routing becomes messy.

Instead, we use **one handler per route + switch by method**

---

# ✅ 2. Clean Handler (Connected to Repo)

```go
package handlers

import (
	"encoding/json"
	"expenseTracker/internal/models"
	"expenseTracker/internal/repository"
	"net/http"
	"strconv"
)

type ExpenseHandler struct {
	repo *repository.ExpenseRepository
}

func NewExpenseHandler(r *repository.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: r}
}

// Main entry point for /expenses
func (h *ExpenseHandler) HandleExpenses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		h.createExpense(w, r)

	case http.MethodGet:
		h.getExpenses(w, r)

	case http.MethodDelete:
		h.deleteExpense(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
```

---

# 🧱 3. Create Expense

```go
func (h *ExpenseHandler) createExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var expense models.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TEMP: hardcoded user (until auth is added)
	expense.UserID = 1

	err = h.repo.Create(r.Context(), &expense)
	if err != nil {
		http.Error(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}
```

---

# 📊 4. Get Expenses (All or by ID)

```go
func (h *ExpenseHandler) getExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")

	// 👉 If ID is provided → get one
	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		expense, err := h.repo.GetByID(r.Context(), id)
		if err != nil {
			if err == repository.ErrNotFound {
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Error fetching expense", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(expense)
		return
	}

	// 👉 Otherwise → get all
	expenses, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch expenses", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}
```

---

# ❌ 5. Delete Expense

```go
func (h *ExpenseHandler) deleteExpense(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.repo.Delete(r.Context(), id)
	if err != nil {
		if err == repository.ErrNotFound {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
```

---

# 🔌 6. Connect in `main.go`

Now wire everything:

```go
import (
	"expenseTracker/internal/handlers"
	"expenseTracker/internal/repository"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()

	// 🔌 Inject dependencies
	repo := repository.NewExpenseRepository(db)
	expenseHandler := handlers.NewExpenseHandler(repo)

	// 🌐 Routes
	http.HandleFunc("/expenses", expenseHandler.HandleExpenses)
	http.HandleFunc("/health", healthHandler)

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

# 🧪 7. Test with Postman (IMPORTANT)

## ✅ Create Expense

```
POST /expenses
```

```json
{
  "category_id": 1,
  "title": "Lunch",
  "amount": 300
}
```

---

## ✅ Get All

```
GET /expenses
```

---

## ✅ Get One

```
GET /expenses?id=1
```

---

## ✅ Delete

```
DELETE /expenses?id=1
```

---

# 🧠 8. What You Just Built

You now have:

* ✔ Clean dependency injection
* ✔ Context-aware DB calls
* ✔ Proper HTTP method handling
* ✔ Error mapping (404, 500, etc.)

👉 This is **real backend structure**

---

# ⚡ Reality Check

This is no longer beginner code.

Next level would be:

* Proper router (like `gorilla/mux` or `chi`)
* Middleware (JWT)
* Service layer

---

# 🚀 Your Next Move

Don’t rush.

👉 Test everything first.

---

When ready, say:

👉 **“next: clean routing + middleware”**

and I’ll upgrade your system again.
