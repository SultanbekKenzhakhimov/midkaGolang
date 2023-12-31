package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"midkaGolang/mailer"
	"midkaGolang/models"
	"net/http"
	"strconv"
	"time"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
	db.AutoMigrate(&models.PowerTool{}, &models.Paint{}, &models.NailScrew{}, &models.PlumbingSupply{}, &models.ElectricalFixture{}, models.User{})
}

//SIGN UP user

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Role = "user" // set default role
	db.Create(&user)

	//добавлю таймаут задержку
	time.Sleep(5 * time.Second)

	// После создания пользователя вызывайте функцию handleSuccessfulRegistration
	mailer.HandleSuccessfulRegistration(user)

	fmt.Println("Endpoint Hit: CreateUser")
	json.NewEncoder(w).Encode(user)
}

//CRUD для сущностей

func GetAllPowerTools(w http.ResponseWriter, r *http.Request) {
	var powerTools []models.PowerTool
	db.Find(&powerTools)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)

	json.NewEncoder(w).Encode(powerTool)
}

func CreatePowerTool(w http.ResponseWriter, r *http.Request) {
	var powerTool models.PowerTool
	json.NewDecoder(r.Body).Decode(&powerTool)
	db.Create(&powerTool)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	fmt.Println("Endpoint Hit: CreatePowerTool")
	json.NewEncoder(w).Encode(powerTool)
}

func UpdatePowerTool(w http.ResponseWriter, r *http.Request) {
	var updatedPowerTool models.PowerTool
	json.NewDecoder(r.Body).Decode(&updatedPowerTool)
	db.Save(&updatedPowerTool)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)

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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	w.WriteHeader(http.StatusNoContent)
}

func GetAllPaints(w http.ResponseWriter, r *http.Request) {
	var paints []models.Paint
	db.Find(&paints)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	fmt.Println("Endpoint Hit: GetAllPaints")
	json.NewEncoder(w).Encode(paints)
}

func CreatePaint(w http.ResponseWriter, r *http.Request) {
	var paint models.Paint
	json.NewDecoder(r.Body).Decode(&paint)
	db.Create(&paint)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)

	json.NewEncoder(w).Encode(paint)
}

// UpdatePaint updates a paint by its ID
func UpdatePaint(w http.ResponseWriter, r *http.Request) {
	var updatedPaint models.Paint
	json.NewDecoder(r.Body).Decode(&updatedPaint)
	db.Save(&updatedPaint)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	w.WriteHeader(http.StatusNoContent)
}
func UpdatePaintPatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var newData models.Paint
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingData models.Paint
	result := db.First(&existingData, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	db.Model(&existingData).Updates(newData)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingData)
}

func GetAllNailScrews(w http.ResponseWriter, r *http.Request) {
	var nailScrews []models.NailScrew
	db.Find(&nailScrews)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	fmt.Println("Endpoint Hit: GetAllNailScrews")
	json.NewEncoder(w).Encode(nailScrews)
}

func CreateNailScrew(w http.ResponseWriter, r *http.Request) {
	var nailScrew models.NailScrew
	json.NewDecoder(r.Body).Decode(&nailScrew)
	db.Create(&nailScrew)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)

	json.NewEncoder(w).Encode(nailScrew)
}

// UpdateNailScrew updates a nail screw by its ID
func UpdateNailScrew(w http.ResponseWriter, r *http.Request) {
	var updatedNailScrew models.NailScrew
	json.NewDecoder(r.Body).Decode(&updatedNailScrew)
	db.Save(&updatedNailScrew)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	w.WriteHeader(http.StatusNoContent)
}

func GetAllPlumbingSupplies(w http.ResponseWriter, r *http.Request) {
	var plumbingSupplies []models.PlumbingSupply
	db.Find(&plumbingSupplies)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	fmt.Println("Endpoint Hit: GetAllPlumbingSupplies")
	json.NewEncoder(w).Encode(plumbingSupplies)
}

