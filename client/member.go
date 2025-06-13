package client

import "fmt"

const (
	memberProfileURL = "member/profile"
	memberChartURL   = "member/chart_data"
	memberAwardsURL  = "member/awards"
	memberDataURL    = "member/get"
)

type LicenceType int

const (
	Oval LicenceType = iota + 1
	Road
	DirtOval
	DirtRoad
	SportsCar
	FormulaCar
)

type ChartType int

const (
	IRating ChartType = iota + 1
	TTRating
	SafetyRating
)

type MemberProfile struct {
	Activity struct {
		Recent30DaysCount int `json:"recent_30days_count"`
		ConsecutiveWeeks  int `json:"consecutive_weeks"`
	} `json:"activity"`

	Success bool `json:"success"`

	LicenceHistory []struct {
		CategoryID   int     `json:"category_id"`
		SafetyRating float32 `json:"safety_rating"`
	} `json:"license_history"`

	MemberInfo struct {
		Licenses []struct {
			CategoryID   int     `json:"category_id"`
			SafetyRating float32 `json:"safety_rating"`
			CategoryName string  `json:"category_name"`
			GroupName    string  `json:"group_name"`
		} `json:"licenses"`

		DisplayName string `json:"display_name"`
		MemberSince string `json:"member_since"`
	} `json:"member_info"`

	RecentEvents []struct {
		EventType string `json:"event_type"`
		StartTime string `json:"start_time"`
		CarName   string `json:"car_name"`
	} `json:"recent_events"`
}

func (c *Client) GetMemberProfile(customerId int) (MemberProfile, error) {
	queryPath := fmt.Sprintf("?cust_id=%d", customerId)
	return followLink[MemberProfile](c, createURL(memberProfileURL)+queryPath)
}

type ChartData struct {
	Blackout   bool `json:"blackout"`
	CategoryId int  `json:"category_id"`
	ChartType  int  `json:"chart_type"`
	Data       []struct {
		When  string `json:"when"`
		Value int    `json:"value"`
	} `json:"data"`
	Success    bool `json:"success"`
	CustomerId int  `json:"cust_id"`
}

func (c *Client) GetUserCharts(customerId int, category LicenceType, chartType ChartType) (ChartData, error) {
	queryPath := fmt.Sprintf("?cust_id=%d&category_id=%d&chart_type=%d", customerId, category, chartType)
	return followLink[ChartData](c, createURL(memberChartURL)+queryPath)
}
