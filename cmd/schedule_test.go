package cmd

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestFetchSchedule(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Get the current dir and it's parent dir
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting the current directory: %v", err)
	}
	parentDir := filepath.Dir(wd)

	err = godotenv.Load(parentDir + "/.env")
	if err != nil {
		log.Fatalf("An Error Occurred: %v", err)
	}

	scheduleService := NewScheduleService()
	schedule, err := scheduleService.FetchSchedule()

	if err != nil {

	}

	// Assert that there is no error
	assert.NoError(t, err)

	log.Println("Response: ", schedule)

	// TODO: Add more assertions based on the expected behavior of your function
	// For example, you can use the testify/assert package for assertions
	// For demonstration purposes, let's assert that the response is as expected

}
