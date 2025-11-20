package logger

import (
	"log"
)

func Err(txt string) {
	log.Printf("Ошибка %s", txt)
}

func Info(txt string) {
	log.Printf("Инфо %s", txt)
}
