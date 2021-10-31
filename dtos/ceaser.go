package dtos

import "time"

type CeasersPromotionResponse struct {
	SportID      string `json:"sportId"`
	Name         string `json:"name"`
	EventCount   int    `json:"eventCount"`
	Competitions []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Events []struct {
			Markets []struct {
				ID                    string `json:"id"`
				Name                  string `json:"name"`
				TemplateName          string `json:"templateName"`
				DisplayOrder          int    `json:"displayOrder"`
				Type                  string `json:"type"`
				Display               bool   `json:"display"`
				TradedInPlay          bool   `json:"tradedInPlay"`
				TemplateID            string `json:"templateId"`
				SpOffered             bool   `json:"spOffered"`
				Active                bool   `json:"active"`
				SportID               string `json:"sportId"`
				CompetitionID         string `json:"competitionId"`
				EventID               string `json:"eventId"`
				KeyMarket             bool   `json:"keyMarket"`
				DisplayName           string `json:"displayName"`
				CollectionName        string `json:"collectionName"`
				PrimaryDisplayOrder   int    `json:"primaryDisplayOrder"`
				SecondaryDisplayOrder int    `json:"secondaryDisplayOrder"`
				Selections            []struct {
					ID            string `json:"id"`
					Display       bool   `json:"display"`
					Name          string `json:"name"`
					SportID       string `json:"sportId"`
					CompetitionID string `json:"competitionId"`
					EventID       string `json:"eventId"`
					MarketID      string `json:"marketId"`
					Active        bool   `json:"active"`
					Type          string `json:"type"`
					Price         struct {
						F string  `json:"f"`
						D float64 `json:"d"`
						A int     `json:"a"`
					} `json:"price"`
				} `json:"selections"`
				SixPackView bool    `json:"sixPackView"`
				Line        float64 `json:"line,omitempty"`
				BetLimits   struct {
					LiabilityLimit int `json:"liabilityLimit"`
					MaxMultipleBet int `json:"maxMultipleBet"`
					MaxWin         int `json:"maxWin"`
					MaxBet         int `json:"maxBet"`
				} `json:"betLimits,omitempty"`
				Metadata struct {
					Line           string `json:"line"`
					MarketCategory string `json:"marketCategory"`
					MarketCode     string `json:"marketCode"`
					PricingSource  string `json:"pricingSource"`
					Period         string `json:"period"`
					MarketType     string `json:"marketType"`
				} `json:"metadata,omitempty"`
				ByoEligible bool `json:"byoEligible"`
			} `json:"markets"`
			ID                        string    `json:"id"`
			Name                      string    `json:"name"`
			StartTime                 time.Time `json:"startTime"`
			Type                      string    `json:"type"`
			Display                   bool      `json:"display"`
			TradedInPlay              bool      `json:"tradedInPlay"`
			SportID                   string    `json:"sportId"`
			CompetitionID             string    `json:"competitionId"`
			CompetitionName           string    `json:"competitionName"`
			Active                    bool      `json:"active"`
			Started                   bool      `json:"started"`
			MarketCountActivePreMatch int       `json:"marketCountActivePreMatch"`
			MarketCountActiveInPlay   int       `json:"marketCountActiveInPlay"`
			ByoEligible               bool      `json:"byoEligible"`
			LiveEventData             struct {
				EventID  string `json:"eventId"`
				GameTime struct {
					Seconds int `json:"seconds"`
					Minutes int `json:"minutes"`
				} `json:"gameTime"`
				Scores []struct {
					Team   string `json:"team"`
					Points int    `json:"points"`
				} `json:"scores"`
				StatsID   int       `json:"statsId"`
				StartTime time.Time `json:"startTime"`
				Quarter   int       `json:"quarter"`
			} `json:"liveEventData,omitempty"`
		} `json:"events"`
		GroupName string `json:"groupName"`
	} `json:"competitions"`
}