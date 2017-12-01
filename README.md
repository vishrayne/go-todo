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
3. go build `github.com/vishrayne/go-todo/cmd/todo` or `github.com/vishrayne/go-todo/cmd/todo-server`
4. go install `<which ever component you have built>`
5. `$ todo <or> todo-server`


TODO: clean up around todo-server
