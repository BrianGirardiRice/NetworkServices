package taskHandler

import (
	"encoding/json"
	"net/http"
	"sync"
)

type toDo struct {
	Id          int
	Description string
}

type toDoList struct {
	mu   sync.Mutex
	list map[int]string
}

func NewServer() http.Handler {
	todos := toDoList{list: make(map[int]string)}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", todos.getAllToDos)
	// mux.HandleFunc("POST /todos", todos.createToDo)
	// mux.HandleFunc("OPTIONS /todos", todos.optionsAllToDos)
	// mux.HandleFunc("GET /todos/", todos.getToDo)
	// mux.HandleFunc("PUT /todos/", todos.createOrReplaceToDo)
	// mux.HandleFunc("DELETE /todos/", todos.deleteToDo)
	// mux.HandleFunc("OPTIONS /todos/", todos.optionsToDo)
	return mux
}

func (l *toDoList) getAllToDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	l.mu.Lock()
	defer l.mu.Unlock()
	todos := make([]toDo, 0, len(l.list))

	for id, d := range l.list {
		todos = append(todos, toDo{Id: id, Description: d})
	}
	data, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
