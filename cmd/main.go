package main

import (
	"log"
)

var mongoConnectionString = "mongodb://root:rootpw@127.0.0.1:27017"

type App struct {
	mcn     string
}

func (app App) run() {

}

func main() {
	log.Println("Go works")
	app := App{mcn: mongoConnectionString}
	app.run()
}
