package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Kinirok/ctfd-acc-gen/ctfdgen"
	"github.com/Kinirok/ctfd-acc-gen/internal/ctfd"
	"github.com/Kinirok/ctfd-acc-gen/internal/gen"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var generator *ctfdgen.Generator
var ctx context.Context

func main() {
	pass1 := gen.GeneratePassword()
	fmt.Printf(pass1)
	for range 1000 {
		pass2 := gen.GeneratePassword()
		if pass2 == pass1 {
			fmt.Printf("%s: %s", pass2, pass1)
		}
		pass1 = pass2
	}
}
func main1() {

	_ = godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
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
