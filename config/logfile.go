package config

import (
	"log"
	"os"
)

var File *os.File

// / Crear el archivo loh
func CreateLogFile() {

	File, _ := os.OpenFile("logfile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	log.SetOutput(File)
	log.Println("Se ha creado el archivo log")

}

func CloseLogFile() {
	File.Close()
}
