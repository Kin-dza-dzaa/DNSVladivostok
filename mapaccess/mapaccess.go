package mapacess

import (
	"context"
	"io"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"github.com/Kin-dza-dzaa/DNSVladivostok/conv"
)

const (
	IfKeyInMap   = "SELECT EXISTS(SELECT * FROM map WHERE key = $1);"
	InserIntoMap = "INSERT INTO map VALUES($1, $2);"
	UpdateMap    = "UPDATE map SET value = $1 WHERE key = $2;"
)

type MapAccess struct {
	Logger *zerolog.Logger
	Pool   *pgxpool.Pool
	File   io.Reader
}

func (m *MapAccess) InsertOne(row string) error {
	key, value, err := conv.ParseRow(row)
	if err != nil {
		return err
	}
	return m.Insert(key, value)
}

func (m *MapAccess) Insert(key, value int) error {
	return m.Pool.BeginFunc(context.TODO(), func(tx pgx.Tx) error {
		exists, err := check(key, tx)
		if err != nil {
			return err
		}
		if exists {
			if _, err := tx.Exec(context.TODO(), UpdateMap, key, value); err != nil {
				return err
			} else {
				m.Logger.Info().Msg(fmt.Sprintf("value on key %d was updated to %d", key, value))
			}
		} else {
			if _, err := tx.Exec(context.TODO(), InserIntoMap, key, value); err != nil {
				return err
			} else {
				m.Logger.Info().Msg(fmt.Sprintf("value %d on key %d was inserted into table", value, key))
			}
		}
		return nil
	})
}

func NewMapAccess(logger *zerolog.Logger, pool *pgxpool.Pool, file io.Reader) *MapAccess {
	return &MapAccess{
		Logger: logger,
		Pool:   pool,
		File:   file,
	}
}

func check(key int, tx pgx.Tx) (bool, error) {
	res := false
	if err := tx.QueryRow(context.TODO(), IfKeyInMap, key).Scan(&res); err != nil {
		return false, err
	}
	return res, nil
}
