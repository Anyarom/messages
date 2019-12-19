package postgres

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/rs/zerolog"
	"time"
	"worker/config"
)

// структура для создания соединения c бд
type PgxClient struct {
	PgxConnPool *pgx.ConnPool
	Log         zerolog.Logger
}

// создаем клиента Pg
func InitPgxClient(pgConfig config.PgConfig, log zerolog.Logger) (*PgxClient, error) {

	// заполняем конфиг (с помощью Вайпер)
	connConfig := pgx.ConnConfig{
		Host:     pgConfig.Host,
		Port:     uint16(pgConfig.Port),
		Database: pgConfig.Database,
		User:     pgConfig.User,
		Password: pgConfig.Password,
	}

	// создаем соединение
	pgxConnPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 10})
	if err != nil {
		return nil, err
	}

	return &PgxClient{pgxConnPool, log}, nil
}

// 2) создаем соединение с бд (миграция)
func InitSQLPg(pgConfig config.PgConfig, log zerolog.Logger) (*sql.DB, error) {
	var sqlDB *sql.DB
	ttlRetry := 1 * time.Second

	connConfig := pgx.ConnConfig{
		Host:     pgConfig.Host,
		Port:     uint16(pgConfig.Port),
		Database: pgConfig.Database,
		User:     pgConfig.User,
		Password: pgConfig.Password,
	}

	retry := 1
	for retry < 10 {
		sqlDB = stdlib.OpenDB(connConfig)
		if sqlDB == nil {
			log.Error().Caller().Err(errors.New("ошибка при создании соединения с Pg")).Msgf("#%v retrying after %v sec.", retry, ttlRetry)
			retry++
			time.Sleep(ttlRetry)
			continue
		}
		break
	}

	if sqlDB == nil {
		log.Panic().Caller().Err(errors.New("ошибка при создании соединения с Pg"))
	}

	return sqlDB, nil
}
