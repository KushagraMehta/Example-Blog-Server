package model

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

var testDB DB

func cleanTable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := testDB.db.Exec(ctx, "SELECT truncate_tables('postgres');")
	require.NoError(t, err)
}

func TestMain(m *testing.M) {
	var err error
	if testDB.db, err = pgxpool.Connect(context.Background(), "postgres://postgres:123@localhost:5432/blog"); err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	os.Exit(m.Run())
}
