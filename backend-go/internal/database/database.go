package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Iemaduddin/goweb/backend-go/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func MySQLDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Terjadi kesalahan: ", err)
	}

	log.Println("Koneksi ke database berhasil")

	return db, nil
}
