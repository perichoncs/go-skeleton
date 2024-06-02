package ports

import "github.com/sebaperi/go-skeleton/internal/domain"

type PlayerService interface {
	Create(player domain.Player) (id interface{}, err error)
}

type PlayerRepository interface {
	Insert(player domain.Player) (id interface{}, err error)
}
