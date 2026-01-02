package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

var urlPost = "postgresql://postgres.idpoonqbkjymjhitbmyh:tdtzDap%23w%2Ah%23D8C@aws-1-sa-east-1.pooler.supabase.com:6543/postgres?pgbouncer=true"

func InitConn() {
	var err error
	pool, err = pgxpool.New(ctx, urlPost)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	// Verify the connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}

func main() {
	InitConn()
}
