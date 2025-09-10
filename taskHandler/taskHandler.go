package taskhandler

import (
	"fmt"
	"strings"
)

type toDo struct {
    Id          int
    Description string
}

type toDoList struct {
    todos  map[int]toDo
    nextId int
}

http.HandleFunc("/todos", handleTodos)

func (l *toDoList) HandleTodos(w http.ResponseWriter, r *http.Request) {
    switch r.method {
	case http.methodGet:
		var list []toDo
		for _, d : range l.todos {
			list.append(list, d)
		}
		data, _ := json.Marshal(list)
		w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(data)
	case http.methodPost:
	case http.methodPut:
	case http.methodDelete:
	case http.methodOptions:

	}
}