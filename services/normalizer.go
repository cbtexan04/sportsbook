package services

import (
	"strings"
)

func normalizeTeamName(givenName string) string {
	if strings.Contains(givenName, "Bengals") { return "Cincinnati Bengals" }
	if strings.Contains(givenName, "Cincinnati") { return "Cincinnati Bengals" }

	if strings.Contains(givenName, "Houston") { return "Houston Texans" }
	if strings.Contains(givenName, "Texans") { return "Houston Texans" }

	if strings.Contains(givenName, "New England") { return "New England Patriots" }
	if strings.Contains(givenName, "Patriots") { return "New England Patriots" }

	if strings.Contains(givenName, "New York Jets") { return "New York Jets" }
	if strings.Contains(givenName, "Jets") { return "New York Jets" }

	if strings.Contains(givenName, "Steelers") { return "Pittsburgh Steelers" }
	if strings.Contains(givenName, "Pittsburgh") { return "Pittsburgh Steelers" }

	if strings.Contains(givenName, "Browns") { return "Cleveland Browns" }
	if strings.Contains(givenName, "Cleveland") { return "Cleveland Browns" }

	if strings.Contains(givenName, "Baltimore") { return "Baltimore Ravens" }
	if strings.Contains(givenName, "Ravens") { return "Baltimore Ravens" }

	if strings.Contains(givenName, "Denver") { return "Denver Broncos" }
	if strings.Contains(givenName, "Broncos") { return "Denver Broncos" }

	if strings.Contains(givenName, "Jacksonville") { return "Jacksonville Jaguars" }
	if strings.Contains(givenName, "Jaguars") { return "Jacksonville Jaguars" }

	if strings.Contains(givenName, "Tampa Bay") { return "Tampa Bay Buccaneers" }
	if strings.Contains(givenName, "Buccaneers") { return "Tampa Bay Buccaneers" }

	if strings.Contains(givenName, "San Fransisco") { return "San Fransisco 49ers" }
	if strings.Contains(givenName, "49ers") { return "San Fransisco 49ers" }

	if strings.Contains(givenName, "Arizona") { return "Arizona Cardinals" }
	if strings.Contains(givenName, "Cardinals") { return "Arizona Cardinals" }

	if strings.Contains(givenName, "Buffalo") { return "Buffalo Bills" }
	if strings.Contains(givenName, "Bills") { return "Buffalo Bills" }

	if strings.Contains(givenName, "Green Bay") { return "Packers" }
	if strings.Contains(givenName, "Packers") { return "Packers" }

	if strings.Contains(givenName, "Dallas") { return "Dallas Cowboys" }
	if strings.Contains(givenName, "Cowboys") { return "Dallas Cowboys" }

	if strings.Contains(givenName, "Detroit") { return "Detroit Lions" }
	if strings.Contains(givenName, "Lions") { return "Detroit Lions" }

	if strings.Contains(givenName, "Chicago") { return "Chicago Bears" }
	if strings.Contains(givenName, "Bears") { return "Chicago Bears" }

	if strings.Contains(givenName, "New Orleans") { return "New Orleans Saints" }
	if strings.Contains(givenName, "Saints") { return "New Orleans Saints" }

	if strings.Contains(givenName, "Los Angeles") { return "Los Angeles Rams" }
	if strings.Contains(givenName, "Rams") { return "Los Angeles Rams" }

	if strings.Contains(givenName, "Minnesota") { return "Minnesota Vikings" }
	if strings.Contains(givenName, "Vikings") { return "Minnesota Vikings" }

	if strings.Contains(givenName, "Washington") { return "Washington Football Team" }
	if strings.Contains(givenName, "Football Team") { return "Washington Football Team" }

	if strings.Contains(givenName, "Tennessee") { return "Tennessee Titans" }
	if strings.Contains(givenName, "Titans") { return "Tennessee Titans" }

	if strings.Contains(givenName, "Seattle") { return "Seattle Seahawks" }
	if strings.Contains(givenName, "Seahawks") { return "Seattle Seahawks" }

	if strings.Contains(givenName, "Philadelphia") { return "Philadelphia Eagles" }
	if strings.Contains(givenName, "Eagles") { return "Philadelphia Eagles" }

	if strings.Contains(givenName, "Indianapolis") { return "Indianapolis Colts" }
	if strings.Contains(givenName, "Colts") { return "Indianapolis Colts" }

	if strings.Contains(givenName, "Kansas City") { return "Kansas City Chiefs" }
	if strings.Contains(givenName, "Chiefs") { return "Kansas City Chiefs" }

	if strings.Contains(givenName, "Las Vegas") { return "Las Vegas Raiders" }
	if strings.Contains(givenName, "Raiders") { return "Las Vegas Raiders" }

	if strings.Contains(givenName, "Carolina") { return "Carolina Panthers" }
	if strings.Contains(givenName, "Panthers") { return "Carolina Panthers" }

	if strings.Contains(givenName, "Atlanta") { return "Atlanta Falcons" }
	if strings.Contains(givenName, "Falcons") { return "Atlanta Falcons" }

	if strings.Contains(givenName, "Miami") { return "Miami Dolphins" }
	if strings.Contains(givenName, "Dolphins") { return "Miami Dolphins" }

	if strings.Contains(givenName, "New York Giants") { return "New York Giants" }
	if strings.Contains(givenName, "Giants") { return "New York Giants" }

	return ""
}