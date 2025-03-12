package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type ReadYml struct {
	Port string `yaml:"port"`
}

func (c *ReadYml) GetConf() *ReadYml {
	// get work directory

	// DebugMode = local
	// TestMode = Staging / Dev
	// ReleaseMode = Production
	gin.SetMode(gin.DebugMode)
	name := ""
	osData := runtime.GOOS
	if strings.EqualFold(osData, "windows") || strings.EqualFold(osData, "darwin") {
		directory, err := os.Getwd() //get the current directory using the built-in function
		if err != nil {
			fmt.Println(err) //print the error if obtained
		}
		if gin.Mode() == gin.DebugMode {
			name = directory + "/env/local.yaml"
		} else if gin.Mode() == gin.TestMode {
			name = directory + "/env/staging.yaml"
		} else {
			name = directory + "/env/production.yaml"

		}
	} else {
		directory, err := filepath.Abs(filepath.Dir(os.Args[0])) //get the current working directory
		if err != nil {
			log.Fatal(err) //print the error if obtained
		}

		if gin.Mode() == gin.DebugMode {
			name = directory + "/env/local.yaml"
		} else if gin.Mode() == gin.TestMode {
			name = directory + "/env/staging.yaml"
		} else {
			name = directory + "/env/production.yaml"

		}
	}

	yamlFile, err := os.ReadFile(name)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
