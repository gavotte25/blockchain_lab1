package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Logger struct {
	Enable bool
}

func (l *Logger) Println(v ...any) {
	if l.Enable {
		log.Println(v...)
	}
}

func (l *Logger) Panicln(v ...any) {
	if l.Enable {
		log.Panicln(v...)
	}
}

// used to convert an int64 to a byte array
func ConvertTimestampToByte(num int64) []byte {
	/*
		Function to convert an int64 to a byte array in Big Endian format.

		:param num: int64 number to be converted
		:return: byte array representing the hex value of the input number
	*/
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		// Handle the error in a way suitable for your application
		panic(err)
	}

	return buff.Bytes()
}

func GetTimestampFormat(timestamp int64) string {
	unixTimeUTC := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return unixTimeUTC
}

func WipeFolder(folderPath string) error {
	targetFolder, err := os.Open(folderPath)
	if err != nil {
		return err
	}

	defer targetFolder.Close()
	files, err := targetFolder.Readdir(-1)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(folderPath, file.Name())
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetStringEncode(input []byte) string {
	return hex.EncodeToString(input[:])
}

func GetStringDecode(input []byte) string {
	return hex.EncodeToString(input[:])
}

func GetLog(logType, message string) {
	// Default logType to "INFO" if it is nil or empty
	if logType == "" {
		logType = "info"
	}

	// Format the current timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Print the log to the screen
	fmt.Printf("%s %s: %s\n", timestamp, strings.ToUpper(logType), strings.ToLower(message))
}

func WriteFile(data []string, fileName string, path string) error {
	filePath := filepath.Join(path, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, row := range data {
		_, err := f.WriteString(row + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFile(fileName string, path string) []string {
	filePath := filepath.Join(path, fileName)
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var inputFiles []string
	for scanner.Scan() {
		inputFiles = append(inputFiles, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil
	}
	return inputFiles
}

func ConvertStringToInt(data string) (int, error) {
	result, err := strconv.Atoi(data)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func TrimInputByOS(input string) string {
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(input, "\r\n")
	} else {
		input = strings.TrimRight(input, "\n")
	}
	return input
}

func DeleteVerifiedRowInFile(dir string, filename string, valueToRemove string) error {
	inFile, err2 := os.OpenFile(dir+"/"+filename, os.O_RDWR, 0644)
	if err2 != nil {
		//fmt.Println(err2)
		return err2
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	inFile.Seek(0, 0)
	lineNum := 0
	var data []string
	for scanner.Scan() {
		scannerValue := strings.Replace(scanner.Text(), "\t", "", -1)
		scannerValue = strings.Replace(scannerValue, "\n", "", -1)
		fmt.Println(scannerValue)
		if scannerValue == valueToRemove {
			continue
		} else {
			// Skip writing this line
			lineNum++
			data = append(data, scanner.Text())
		}
	}
	err := WriteFile(data, filename, dir)
	if err != nil {
		return err
	}
	return nil
}
