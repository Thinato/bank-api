
autenticação -> usuario/senha



* client
* manager
* transaction
* stock


estou construindo uma aplicação de um banco digital em GO com um framework chamado gin gonic para API's Rest.

estou seguindo solid design pattern, dependency injection e Domain driven design.

minha estrutura de pastas está assim:

/modules
    /client
        /application
            /controller
                dto.go
                interface.go
                main.go
            /resolver
                dto.go
                interface.go
                main.go
        /domain
            /entities
                client.go
            /usecase
                /login
                    dto.go
                    interface.go
                    main.go
                /signup
                    dto.go
                    interface.go
                    main.go
        /infra
            dto.go
            interface.go
            main.go
        interface.go
        module.go
load_env.go
main.go

preciso da sua ajuda para montar essa aplicação, estou com dificuldade de implementar a conexão com o banco de dados, quero usar o conector do mysql "github.com/go-sql-driver/mysql" preciso também entender onde eu vou instanciar o banco e como vou usa-lo nesse sistema de dependecy injection e domain driven design