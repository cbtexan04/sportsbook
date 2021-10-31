package dtos

type MongoPromotionRecords []MongoPromotionRecord

type MongoPromotionRecord struct {
	EventName string
	Spread []MongoGameLineRecord
	Moneyline []MongoGameLineRecord
	Totals []MongoGameLineRecord
}

type MongoGameLineRecord struct {
	Book string
	LineChoice string
	Line string
	LineOdds string
	LineDecimal float64
}