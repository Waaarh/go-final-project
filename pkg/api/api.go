package api

import (
	"go1f/pkg/db"
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", NextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	// Минимальный обработчик для получения списка задач он типо нужен фронтенду
	// он запрашивает /api/tasks после добавления, чтобы обновить UI
	http.HandleFunc("/api/tasks", tasksHandler)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем список задач в простом формате, ожидаемом фронтендом/tests
	// Здесь не делается сложной фильтрации, только простой LIKE по полям
	if r.Method != http.MethodGet {
		writeJSON(w, map[string]string{"error": "method not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	search := r.URL.Query().Get("search")
	tasks, err := db.ListTasks(search)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	out := make([]map[string]string, 0, len(tasks))
	for _, t := range tasks {
		out = append(out, map[string]string{
			"id":      t.ID,
			"date":    t.Date,
			"title":   t.Title,
			"comment": t.Comment,
			"repeat":  t.Repeat,
		})
	}
	writeJSON(w, map[string]any{"tasks": out}, http.StatusOK)
}
