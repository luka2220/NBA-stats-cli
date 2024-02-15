/*
Copyright © 2024 Luka Piplica piplicaluka64@gmail.com
*/
package cmd

import (
	"io"
	"nba-stats/models"
	"nba-stats/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Standings API endpoint
const (
	standingsUrl = "https://api-basketball.p.rapidapi.com/standings?league=12&season=2023-2024"
)

// Handles interactions with the standings API
type StandingsService struct {
	URL  string
	Body io.Reader
}

// Create a new StandingsService with default values
func NewStandingsService() *StandingsService {
	return &StandingsService{
		standingsUrl,
		nil,
	}
}

// Fetch the standings data from NBA api
// Receiver method associated with the StandingService type
func (s *StandingsService) FetchStandings() (*models.Standings, error) {
	req, err := http.NewRequest("GET", s.URL, s.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(utils.ApiKeyHeader, os.Getenv("API_KEY"))
	req.Header.Add(utils.ApiHostHeader, "api-basketball.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var standings models.Standings

	// Convert JSON data in the defined models
	err = json.NewDecoder(res.Body).Decode(&standings)
	if err != nil {
		return nil, err
	}

	return &standings, nil
}

// DisplayStandingsTable displays the standings in a table
func DisplayStandingsTable(standings *models.Standings) {
	// Create West conference standings
	westConference := tablewriter.NewWriter(os.Stdout)
	westConference.SetHeader([]string{"Position", "Conference", "Team", "W", "L", "PCT"})
	westConference.SetCaption(true, "West Standings")

	// Create East conference standings
	eastConference := tablewriter.NewWriter(os.Stdout)
	eastConference.SetHeader([]string{"Position", "Conference", "Team", "W", "L", "PCT"})
	eastConference.SetCaption(true, "East Standings")

	for _, response := range standings.Response {
		for _, r := range response {
			// Extract relevant fields from each Response struct
			position := fmt.Sprintf("%d", r.Position)
			team := r.Team.Name
			wins := fmt.Sprintf("%d", r.Games.Win.Total)
			losses := fmt.Sprintf("%d", r.Games.Lose.Total)
			winPercentage := r.Games.Win.Percentage

      if conference := r.Group.Name; conference == "Western Conference" {
				westConference.Append([]string{position, conference, team, wins, losses, winPercentage})
			} else if conference == "Eastern Conference" {
				eastConference.Append([]string{position, conference, team, wins, losses, winPercentage})
			}
		}
	}

	// Display tables
	westConference.Render()
	eastConference.Render()
}

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "View the updated standings in the NBA",
	Long:  "Nba stat data displayed in table format",

	Run: func(cmd *cobra.Command, args []string) {
		// Load in env variables
		err := godotenv.Load()
		if err != nil {
			fmt.Println("An Error Occurred: ", err)
			return
		}

		standingsService := NewStandingsService()

		standings, err := standingsService.FetchStandings()
		if err != nil {
			fmt.Println("An Error Occurred: ", err)
			return
		}

		DisplayStandingsTable(standings)
	},
}

func init() {
	rootCmd.AddCommand(standingsCmd)
}
