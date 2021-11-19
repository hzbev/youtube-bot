package helper

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var prod string = GetEnv()

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ReadtoArray(filePath string) []string {
	var workingDir string
	if prod == "dev" {
		workingDir, _ = os.Getwd()
	} else {
		exeDir, _ := os.Executable()
		workingDir = filepath.Dir(exeDir)
	}

	file, err := os.Open(filepath.Join(workingDir, filePath))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content
}

func Write(filePath, text string) {
	var workingDir string
	if prod == "dev" {
		workingDir, _ = os.Getwd()
	} else {
		exeDir, _ := os.Executable()
		workingDir = filepath.Dir(exeDir)
	}

	file, err := os.OpenFile(workingDir+`/`+filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := file.Write([]byte("\n" + text)); err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func GetEnv() string {
	return os.Getenv("APP_ENV")
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}
