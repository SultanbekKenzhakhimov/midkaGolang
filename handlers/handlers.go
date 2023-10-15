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
