package models

import (
	"Sportsbook/dtos"
)

type Promotions []Promotion
type Promotion struct {
	Name string
	Book string
	Spread GameLines
	Moneyline GameLines
	Totals GameLines
}

func (p Promotion) ToMongoPromotion() dtos.MongoPromotionRecord {
	return dtos.MongoPromotionRecord{
		EventName: p.Name,
		Spread: p.Spread.toMongoGameLines(p.Book),
		Totals: p.Totals.toMongoGameLines(p.Book),
		Moneyline: p.Moneyline.toMongoGameLines(p.Book),
	}
}

type GameLines []GameLine

type GameLine struct {
	LineChoice string
	LineOdds string
	LineDecimal float64
	Line string
}

func (g GameLines) toMongoGameLines(book string) []dtos.MongoGameLineRecord {
	gameLines := make([]dtos.MongoGameLineRecord, len(g))
	for i, game := range g {
		gameLines[i] = game.toMongoGameLine(book)
	}
	return gameLines
}

func (g GameLine) toMongoGameLine(book string) dtos.MongoGameLineRecord {
	return dtos.MongoGameLineRecord{
		Book: book,
		LineDecimal: g.LineDecimal,
		LineOdds: g.LineOdds,
		LineChoice: g.LineChoice,
		Line: g.Line,
	}
}