package main

import (
	"BalanceRange/internal/config"
	"BalanceRange/internal/db"
	"BalanceRange/internal/payment"
	pg_conn "BalanceRange/pkg/pg"
	redis_conn "BalanceRange/pkg/redis"
	"math/rand"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type DB struct {
	pgDB    *sqlx.DB
	redisDB *redis.Client
}

func main() {
	rand.Seed(time.Now().UnixNano())
	config := config.GetConfig()
	pgClient, err := pg_conn.NewPgClient(config.Pg)
	if err != nil {
		panic(err)
	}

	rdsClient, err := redis_conn.NewRedisClient(config.Redis)
	if err != nil {
		panic(err)
	}

	dbs := DB{
		pgDB:    pgClient,
		redisDB: rdsClient,
	}
	if err != nil {
		panic(err)
	}

	dbService := db.NewDb(dbs.pgDB)
	dbService.Clear()

	paymentService := payment.NewPaymentsService(&dbService)
	err = paymentService.Start()
	if err != nil {
		panic(err)
	}
}
