/*
Copyright © 2024 Luka Piplica piplicaluka64@gmail.com
*/
package cmd

import (
	"errors"
	"nba-stats/models"
	"nba-stats/utils"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// Schedule API endpoint
var scheduleUrl string = "https://api-basketball.p.rapidapi.com/games?league=12&season=2023-2024&timezone=America/Toronto&date=" + utils.GetCurrentDate()

// ScheduleService for handling interaction with the schedule API
type ScheduleService struct {
	URL  string
	Body io.Reader
}

// NewScheduleService to initialize a ScheduleService type with default values
func NewScheduleService() *ScheduleService {
	return &ScheduleService{
		scheduleUrl,
		nil,
	}
}

// FetchSchedule receiver method associated with the ScheduleService type for
// fetch the schedule data from the API
func (s *ScheduleService) FetchSchedule() (*models.Schedule, error) {
	reqSchedule, err := http.NewRequest(http.MethodGet, s.URL, s.Body)
	if err != nil {
		return nil, err
	}

	reqSchedule.Header.Add(utils.ApiKeyHeader, os.Getenv("API_KEY"))
	reqSchedule.Header.Add(utils.ApiHostHeader, "api-basketball.p.rapidapi.com")

	resSchedule, err := http.DefaultClient.Do(reqSchedule)
	if err != nil {
		statusCode := resSchedule.StatusCode
		errorString := fmt.Sprintf("Unexpected response code: %d", statusCode)
		errors.New(errorString)
	}

	defer resSchedule.Body.Close()

	var schedule models.Schedule

	err = json.NewDecoder(resSchedule.Body).Decode(&schedule)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "A brief description of your command",
	Long:  "Nba stat data displayed in table format",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching Schedule Data")

		// Load in env variables
		err := godotenv.Load()
		if err != nil {
			fmt.Println("An Error Occurred: ", err)
			return
		}

		scheduleService := NewScheduleService()
		fmt.Println(scheduleService.FetchSchedule())
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
