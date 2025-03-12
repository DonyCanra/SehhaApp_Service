package logger

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"sehha/pkg"
	"time"
)

func LoggerFile(*gin.Context) {
	currentDate := time.Now().Format("2006-01-02")

	directory, isLinuxOS := checkDirFile()
	logDir := filepath.Join(directory, "logger", "data")
	if isLinuxOS {
		logDir = directory
	}

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	fileName := fmt.Sprintf("sehha-service-%s.log", currentDate)
	filePath := filepath.Join(logDir, fileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.New(file, "", log.LstdFlags)
}

func Info(message interface{}, data interface{}) {
	// Get current time and format it
	currentTime := time.Now()
	currentTimeString := currentTime.Format(pkg.StdDateTimeFormat)
	currentDate := currentTime.Format(pkg.StdDateFormat)

	// Marshal data to JSON
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling data to JSON: %v", err)
		return
	}

	// Determine log directory based on OS
	directory, isLinuxOS := checkDirFile()
	logDir := filepath.Join(directory, "logger", "data")
	if isLinuxOS {
		logDir = directory
	}

	// Construct log file path
	fileName := fmt.Sprintf("sehha-service-%s.log", currentDate)
	filePath := filepath.Join(logDir, fileName)

	// Get caller information
	pc, file, line, _ := runtime.Caller(1)
	funcName := getFunctionName(pc)

	// Log message format
	logMessage := fmt.Sprintf("[%s] [%s:%d] [%s] [%v] %s", currentTimeString, file, line, funcName, message, string(dataJSON))
	logPrintln := fmt.Sprintf("[%s] [%v] %s", currentTimeString, message, string(dataJSON))
	fmt.Println(logPrintln)
	// Open or create the log file
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer logFile.Close()

	// Create a new logger instance
	logWriter := log.New(logFile, " ", 0)

	// Write log message
	logWriter.Println(logMessage)
}

func Error(message interface{}, data interface{}) {
	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	currentDate := currentTime.Format("2006-01-02")

	// Marshal data to JSON
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling data to JSON: %v", err)
		return
	}

	// Determine log directory based on OS
	directory, isLinuxOS := checkDirFile()
	logDir := filepath.Join(directory, "logger", "data")
	if isLinuxOS {
		logDir = directory
	}

	// Construct log file path
	fileName := fmt.Sprintf("sehha-service-%s.log", currentDate)
	filePath := filepath.Join(logDir, fileName)

	// Get caller information
	pc, file, line, _ := runtime.Caller(1)
	funcName := getFunctionName(pc)

	// Log message format
	logMessage := fmt.Sprintf("[%s] [%s:%d] [%s] [%v] %s", currentTimeString, file, line, funcName, message, dataJSON)

	// Open or create the log file
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer logFile.Close()

	// Create a new logger instance
	logWriter := log.New(logFile, " ", 0)

	// Write log message
	logWriter.Println(logMessage)
}

func InfoRawString(message interface{}, data interface{}) {
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	directory, isLinuxOS := checkDirFile()
	logDir := filepath.Join(directory, "logger", "data")
	if isLinuxOS {
		logDir = directory
	}

	fileName := fmt.Sprintf("sehha-service-%s.log", currentDate)
	fileName = filepath.Join(logDir, fileName)

	pc, file, line, _ := runtime.Caller(1)
	funcName := getFunctionName(pc)

	// Log message format
	logMessage := fmt.Sprintf("[%s] [%s:%d] [%s] [%v] %s", currentTimeString, file, line, funcName, message, data)
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer logFile.Close()
	logWriter := log.New(logFile, "", 0)
	fmt.Println(logMessage)

	logWriter.Println(logMessage)
}

func checkDirFile() (string, bool) {
	osData := runtime.GOOS
	var directory string
	var err error
	isLinux := false
	if osData == "windows" || osData == "darwin" {
		directory, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		isLinux = true
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		localAddress := conn.LocalAddr().(*net.UDPAddr).IP

		directory = "/mnt/Sehha_Service_" + localAddress.String()

		//directory, err = filepath.Abs(filepath.Dir(os.Args[0]))
		//if err != nil {
		//	log.Fatal(err)
		//}
	}

	return directory, isLinux
}

func getFunctionName(pc uintptr) string {
	if fn := runtime.FuncForPC(pc); fn != nil {
		return fn.Name()
	}
	return ""
}
