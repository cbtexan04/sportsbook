package main

import (
	"Sportsbook/dtos"
	"Sportsbook/services"
	"strings"
)

func main() {
	ceaser, _ := services.GetCeaserGameLines()
	dk, _ := services.GetDkGameLines()

	records := make([]dtos.MongoPromotionRecord, 0)
	for _, r := range *ceaser {
		if strings.Replace(r.Name, " at ", "", -1) == "" {
			continue
		}
		records = append(records, r.ToMongoPromotion())
	}
	for _, r := range *dk {
		if strings.Replace(r.Name, " at ", "", -1) == "" {
			continue
		}
		records = append(records, r.ToMongoPromotion())
	}

	datastore := services.NewDatastore()
	err := datastore.InsertPromotions(records)
	if err != nil {
		panic(err)
	}
}