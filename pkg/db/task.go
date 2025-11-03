package db

import (
	"database/sql"
	"strconv"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	// Если глобальная переменная DB ещё не инициализирована (сервер
	// запускают не через main или вызов Init пропущен), то инициализируем БД
	// здесь, чтобы избежать nil-pointer panic при вставке
	if DB() == nil {
		if err := Init(DBFile); err != nil {
			return 0, err
		}
	}
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := DB().Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ListTasks(search string) ([]Task, error) {
	if DB() == nil {
		if err := Init(DBFile); err != nil {
			return nil, err
		}
	}
	var rows *sql.Rows
	var err error
	if search == "" {
		rows, err = DB().Query(`SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC`)
	} else {
		like := "%" + search + "%"
		rows, err = DB().Query(`SELECT id, date, title, comment, repeat FROM scheduler WHERE date LIKE ? OR title LIKE ? OR comment LIKE ? ORDER BY date ASC`, like, like, like)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Task
	for rows.Next() {
		var id int64
		var date, title, comment, repeat string
		if err := rows.Scan(&id, &date, &title, &comment, &repeat); err != nil {
			return nil, err
		}
		out = append(out, Task{ID: strconv.FormatInt(id, 10), Date: date, Title: title, Comment: comment, Repeat: repeat})
	}
	return out, nil
}
