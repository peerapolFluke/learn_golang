package injection

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"log"
	"seamoor/config"
	"seamoor/ent"
	"seamoor/ent/migrate"
	"time"
)

var DB *ent.Client

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil
	}
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(100)
	db.SetConnMaxLifetime(time.Hour)
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func InitDatabase() {
	client := Open(config.Env.DatabaseUrl)
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	DB = client
}
