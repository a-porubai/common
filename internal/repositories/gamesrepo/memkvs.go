package gamesrepo

import (
	"encoding/json"
	"errors"

	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/domain"
)

type Memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *Memkvs {
	return &Memkvs{kvs: map[string][]byte{}}
}

func (repo *Memkvs) Get(id string) (*domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := &domain.Game{}
		err := json.Unmarshal(value, &game)

		if err != nil {
			return nil, errors.New("fail to get value from kvs")
		}

		return game, nil
	}

	return nil, errors.New("game not found in kvs")
}

func (repo *Memkvs) Save(game domain.Game) error {
	bytes, err := json.Marshal(game)
	if err != nil {
		return errors.New("game fails at marshal into json string")
	}

	repo.kvs[game.ID] = bytes

	return nil
}
