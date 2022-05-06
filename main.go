package main

import (
	"github.com/SuddenGunter/lapunder/config"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
	"path"
)

func main() {
	cfg, err := config.Load()
	handleErr(err)

	db, err := bbolt.Open(path.Join(cfg.Workdir, cfg.DB), 0600, nil)
	handleErr(err)

	defer db.Close()
}

func handleErr(err error) {
	if err != nil {
		logger := log.With().Err(err).Logger()
		logger.Fatal().Msg("failed to start lapunder app")
	}
}
