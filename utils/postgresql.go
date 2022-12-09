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
