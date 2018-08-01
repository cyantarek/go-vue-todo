package main

import (
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	rows, _ := Db.Query("select * from tasks")
	defer rows.Close()
	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Name)
		tasks = append(tasks, task)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func PostTasks(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	result, _ := Db.Exec("insert into tasks(name) values(?)", task.Name)
	id, _ := result.LastInsertId()
	task.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	Db.Exec("delete from tasks where id = ?", id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "DELETED Tasks " + id})
}
