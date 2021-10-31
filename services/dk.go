package services

import (
	"Sportsbook/clients"
	"Sportsbook/dtos"
	"Sportsbook/models"
	"errors"
	"fmt"
)

func GetDkGameLines() (*[]models.Promotion, error) {
	response, err := clients.GetDraftKingsPromotions()
	if err != nil {
		return nil, err
	}

	offers, err := getOffers(response)
	if err != nil {
		return nil, err
	}

	return getDkPromotions(offers), nil
}

func getOffers(response *dtos.DraftKingsResponse) ([][]dtos.Offers, error) {
	categories := response.EventGroup.OfferCategories
	for _, category := range categories {
		if category.Name == "Game Lines"{
			for _, descriptors := range category.OfferSubcategoryDescriptors {
				if descriptors.Name == "Game" {
					return descriptors.OfferSubcategory.Offers, nil
				}
			}
		}
	}

	return nil, errors.New("unable to find game lines")
}

func getDkPromotions(offers [][]dtos.Offers) *[]models.Promotion {
	if len(offers) == 0 {
		return nil
	}

	promotions := make([]models.Promotion, 0, len(offers[0]))
	for _, game := range offers {
		promotion := models.Promotion{
			Book: "Draft Kings",
		}

		if len(game) == 0{
			continue
		}

		for _, bettingOption := range game {
			if !bettingOption.IsOpen {
				continue
			}

			teamName1 := normalizeTeamName(bettingOption.Outcomes[0].Label)
			teamName2 := normalizeTeamName(bettingOption.Outcomes[1].Label)
			if teamName1 == "" || teamName2 == "" {
				continue
			}

			gameline1 := models.GameLine{
				LineOdds: bettingOption.Outcomes[0].OddsAmerican,
				LineChoice: teamName1,
				LineDecimal: bettingOption.Outcomes[0].OddsDecimal,
				Line: fmt.Sprintf("%v", bettingOption.Outcomes[0].Line),
			}

			gameline2 := models.GameLine{
				LineOdds: bettingOption.Outcomes[1].OddsAmerican,
				LineChoice: teamName2,
				LineDecimal: bettingOption.Outcomes[1].OddsDecimal,
				Line: fmt.Sprintf("%v", bettingOption.Outcomes[1].Line),
			}

			switch bettingOption.Label {
			case "Spread":
				gameline1.LineChoice = teamName1
				gameline2.LineChoice = teamName2
				promotion.Name =  teamName1 + " at " + teamName2
				promotion.Spread = append(promotion.Spread, gameline1)
				promotion.Spread = append(promotion.Spread, gameline2)
			case "Total":
				promotion.Totals = append(promotion.Totals, gameline1)
				promotion.Totals = append(promotion.Totals, gameline2)
			case "Moneyline":
				gameline1.LineChoice = teamName1
				gameline2.LineChoice = teamName2
				promotion.Moneyline = append(promotion.Moneyline, gameline1)
				promotion.Moneyline = append(promotion.Moneyline, gameline2)
			default:
				continue
			}
		}
		promotions = append(promotions, promotion)
	}

	return &promotions
}