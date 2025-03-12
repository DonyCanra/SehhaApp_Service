package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var db *gorm.DB

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Connection() *gorm.DB {
	fmt.Println("yaml config")
	gin.SetMode(gin.DebugMode) // Adjust mode if needed

	var config Config
	var name string
	osData := runtime.GOOS

	if strings.EqualFold(osData, "windows") || strings.EqualFold(osData, "darwin") {
		directory, err := os.Getwd()
		check(err)
		fmt.Println("gin mode " + gin.Mode())
		switch gin.Mode() {
		case gin.DebugMode:
			name = directory + "/env/local.yaml"
		case gin.TestMode:
			name = directory + "/env/staging.yaml"
		default:
			name = directory + "/env/production.yaml"
		}
	} else {
		directory, err := filepath.Abs(filepath.Dir(os.Args[0]))
		check(err)
		fmt.Println("gin mode " + gin.Mode())
		switch gin.Mode() {
		case gin.DebugMode:
			name = directory + "/env/local.yaml"
		case gin.TestMode:
			name = directory + "/env/staging.yaml"
		default:
			name = directory + "/env/production.yaml"
		}
	}

	file, err := os.Open(name)
	check(err)
	defer file.Close()

	content, err := io.ReadAll(file)
	check(err)

	err = yaml.Unmarshal(content, &config)
	check(err)

	// PostgreSQL DSN format: "host=localhost user=username password=secret dbname=mydb port=5432 sslmode=disable"
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Database.Host, config.Database.User, config.Database.Password,
		config.Database.DBName, config.Database.Port, config.Database.SSLMode,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection failed:", err)
	} else {
		log.Println("connection to db success")
	}

	Migration(db)
	return db
}