func CreatePlumbingSupply(w http.ResponseWriter, r *http.Request) {
	var plumbingSupply models.PlumbingSupply
	json.NewDecoder(r.Body).Decode(&plumbingSupply)
	db.Create(&plumbingSupply)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)

	json.NewEncoder(w).Encode(plumbingSupply)
}

// UpdatePlumbingSupply updates a plumbing supply by its ID
func UpdatePlumbingSupply(w http.ResponseWriter, r *http.Request) {
	var updatedPlumbingSupply models.PlumbingSupply
	json.NewDecoder(r.Body).Decode(&updatedPlumbingSupply)
	db.Save(&updatedPlumbingSupply)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	w.WriteHeader(http.StatusNoContent)
}
func GetAllElectricalFixtures(w http.ResponseWriter, r *http.Request) {
	var electricalFixtures []models.ElectricalFixture
	db.Find(&electricalFixtures)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	fmt.Println("Endpoint Hit: GetAllElectricalFixtures")
	json.NewEncoder(w).Encode(electricalFixtures)
}

func CreateElectricalFixture(w http.ResponseWriter, r *http.Request) {
	var electricalFixture models.ElectricalFixture
	json.NewDecoder(r.Body).Decode(&electricalFixture)
	db.Create(&electricalFixture)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)

	json.NewEncoder(w).Encode(electricalFixture)
}

func UpdateElectricalFixture(w http.ResponseWriter, r *http.Request) {
	var updatedElectricalFixture models.ElectricalFixture
	json.NewDecoder(r.Body).Decode(&updatedElectricalFixture)
	db.Save(&updatedElectricalFixture)
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
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
	//добавляю таймаут задержку
	time.Sleep(5 * time.Second)
	w.WriteHeader(http.StatusNoContent)
}

// Login обрабатывает запрос на вход и выдает JWT токен при успешной аутентификации
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	result := db.Where("email = ? AND password = ?", user.Email, user.Password).First(&existingUser)
	if result.Error != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Создаем токен с добавлением роли пользователя
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": existingUser.Email,
		"role":     existingUser.Role, // Добавляем роль пользователя в токен
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Отправляем токен в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// ProtectedEndpoint является защищенным маршрутом, который требует валидного токена для доступа
func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	// Получаем токен из заголовка Authorization
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Проверяем валидность токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Возвращаем секретный ключ для проверки подписи токена
		return []byte("your-secret-key"), nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Если токен валиден, продолжаем выполнение защищенного кода

	// Пример ответа
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "This is a protected endpoint"}`))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка Authorization
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Проверяем валидность токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Возвращаем секретный ключ для проверки подписи токена
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Получаем информацию о ролях из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Проверяем роли
		role, ok := claims["role"].(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Получаем информацию о маршруте (эндпоинте) с использованием mux.Vars
		vars := mux.Vars(r)
		route := r.URL.Path

		// Проверяем разрешения на основе роли и маршрута
		switch role {
		case "admin":
			// Администратор имеет доступ ко всем эндпоинтам
			next.ServeHTTP(w, r)
		case "user":
			// Пользователь имеет доступ к определенным эндпоинтам
			switch route {
			case "/powerTools", "/powerTools/{id}":
				// Разрешения для эндпоинтов powerTools
				if vars["id"] != "" {
					// Действия для конкретного powerTool с заданным ID
					next.ServeHTTP(w, r)
				} else {
					// Действия для списка powerTools
					next.ServeHTTP(w, r)
				}
			case "/paints", "/paints/{id}":
				// Разрешения для эндпоинтов paints
				if vars["id"] != "" {
					// Действия для конкретного paint с заданным ID
					next.ServeHTTP(w, r)
				} else {
					// Действия для списка paints
					next.ServeHTTP(w, r)
				}
			// Добавьте другие эндпоинты по аналогии
			default:
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		default:
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
