/*
Copyright © 2024 Luka Piplica piplicaluka64@gmail.com
*/
package cmd

import (
	"nba/models"
	"nba/utils"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// Shcedule API endpoint

var scheduleUrl string = "https://api-basketball.p.rapidapi.com/games?league=12&season=2023-2024&timezone=America/Toronto&date=" + getCurrentDate()

// Struct for handeling interaction with the scheudle API
type ScheduleService struct {
	URL  string
	Body io.Reader
}

// Initialize a ScheduleService type with deault values
func NewScheduleService() *ScheduleService {
	return &ScheduleService{
		scheduleUrl,
		nil,
	}
}

// Receiver method associated with the ScheduleService type for
// fetch the schedule data from the API
func (s *ScheduleService) FetchSchedule() {
	reqSchedule, err := http.NewRequest("GET", s.URL, s.Body)
	if err != nil {
		fmt.Println("An error occured: ", err)
	}

	reqSchedule.Header.Add(utils.ApiKeyHeader, os.Getenv("API_KEY"))
	reqSchedule.Header.Add(utils.ApiHostHeader, "api-basketball.p.rapidapi.com")

	resSchedule, err := http.DefaultClient.Do(reqSchedule)
	if err != nil {
		fmt.Println("An error occured: ", err)
	}

	defer resSchedule.Body.Close()

	var schedule models.Schedule
	//var scheduleJSON map[string]interface{}

	err = json.NewDecoder(resSchedule.Body).Decode(&schedule)
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	fmt.Println(schedule.Response)
}

func getCurrentDate() string {
	// Get the current date
	currentTime := time.Now().UTC()

	const YYYYMMDD = "2006-01-02"   // Must use this string to get the format as YYYYMMDD
	date := currentTime.Format(YYYYMMDD)
	fmt.Println(date)

	return date
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
		scheduleService.FetchSchedule()
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
