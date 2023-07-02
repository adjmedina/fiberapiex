package database

import (
	//importamos el paquete GORM

	"adjmedina/fiberapiex/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	//Se define una variable global DBConn que se
	//utiliza para conectarse a la base de datos SQLite.
	DBConn *gorm.DB
)

func InitDatabase() {

	var err error

	//database.DBConn es exportada por el paquete database
	DBConn, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}

	log.Println("Database Conectada")

	// Migra la database
	DBConn.AutoMigrate(&models.User{})

	fmt.Println("Database Migrada")

}

func CloseDatabase() {
	DBConn.Close()
}
