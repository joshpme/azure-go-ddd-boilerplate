# Azure DDD Boilerplate Go

This is a boilerplate for a DDD / CQRS / Event Sourcing project in GoLang. 

It also includes boilerplate for it to be used as an Azure Function.

This is my first attempt at this in Go, so I'm sure there are things that can be improved. Feel free to open an issue or PR.

During development, use `go run cmd/main.go` and debug with CURL by hitting the routes in `cmd/main.go`.

If you are debugging an azure func, first compile into app with `go build -o app cmd/main.go` and then run `func start` to start the func host.