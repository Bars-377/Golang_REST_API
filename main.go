package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	connStr := "user=postgres password=enigmA1418 dbname=todo_list_db sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handleTasks).Methods("GET", "POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", handleTaskByID).Methods("GET", "PUT", "DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
			return
		}

		task.CreatedAt = time.Now()
		task.UpdatedAt = time.Now()

		err := db.QueryRow(
			`INSERT INTO tasks (title, description, due_date, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5) RETURNING id`,
			task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt,
		).Scan(&task.ID)

		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
		return
	}

	if r.Method == http.MethodGet {
		rows, err := db.Query(
			`SELECT id, title, description, due_date, created_at, updated_at FROM tasks`,
		)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		tasks := []Task{}
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}
}

func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if r.Method == http.MethodGet {
		var task Task
		err := db.QueryRow(
			`SELECT id, title, description, due_date, created_at, updated_at
            FROM tasks WHERE id = $1`,
			id,
		).Scan(
			&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt,
		)

		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
		return
	}

	if r.Method == http.MethodPut {
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
			return
		}

		task.UpdatedAt = time.Now()
		_, err := db.Exec(
			`UPDATE tasks SET title = $1, description = $2, due_date = $3, updated_at = $4 WHERE id = $5`,
			task.Title, task.Description, task.DueDate, task.UpdatedAt, id,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Task not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
		return
	}

	if r.Method == http.MethodDelete {
		_, err := db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Task not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}
}
