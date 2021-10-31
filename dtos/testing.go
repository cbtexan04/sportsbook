package dtos

type MongoTestRecord struct {
	Name   string
	Spread TeamBettingOptions
}

// Team name to betting line
type TeamBettingOptions map[string]BettingOptions

// Betting line to Sportsbook
type BettingOptions map[string]Lines

// Book name to book value
type Lines map[string]BookLine

type BookLine struct {
	OddsDecimal float64
	Odds string
}


/*
{
	"name": "Bengals at Jets",
	"spread": {
		"Bengals": {
			"-3.5": {
				"Draft Kings": "-150",
				"Ceasers": "-140"
			}
		},
		"Jets": {
			"+3.5": {
				"Draft Kings": "-110",
				"Ceasers": "-110"
			}
		}
	},
	"totals": {
		"Over": {
			"40": {
				"Draft Kings": "-150",
				"Ceasers": "-140"
			},
			"50": {
				"Draft Kings": "-150",
				"Ceasers": "-140"
			}
		},
		"Under": {
			"40": {
				"Draft Kings": "-110",
				"Ceasers": "-110"
			},
			"50": {
				"Draft Kings": "-110",
				"Ceasers": "-110"
			}
		}
	},
	"moneylines": {
		"Bengals": {
			"Draft Kings": "-150",
			"Ceasers": "-140"
		},
		"Jets": {
			"Draft Kings": "-150",
			"Ceasers": "-140"
		}
	}
}
 */