package handlers

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"midkaGolang/models"
	"net/http"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
	db.AutoMigrate(&models.PowerTool{}, &models.Paint{}, &models.NailScrew{}, &models.PlumbingSupply{}, &models.ElectricalFixture{})
}

func GetAllPowerTools(w http.ResponseWriter, r *http.Request) {
	var powerTools []models.PowerTool
	db.Find(&powerTools)
	fmt.Println("Endpoint Hit: GetAllPowerTools")
	json.NewEncoder(w).Encode(powerTools)
}

func CreatePowerTool(w http.ResponseWriter, r *http.Request) {
	var powerTool models.PowerTool
	json.NewDecoder(r.Body).Decode(&powerTool)
	db.Create(&powerTool)
	fmt.Println("Endpoint Hit: CreatePowerTool")
	json.NewEncoder(w).Encode(powerTool)
}
func GetAllPaints(w http.ResponseWriter, r *http.Request) {
	var paints []models.Paint
	db.Find(&paints)
	fmt.Println("Endpoint Hit: GetAllPaints")
	json.NewEncoder(w).Encode(paints)
}

func CreatePaint(w http.ResponseWriter, r *http.Request) {
	var paint models.Paint
	json.NewDecoder(r.Body).Decode(&paint)
	db.Create(&paint)
	fmt.Println("Endpoint Hit: CreatePaint")
	json.NewEncoder(w).Encode(paint)
}

func GetAllNailScrews(w http.ResponseWriter, r *http.Request) {
	var nailScrews []models.NailScrew
	db.Find(&nailScrews)
	fmt.Println("Endpoint Hit: GetAllNailScrews")
	json.NewEncoder(w).Encode(nailScrews)
}

func CreateNailScrew(w http.ResponseWriter, r *http.Request) {
	var nailScrew models.NailScrew
	json.NewDecoder(r.Body).Decode(&nailScrew)
	db.Create(&nailScrew)
	fmt.Println("Endpoint Hit: CreateNailScrew")
	json.NewEncoder(w).Encode(nailScrew)
}

func GetAllPlumbingSupplies(w http.ResponseWriter, r *http.Request) {
	var plumbingSupplies []models.PlumbingSupply
	db.Find(&plumbingSupplies)
	fmt.Println("Endpoint Hit: GetAllPlumbingSupplies")
	json.NewEncoder(w).Encode(plumbingSupplies)
}

func CreatePlumbingSupply(w http.ResponseWriter, r *http.Request) {
	var plumbingSupply models.PlumbingSupply
	json.NewDecoder(r.Body).Decode(&plumbingSupply)
	db.Create(&plumbingSupply)
	fmt.Println("Endpoint Hit: CreatePlumbingSupply")
	json.NewEncoder(w).Encode(plumbingSupply)
}

func GetAllElectricalFixtures(w http.ResponseWriter, r *http.Request) {
	var electricalFixtures []models.ElectricalFixture
	db.Find(&electricalFixtures)
	fmt.Println("Endpoint Hit: GetAllElectricalFixtures")
	json.NewEncoder(w).Encode(electricalFixtures)
}

func CreateElectricalFixture(w http.ResponseWriter, r *http.Request) {
	var electricalFixture models.ElectricalFixture
	json.NewDecoder(r.Body).Decode(&electricalFixture)
	db.Create(&electricalFixture)
	fmt.Println("Endpoint Hit: CreateElectricalFixture")
	json.NewEncoder(w).Encode(electricalFixture)
}
