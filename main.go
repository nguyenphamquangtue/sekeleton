package main

import (
	"log"
	"net/http"
	"skeleton/define"
	"skeleton/driver"
	"skeleton/router"
)

func main() {

	db, err := driver.GetConnection()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &router.Router{
		JWTKey: define.JWTKey,
		Conn:   db.Conn,
	})
}
