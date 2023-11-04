package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"midkaGolang/models"
	"net/http"
	"strconv"
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

func GetPowerToolById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	powerToolID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var powerTool models.PowerTool
	if err := db.First(&powerTool, powerToolID).Error; err != nil {
		http.Error(w, "Power Tool not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(powerTool)
}

func CreatePowerTool(w http.ResponseWriter, r *http.Request) {
	var powerTool models.PowerTool
	json.NewDecoder(r.Body).Decode(&powerTool)
	db.Create(&powerTool)
	fmt.Println("Endpoint Hit: CreatePowerTool")
	json.NewEncoder(w).Encode(powerTool)
}

func UpdatePowerTool(w http.ResponseWriter, r *http.Request) {
	var updatedPowerTool models.PowerTool
	json.NewDecoder(r.Body).Decode(&updatedPowerTool)
	db.Save(&updatedPowerTool)
	json.NewEncoder(w).Encode(updatedPowerTool)
}

func UpdatePowerToolPatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var newData models.PowerTool
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingData models.PowerTool
	result := db.First(&existingData, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	db.Model(&existingData).Updates(newData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingData)
}

func DeletePowerTool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	powerToolID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var powerTool models.PowerTool
	if err := db.First(&powerTool, powerToolID).Error; err != nil {
		http.Error(w, "Power Tool not found", http.StatusNotFound)
		return
	}

	db.Delete(&powerTool)
	w.WriteHeader(http.StatusNoContent)
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
func GetPaintById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paintID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var paint models.Paint
	if err := db.First(&paint, paintID).Error; err != nil {
		http.Error(w, "Paint not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(paint)
}

// UpdatePaint updates a paint by its ID
func UpdatePaint(w http.ResponseWriter, r *http.Request) {
	var updatedPaint models.Paint
	json.NewDecoder(r.Body).Decode(&updatedPaint)
	db.Save(&updatedPaint)
	json.NewEncoder(w).Encode(updatedPaint)
}

// DeletePaint deletes a paint by its ID
func DeletePaint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paintID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var paint models.Paint
	if err := db.First(&paint, paintID).Error; err != nil {
		http.Error(w, "Paint not found", http.StatusNotFound)
		return
	}

	db.Delete(&paint)
	w.WriteHeader(http.StatusNoContent)
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

// GetNailScrewById retrieves a nail screw by its ID
func GetNailScrewById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nailScrewID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var nailScrew models.NailScrew
	if err := db.First(&nailScrew, nailScrewID).Error; err != nil {
		http.Error(w, "Nail Screw not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(nailScrew)
}

// UpdateNailScrew updates a nail screw by its ID
func UpdateNailScrew(w http.ResponseWriter, r *http.Request) {
	var updatedNailScrew models.NailScrew
	json.NewDecoder(r.Body).Decode(&updatedNailScrew)
	db.Save(&updatedNailScrew)
	json.NewEncoder(w).Encode(updatedNailScrew)
}

// DeleteNailScrew deletes a nail screw by its ID
func DeleteNailScrew(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nailScrewID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var nailScrew models.NailScrew
	if err := db.First(&nailScrew, nailScrewID).Error; err != nil {
		http.Error(w, "Nail Screw not found", http.StatusNotFound)
		return
	}

	db.Delete(&nailScrew)
	w.WriteHeader(http.StatusNoContent)
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

// GetPlumbingSupplyById retrieves a plumbing supply by its ID
func GetPlumbingSupplyById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	plumbingSupplyID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var plumbingSupply models.PlumbingSupply
	if err := db.First(&plumbingSupply, plumbingSupplyID).Error; err != nil {
		http.Error(w, "Plumbing Supply not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(plumbingSupply)
}

// UpdatePlumbingSupply updates a plumbing supply by its ID
func UpdatePlumbingSupply(w http.ResponseWriter, r *http.Request) {
	var updatedPlumbingSupply models.PlumbingSupply
	json.NewDecoder(r.Body).Decode(&updatedPlumbingSupply)
	db.Save(&updatedPlumbingSupply)
	json.NewEncoder(w).Encode(updatedPlumbingSupply)
}

// DeletePlumbingSupply deletes a plumbing supply by its ID
func DeletePlumbingSupply(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	plumbingSupplyID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var plumbingSupply models.PlumbingSupply
	if err := db.First(&plumbingSupply, plumbingSupplyID).Error; err != nil {
		http.Error(w, "Plumbing Supply not found", http.StatusNotFound)
		return
	}

	db.Delete(&plumbingSupply)
	w.WriteHeader(http.StatusNoContent)
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
func GetElectricalFixtureById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	electricalFixtureID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var electricalFixture models.ElectricalFixture
	if err := db.First(&electricalFixture, electricalFixtureID).Error; err != nil {
		http.Error(w, "Electrical Fixture not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(electricalFixture)
}

func UpdateElectricalFixture(w http.ResponseWriter, r *http.Request) {
	var updatedElectricalFixture models.ElectricalFixture
	json.NewDecoder(r.Body).Decode(&updatedElectricalFixture)
	db.Save(&updatedElectricalFixture)
	json.NewEncoder(w).Encode(updatedElectricalFixture)
}

func DeleteElectricalFixture(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	electricalFixtureID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var electricalFixture models.ElectricalFixture
	if err := db.First(&electricalFixture, electricalFixtureID).Error; err != nil {
		http.Error(w, "Electrical Fixture not found", http.StatusNotFound)
		return
	}

	db.Delete(&electricalFixture)
	w.WriteHeader(http.StatusNoContent)
}
