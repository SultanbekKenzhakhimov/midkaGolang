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

	// Paints
	router.HandleFunc("/paints", handlers.GetAllPaints).Methods("GET")
	router.HandleFunc("/paints", handlers.CreatePaint).Methods("POST")

	// Nail Screws
	router.HandleFunc("/nailScrews", handlers.GetAllNailScrews).Methods("GET")
	router.HandleFunc("/nailScrews", handlers.CreateNailScrew).Methods("POST")

	// Plumbing Supplies
	router.HandleFunc("/plumbingSupplies", handlers.GetAllPlumbingSupplies).Methods("GET")
	router.HandleFunc("/plumbingSupplies", handlers.CreatePlumbingSupply).Methods("POST")

	// Electrical Fixtures
	router.HandleFunc("/electricalFixtures", handlers.GetAllElectricalFixtures).Methods("GET")
	router.HandleFunc("/electricalFixtures", handlers.CreateElectricalFixture).Methods("POST")

	// Start the server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
