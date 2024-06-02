# go-skeleton

File structure
- cmd
    - api
        - handlers
            - entity
                CRUD.go
        main.go
        errors.go
- internal
  - domain
    - entity.go
  - services
    - entity
      CRUD.go
      service.go
  - repositories
    - mongo
      - player
        insert.go
        repository.go
      connectClient.go
  - ports
    player.go
  - validator

Create a .env file with the following:
MONGO_URI="mongodb-uri"

cmd:
api:
internal: