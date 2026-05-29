package main

import (
	"log"
	"os"

	//	"os/exec"
	//	"runtime"
	"encoding/json"
	"fmt"
	"student_management_system/cli"
	"student_management_system/models"
	"student_management_system/session"
)

const dataFile = "storage/institute.json"

func main() {

	var institute *models.Institute
	var err error

	institute, err = loadInstituteFromFile(dataFile)

	if err != nil {
		log.Println("No saved data found, creating new Institute...")
		institute = newInstitute("No-Name")
	} else {
		log.Println("Loaded institute from file successfully")
	}

	session.Running = true

	for {
		if !session.Running {
			break
		}
		cli.Home(institute)
	}

	if err := saveInstituteToFile(institute, dataFile); err != nil {
		log.Printf("Failed to save institute: %v", err)
	} else {
		log.Printf("Institute saved successfully")
	}
}

func saveInstituteToFile(institute *models.Institute, filePath string) error {
	data, err := json.MarshalIndent(institute, "", "  ")

	if err != nil {
		return fmt.Errorf("marhsalling failed: %w", err)
	}

	return os.WriteFile(filePath, data, 0666)
}

func loadInstituteFromFile(filePath string) (*models.Institute, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var institute models.Institute
	err = json.Unmarshal(data, &institute)
	if err != nil {
		return nil, err
	}

	return &institute, nil
}

/*
func clearScreen() {

		var cmd *exec.Cmd

		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", "cls")
		} else {
			cmd = exec.Command("clear")
		}

		cmd.Stdout = os.Stdout
		cmd.Run()
	}
*/
func newInstitute(name string) *models.Institute {
	return &models.Institute{
		Name:    name,
		Classes: make(map[string]*models.Class),
		Users:   make(map[int]*models.User),
		NextID:  1,
	}
}
