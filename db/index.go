package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Availibility struct {
	bun.BaseModel `bun:"table:Availibility"`

	Id             int64     `bun:",pk,autoincrement"`
	Time           time.Time `bun:"default:current_timestamp"`
	NumberOfPeople int       `bun:"type:integer"`
}

func ConnectToDB(ctx context.Context) *bun.DB {
	// host, ok := os.LookupEnv("HOST")
	// if !ok {
	host := "localhost"
	// }
	dsn := fmt.Sprintf("postgres://admin:password@%s:5432/fialka?sslmode=disable", host)

	fmt.Println(dsn)
	var err error
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	db.NewCreateTable().Model((*Availibility)(nil)).IfNotExists().Exec(ctx)
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
	return db
}
