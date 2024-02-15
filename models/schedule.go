package models

// Schedule root JSON response type
type Schedule struct {
	Errors   []int              `json:"errors"`
	Response []ScheduleResponse `json:"response"`
}

// ScheduleResponse JSON response type
type ScheduleResponse struct {
	Time  string        `json:"time"`
	Teams HomeAwayTeams `json:"teams"`
}

// HomeAwayTeams and away playing JSON response type
type HomeAwayTeams struct {
	Home ScheduledTeam `json:"home"`
	Away ScheduledTeam `json:"Away"`
}

// ScheduledTeam JSON response type
type ScheduledTeam struct {
	Name string `json:"name"`
}
