# go-todo

A simple todo utility/server written in golang. Run the todo script under **cmd/todo**.
Also a minimal JSON API is also available at **cmd/todo-server**, which will be written in [gin](https://github.com/gin-gonic/gin) framework.

[Dep](https://github.com/golang/dep) is used for dependancy management.

I would recommend to use atom|go-plus combo to build/hack this if you are new to golang!

## Setup

1. Clone this project
2. Install dep, if not, and run
```
$ dep ensure
```
## Build
### todo utility demo
1. `$ go build github.com/vishrayne/go-todo/cmd/todo`
2. `$ go install github.com/vishrayne/go-todo/cmd/todo`
3. `$ todo`

### todo-server demo
1. `$ go build github.com/vishrayne/go-todo/cmd/todo-server`
2. `$ go install github.com/vishrayne/go-todo/cmd/todo-server`
3. `$ todo-server`
4. visit localhost:8080 on your browser
5. CRUD endpoints are available under `/api/v1/todos`, check your terminal for more details.


TODO: clean up around todo-server
