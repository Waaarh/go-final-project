package api

import (
	"encoding/json"
	"fmt"
	"go1f/pkg/dateutils"
	"go1f/pkg/db"
	_ "modernc.org/sqlite"
	"net/http"
	"time"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format("20060102")
	}
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return err
	}
	var next string

	if task.Repeat != "" {
		next, err = dateutils.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}
	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format("20060102")
		} else {
			task.Date = next
		}
	}
	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---- Incoming request ----")
	defer func() {
		fmt.Println("---- END of request ----")
	}()

	var task db.Task
	fmt.Printf("Response JSON: %+v\n", map[string]any{"id": task.ID})

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"id": id})
}
