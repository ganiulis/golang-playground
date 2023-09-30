# go-ticket

Everything is a draft. It always was; it always will.

```sh
docker-compose up -d --build

docker exec -it api /bin/zsh

# Inside the shell.

go run ./main.go
```

I used Gin and GORM to set up the initial codebase. Both are the most popular Go solutions to common infrastructure/ORM problems.

A lot of the code was originally taken from [this article](https://tanmay-vaish.hashnode.dev/how-to-implement-authentication-and-authorization-in-golang) by Tanmay Vaish.

Tests are done thanks to [the manual](https://go.dev/doc/tutorial/add-a-test) and [TableDrivenTests](https://github.com/golang/go/wiki/TableDrivenTests).

The project layout is still unfinished. Current inspiration can be found in [this repository](https://github.com/golang-standards/project-layout/tree/master).

## Sources
1. https://hub.docker.com/_/golang
2. https://github.com/golang-standards/project-layout
3. https://docs.docker.com/language/golang/build-images/
