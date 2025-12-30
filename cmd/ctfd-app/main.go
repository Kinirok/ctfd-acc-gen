package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Kinirok/ctfd-acc-gen/ctfdgen"
	"github.com/Kinirok/ctfd-acc-gen/internal/ctfd"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var generator *ctfdgen.Generator
var ctx context.Context

func main() {
	_ = godotenv.Load()
	teamID := 1
	url := fmt.Sprintf("%s/teams/%d/members", os.Getenv("CTFD_URL"), teamID)
	//url := fmt.Sprintf("%s/users", os.Getenv("CTFD_URL"))
	userData := map[string]int{
		"id": 6,
	}
	fmt.Println(url)
	//user := ctfd.CreateUserRequest{Email: gen.GenerateEmail(), Name: "user_2526805876", Password: gen.GeneratePassword()}
	jsonData, _ := json.Marshal(userData)
	//bytes.NewBuffer(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	req.Header.Set("Authorization", "Token "+os.Getenv("CTFD_ADMIN_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("ошибка чтения ответа")
	}
	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		// Если ответ не JSON, возвращаем как есть
		fmt.Printf("Ответ не в формате JSON: %v\n", err)
	}
	// Проверка ответа
	if resp.StatusCode != 200 {
		fmt.Print(fmt.Errorf("ошибка создания: %d", resp.StatusCode))
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, bodyBytes, "", "  "); err == nil {
		fmt.Println("Тело ответа (JSON):")
		fmt.Println(prettyJSON.String())
	}
}

func main2() {

	_ = godotenv.Load()

	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable port=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	client := ctfd.NewCTFdClient(os.Getenv("CTFD_URL"), os.Getenv("CTFD_ADMIN_TOKEN"))

	cfg := ctfdgen.Config{CTFDClient: client, DB: db}

	generator, err = ctfdgen.NewGenerator(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx = context.Background()

	//Execute()
}
