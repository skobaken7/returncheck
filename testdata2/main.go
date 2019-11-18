package main

import (
	"log"

	"github.com/go-xorm/xorm"
)

func Hoge() bool {
	createEngine()
	return true
}

func createEngine() *xorm.Engine {
	return &xorm.Engine{}
}

func main() {
	log.Println(Hoge())
}
