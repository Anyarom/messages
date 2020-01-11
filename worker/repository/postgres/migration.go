package postgres

import (
	"github.com/rs/zerolog"
	"github.com/rubenv/sql-migrate"
	"os"
	"worker/config"
)

var migrations = &migrate.MemoryMigrationSource{
	Migrations: []*migrate.Migration{
		{
			Id: "1",
			Up: []string{
				`create table if not exists messages
				(
					id bigserial not null constraint messages_pk primary key,
					phone text,
					text text
				);
				
				alter table messages owner to postgres;
				
				create unique index if not exists messages_id_uindex on messages (id);
				`,
			},
			Down: []string{
				`drop table messages;`,
			},
		},
	},
}

func MigratePgUp(pgConfig config.PgConfig) {

	// зададим настройки логирования в приложении
	log := zerolog.New(os.Stdout).With().Caller().Logger()

	// создадим подключение в БД
	sqlDB, err := InitSQLPg(pgConfig, log)
	if err != nil {
		log.Fatal().Err(err).Msg("can't init pg")
	}

	// выполним миграцию
	affected, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("не выполнено")
	}
	log.Info().Msgf("up migration applied: %v", affected)

}

func MigratePgDown(cfg config.Config) {

	// зададим настройки логирования в приложении
	log := zerolog.New(os.Stdout).With().Caller().Logger()

	// создадим подключение в БД
	sqlDB, err := InitSQLPg(cfg.PgConfig, log)
	if err != nil {
		log.Fatal().Err(err).Msg("can't init pg")
	}

	// выполним миграцию
	affected, err := migrate.ExecMax(sqlDB, "postgres", migrations, migrate.Down, 1)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}
	log.Info().Msgf("down migration applied: %v", affected)

}
