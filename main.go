package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Practice struct {
	ID          uint
	IntValue    int
	RealValue   float32
	StringValue string
}

type GormStandard struct {
	gorm.Model
	IntValue    int
	RealValue   float32
	StringValue string
}

func main() {
	log.Println("main run")

	rand.Seed(time.Now().UnixNano())

	db, err := connectDb()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	createTables(db)
	createSingularTables(db)

	insertRecords(db)

	updateRecords(db)
}

func connectDb() (*gorm.DB, error) {
	configs := map[string]string{}
	configs["host"] = os.Getenv("DB_HOST")
	configs["port"] = os.Getenv("DB_PORT")
	configs["user"] = os.Getenv("DB_USER")
	configs["password"] = os.Getenv("DB_PASSWORD")
	configs["dbname"] = os.Getenv("DB_SCHEMA")
	configs["sslmode"] = os.Getenv("DB_SSL")

	buf := []string{}
	for k, v := range configs {
		buf = append(buf, k+"="+v)
	}
	params := strings.Join(buf, " ")

	db, err := gorm.Open("postgres", params)
	if err != nil {
		log.Println("DB connect error!!")
		log.Println(err)
		return nil, err
	}

	log.Println("DB connect success!")

	return db, nil
}

func createTables(db *gorm.DB) {
	db.CreateTable(&GormStandard{})

	log.Println("Create tables success")
}

func createSingularTables(db *gorm.DB) {
	db.SingularTable(true)

	db.CreateTable(&Practice{})

	db.SingularTable(false)

	log.Println("Create singular tables success")
}

func insertRecords(db *gorm.DB) {
	practice := Practice{
		IntValue:    rand.Intn(100),
		StringValue: strconv.Itoa(rand.Intn(100)),
		RealValue:   rand.Float32() * 1000}
	db.Table("practice").Create(&practice)

	standard := GormStandard{
		IntValue:    rand.Intn(100),
		StringValue: strconv.Itoa(rand.Intn(100)),
		RealValue:   rand.Float32() * 1000}
	db.Create(&standard)

	log.Println("Insert records success")
}

func updateRecords(db *gorm.DB) {
	var practice Practice
	db.Table("practice").Order("id").First(&practice)
	practice.IntValue = rand.Intn(100)
	db.Table("practice").Save(&practice)

	var standard GormStandard
	db.Order("id").First(&standard)
	standard.IntValue = rand.Intn(100)
	db.Save(&standard)

	log.Println("Update records success")
}
