/*
Copyright © 2024 Luka Piplica piplicaluka64@gmail.com
*/
package cmd

import (
	"io"
	"nba/models"

	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Constants for API endpoint and headers
const (
	apiUrl        = "https://api-basketball.p.rapidapi.com/standings?league=12&season=2023-2024"
	apiKeyHeader  = "X-RapidAPI-Key"
	apiHostHeader = "X-RapidAPI-Host"
	apiKey        = "API_KEY"
)

// Handles interactions with the standings API
type StandingsService struct {
	URL  string
	Body io.Reader
}

// Create a new StandingsService with default values
func NewStandingsService() *StandingsService {
	return &StandingsService{
		apiUrl,
		nil,
	}
}

// Fetch the standings data from NBA api
func (s *StandingsService) FetchStandings() (*models.Standings, error) {
	req, err := http.NewRequest("GET", s.URL, s.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(apiKeyHeader, apiKey)
	req.Header.Add(apiHostHeader, "api-basketball.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var standings models.Standings
	err = json.NewDecoder(res.Body).Decode(&standings)
	if err != nil {
		return nil, err
	}

	return &standings, nil
}

// DisplayStandingsTable displays the standings in a table
func DisplayStandingsTable(standings *models.Standings) {
	// Display standings in table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Position", "Conference", "Team", "W", "L", "PCT"})

	for _, response := range standings.Response {
		for _, r := range response {
			// Extract relevant fields from each Response struct
			position := fmt.Sprintf("%d", r.Position)
			conference := r.Group.Name
			team := r.Team.Name
			wins := fmt.Sprintf("%d", r.Games.Win.Total)
			losses := fmt.Sprintf("%d", r.Games.Lose.Total)
			winPercentage := r.Games.Win.Percentage

			// Append a slice of strings to the table
			table.Append([]string{position, conference, team, wins, losses, winPercentage})
		}
	}

	// Display the table
	table.Render()
}

// standingsCmd represents the standings command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "View the updated standings in the NBA",
	Long:  "Nba stat data displayed in table format",

	Run: func(cmd *cobra.Command, args []string) {
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
