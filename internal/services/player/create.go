package player

import (
	"fmt"
	"github.com/sebaperi/go-skeleton/internal/domain"
	"log"
	"time"
)

func (s Service) Create(player domain.Player) (id interface{}, err error) {
	// Set creation time
	// Save to repo
	// Response with created id resource
	player.CreationTime = time.Now().UTC()

	insertedId, err := s.Repo.Insert(player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating player %w", err)
	}
	return insertedId, nil
}
