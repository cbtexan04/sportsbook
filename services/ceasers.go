package services

import (
	"Sportsbook/clients"
	"Sportsbook/dtos"
	"Sportsbook/models"
	"fmt"
)

func GetCeaserGameLines() (*[]models.Promotion, error) {
	response, err := clients.GetCeaserPromotions()
	if err != nil {
		return nil, err
	}

	return getCeaserPromotions(response), nil
}

func getCeaserPromotions(response *dtos.CeasersPromotionResponse) *[]models.Promotion {
	promotions := make([]models.Promotion, 0)
	for _, promo := range response.Competitions {
		if promo.Name == "NFL" {
			for _, game := range promo.Events {
				if  game.Name == "" {
					continue
				}
				promotion := models.Promotion{ Book: "Ceaser" }

				for _, bettingOption := range game.Markets {
					if len(bettingOption.Selections) < 2 {
						continue
					}

					gameline1 := models.GameLine{
						LineOdds: fmt.Sprintf("%d", bettingOption.Selections[0].Price.A),
						LineChoice: bettingOption.Selections[0].Name,
						LineDecimal: bettingOption.Selections[0].Price.D,
						Line: fmt.Sprintf("%v", bettingOption.Line*-1),
					}

					gameline2 := models.GameLine{
						LineOdds: fmt.Sprintf("%d", bettingOption.Selections[1].Price.A),
						LineChoice: bettingOption.Selections[1].Name,
						LineDecimal: bettingOption.Selections[1].Price.D,
						Line: fmt.Sprintf("%v", bettingOption.Line),
					}

					teamName1 := normalizeTeamName(gameline1.LineChoice)
					teamName2 := normalizeTeamName(gameline2.LineChoice)

					switch bettingOption.DisplayName {
					case "Spread Live": fallthrough
					case "Spread":
						gameline1.LineChoice = teamName1
						gameline2.LineChoice = teamName2
						promotion.Name =  teamName1 + " at " + teamName2
						promotion.Spread = append(promotion.Spread, gameline1)
						promotion.Spread = append(promotion.Spread, gameline2)
					case "Total Points Live": fallthrough
					case "Total Points":
						promotion.Totals = append(promotion.Totals, gameline1)
						promotion.Totals = append(promotion.Totals, gameline2)
					case "Money Line Live": fallthrough
					case "Money Line":
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
		}
	}

	return &promotions
}