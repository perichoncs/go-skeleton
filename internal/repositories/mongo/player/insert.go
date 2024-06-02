package player

import (
	"context"
	"fmt"
	"github.com/sebaperi/go-skeleton/internal/domain"
	"log"
)

func (r Repository) Insert(player domain.Player) (id interface{}, err error) {
	// Insert to the db
	// Respond with the resource

	collection := r.Client.Database("godb").Collection("players")
	insertedResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting player %w", err)
	}

	return insertedResult.InsertedID, nil
}
