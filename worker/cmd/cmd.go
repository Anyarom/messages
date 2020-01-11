package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Worker для обработки сообщений из очереди",
}

func Execute() {
	// добавляем команды
	rootCmd.AddCommand(migrateUp)
	rootCmd.AddCommand(startWorker)

	err := rootCmd.Execute()
	if err != nil {
		log.Error().Err(err).Caller().Msg("")
		os.Exit(1)
	}
}
