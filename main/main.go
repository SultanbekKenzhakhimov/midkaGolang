package main

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"midkaGolang/handlers"
	"net/http"
)

func main() {
	dsn := "user=postgres password=postgres dbname=midkaDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	handlers.InitDB(db)

	router := mux.NewRouter()
	// Роут для Sign Up User
	router.HandleFunc("/sign-up", handlers.CreateUser).Methods("POST")
	// Роуты для Power Tools
	router.Handle("/powerTools", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetAllPowerTools))).Methods("GET")
	router.Handle("/powerTools/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetPowerToolById))).Methods("GET")
	router.Handle("/powerTools", handlers.AuthMiddleware(http.HandlerFunc(handlers.CreatePowerTool))).Methods("POST")
	router.Handle("/powerTools/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.UpdatePowerTool))).Methods("PUT")
	router.Handle("/powerTools/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.DeletePowerTool))).Methods("DELETE")
	router.Handle("/powerTools/{id}", handlers.AuthMiddleware(http.HandlerFunc(handlers.UpdatePowerToolPatch))).Methods("PATCH")

	// Роуты для Paints
	router.HandleFunc("/paints", handlers.GetAllPaints).Methods("GET")
	router.HandleFunc("/paints/{id}", handlers.GetPaintById).Methods("GET")
	router.HandleFunc("/paints", handlers.CreatePaint).Methods("POST")
	router.HandleFunc("/paints/{id}", handlers.UpdatePaint).Methods("PUT")
	router.HandleFunc("/paints/{id}", handlers.DeletePaint).Methods("DELETE")
	router.HandleFunc("/paints/{id}", handlers.UpdatePaintPatch).Methods("PATCH")
	// Nail Screws
	router.HandleFunc("/nailScrews", handlers.GetAllNailScrews).Methods("GET")
	router.HandleFunc("/nailScrews/{id}", handlers.GetNailScrewById).Methods("GET")
	router.HandleFunc("/nailScrews", handlers.CreateNailScrew).Methods("POST")
	router.HandleFunc("/nailScrews/{id}", handlers.UpdateNailScrew).Methods("PUT")
	router.HandleFunc("/nailScrews/{id}", handlers.DeleteNailScrew).Methods("DELETE")

	// Plumbing Supplies
	router.HandleFunc("/plumbingSupplies", handlers.GetAllPlumbingSupplies).Methods("GET")
	router.HandleFunc("/plumbingSupplies/{id}", handlers.GetPlumbingSupplyById).Methods("GET")
	router.HandleFunc("/plumbingSupplies", handlers.CreatePlumbingSupply).Methods("POST")
	router.HandleFunc("/plumbingSupplies/{id}", handlers.UpdatePlumbingSupply).Methods("PUT")
	router.HandleFunc("/plumbingSupplies/{id}", handlers.DeletePlumbingSupply).Methods("DELETE")

	// Electrical Fixtures
	router.HandleFunc("/electricalFixtures", handlers.GetAllElectricalFixtures).Methods("GET")
	router.HandleFunc("/electricalFixtures/{id}", handlers.GetElectricalFixtureById).Methods("GET")
	router.HandleFunc("/electricalFixtures", handlers.CreateElectricalFixture).Methods("POST")
	router.HandleFunc("/electricalFixtures/{id}", handlers.UpdateElectricalFixture).Methods("PUT")
	router.HandleFunc("/electricalFixtures/{id}", handlers.DeleteElectricalFixture).Methods("DELETE")

	// Роут для входа пользователя
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Защищенный роут, требующий валидного JWT токена
	router.Handle("/protected", handlers.AuthMiddleware(http.HandlerFunc(handlers.ProtectedEndpoint))).Methods("GET")
	// Start the server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
