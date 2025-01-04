package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"go-gin/dto"
	"go-gin/infra"
	"go-gin/models"
	"go-gin/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestMain関数はGoのテストランナーによって全てのテストが実行される前に呼び出される
func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	// このファイルに含まれる全てのテストがここで呼び出される
	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "テストアイテム１", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム２", Price: 2000, Description: "テスト２", SoldOut: true, UserID: 1},
		{Name: "テストアイテム３", Price: 3000, Description: "テスト３", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@test.com", Password: "test1password"},
		{Email: "test2@test.com", Password: "test2password"},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	// given
	router := setup()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/items", nil)

	// when
	router.ServeHTTP(w, r)

	// then
	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 12, len(res["data"]))
}

func TestCreate(t *testing.T) {
	// given
	router := setup()

	token, err := services.CreateToken(1, "test1@test.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム４",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(reqBody))
	r.Header.Set("Authorization", "Bearer "+*token)

	// when
	router.ServeHTTP(w, r)

	// then
	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}
