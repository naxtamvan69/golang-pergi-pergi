package utils

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"pergipergi/model"
)

var db *gorm.DB

/*
Melakukan setup untuk koneksi ke Database menggunakan gorm pgx
*/
func ConnectDB() error {
	//connect using gorm pgx
	connection, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	connection.AutoMigrate(model.Destinasi{}, model.Role{}, model.TravelAgensi{}, model.User{})
	SetDBConnection(connection)

	return nil
}

func SetDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}

func Reset(db *gorm.DB, table string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE " + table).Error; err != nil {
			return err
		}

		if err := tx.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1").Error; err != nil {
			return err
		}

		return nil
	})
}
