package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"midkaGolang/handlers"
)

func main() {
	dsn := "user=postgres password=postgres dbname=midkaDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	handlers.InitDB(db)

	router := mux.NewRouter()

	// Define API routes for Power Tools
	router.HandleFunc("/powerTools", handlers.GetAllPowerTools).Methods("GET")
	router.HandleFunc("/powerTools", handlers.CreatePowerTool).Methods("POST")

	// Start the server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
