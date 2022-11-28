package main

import (
	"bufio"
	"context"
	"flag"
	"github.com/Kin-dza-dzaa/DNSVladivostok/mapaccess"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"os"
	"time"
)

const (
	DATABASE_URL = "postgres://dns_vladivostok:12345@localhost:5432/dns_assigment"
)

var (
	fileNameFlag *string
	timeFlag     *uint
	logger       zerolog.Logger
	mapAccess    *mapacess.MapAccess
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger = zerolog.New(zerolog.SyncWriter(os.Stdout))

	timeFlag = flag.Uint("time", 1, "wait span")
	fileNameFlag = flag.String("name", "", "name of a file")
	flag.Parse()
	if *fileNameFlag == "" {
		flag.Usage()
		os.Exit(0)
	}

	file, err := os.Open(*fileNameFlag)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	config, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	mapAccess = mapacess.NewMapAccess(&logger, pool, file)
}

func main() {
	scanner := bufio.NewScanner(mapAccess.File)
	for scanner.Scan() {
		if err := mapAccess.InsertOne(scanner.Text()); err != nil {
			logger.Error().Msg(err.Error())
		}
		time.Sleep(time.Second * (time.Duration(*timeFlag)))
	}
}
