package db

import (
	"log"
	"os"

	"database/sql"
	"github.com/google/uuid"
	sqliteGo "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"restapi/env"
)

var Conn *gorm.DB

func Connect() {
	sql.Register("sqlite3_extended",
		&sqliteGo.SQLiteDriver{
			ConnectHook: func(conn *sqliteGo.SQLiteConn) error {
				err := conn.RegisterFunc(
					"gen_random_uuid",
					func(arguments ...interface{}) (string, error) {
						return uuid.New().String(), nil // Return a string value.
					},
					true,
				)
				return err
			},
		},
	)

	conn, err := sql.Open("sqlite3_extended", "boox.db")
	if err != nil {
		panic("Failed to connect to the DB!")
	}

	mode := env.Get[string]("GIN_MODE")
	config := &gorm.Config{}
	if mode != "release" {
		config = &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{LogLevel: logger.Info},
			),
		}
	}

	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite3_extended",
		DSN:        "boox.db",
		Conn:       conn,
	}, config)

	err = db.AutoMigrate(&User{}, &Book{})
	if err != nil {
		return
	}

	if db == nil {
		panic("DB is <nil>")
	}
	Conn = db
}
