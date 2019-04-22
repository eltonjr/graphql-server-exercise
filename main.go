package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/eltonjr/graphql-server-exercise/db"
	"github.com/eltonjr/graphql-server-exercise/web"
)

func main() {
	addr := "localhost:8080"

	r := httprouter.New()

	driverDao := db.NewDriverDao()
	router := web.NewRouter(driverDao)

	router.RegisterIndex(r)
	router.RegisterREST(r)

	fmt.Printf("Server running at: %s\n", addr)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		fmt.Printf("Server stopped: %v\n", err)
	}
}
