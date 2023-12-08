package services

import (
	"fmt"
	"log"
)

type (
	Log     struct{}
	LogFunc struct {
		name string
	}
)

func NewLog() *Log { return &Log{} }

func (l Log) Func(n string) LogFunc { return LogFunc{name: n} }

func (f LogFunc) BadRequest(msg string, args ...any) {
	log.Println(f.name + " BadRequest: " + fmt.Sprintf(msg, args...))
}

func (f LogFunc) NotFound(msg string, args ...any) {
	log.Println(f.name + " Not Found: " + fmt.Sprintf(msg, args...))
}

func (f LogFunc) OK(msg string, args ...any) {
	log.Println(f.name + " OK: " + fmt.Sprintf(msg, args...))
}

func (f LogFunc) InternalSerer(msg string, args ...any) {
	log.Println(f.name + " InternalSerer: " + fmt.Sprintf(msg, args...))
}
