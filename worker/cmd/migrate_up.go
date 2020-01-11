package cmd

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"os"
	"worker/config"
	"worker/postgres"
)

var migrateUp = &cobra.Command{
	Use:   "migrate_up",
	Short: "start migration up in Postgres",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrationUp()
	},
}

func runMigrationUp() {

	// зададим настройки логирования в приложении
	log := zerolog.New(os.Stdout).With().Logger()

	// чтение с конфига с помощью библиотеки Viper
	cfg, err := config.InitConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("Ошибка")
	}

	// запуск миграции Pg
	postgres.MigratePgUp(cfg.PgConfig)

}
