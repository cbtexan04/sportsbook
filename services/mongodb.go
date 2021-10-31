package services

import (
	"Sportsbook/dtos"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type MongoDatastore struct {
	db      *mongo.Database
	Session *mongo.Client
}

func NewDatastore() *MongoDatastore {
	var mongoDataStore *MongoDatastore
	db, session := connect()
	if db != nil && session != nil {
		mongoDataStore = new(MongoDatastore)
		mongoDataStore.db = db
		mongoDataStore.Session = session
		return mongoDataStore
	}

	return nil
}

func connect() (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo()
	})

	return db, session
}

func connectToMongo() (a *mongo.Database, b *mongo.Client) {
	var cred options.Credential

	cred.AuthSource = "admin"
	cred.Username = "root"
	cred.Password = "example"

	var err error
	session, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred))
	if err != nil {
		panic(err)
	}
	err = session.Connect(context.TODO())
	if err != nil {
		panic(err)
	}

	var DB = session.Database("testing")
	return DB, session
}

func aggregatePromotions(models dtos.MongoPromotionRecords) []dtos.MongoTestRecord {
	foo := make(map[string]dtos.MongoTestRecord)

	for _, givenBet := range models {
		if event, ok := foo[givenBet.EventName]; ok {
			for _, spreadBet := range givenBet.Spread {
				if lineOptions, ok := event.Spread[spreadBet.LineChoice]; ok {
					if sportsBooks, ok := lineOptions[spreadBet.Line]; ok {
						if _, ok := sportsBooks[spreadBet.Book]; !ok {
							line := dtos.BookLine{
								OddsDecimal: spreadBet.LineDecimal,
								Odds: spreadBet.LineOdds,
							}
							sportsBooks[spreadBet.Book] = line
						}
					} else {
						line := dtos.BookLine{
							OddsDecimal: spreadBet.LineDecimal,
							Odds: spreadBet.LineOdds,
						}
						bookMap := make(dtos.Lines)
						bookMap[spreadBet.Book] = line
						lineOptions[spreadBet.Line] = bookMap
					}
				} else {
					line := dtos.BookLine{
						OddsDecimal: spreadBet.LineDecimal,
						Odds: spreadBet.LineOdds,
					}
					teamMap := make(dtos.BettingOptions)
					bookMap := make(dtos.Lines)
					bookMap[spreadBet.Book] = line
					teamMap[spreadBet.Line] = bookMap
					event.Spread[spreadBet.LineChoice] = teamMap
				}
			}
		} else {
			spreads := make(dtos.TeamBettingOptions)

			for _, spread := range givenBet.Spread {
				line := dtos.BookLine{
					OddsDecimal: spread.LineDecimal,
					Odds:        spread.LineOdds,
				}
				teamMap := make(dtos.BettingOptions)
				bookMap := make(dtos.Lines)
				bookMap[spread.Book] = line
				teamMap[spread.Line] = bookMap
				spreads[spread.LineChoice] = teamMap
			}

			foo[givenBet.EventName] = dtos.MongoTestRecord{
				Name:   givenBet.EventName,
				Spread: spreads,
			}
		}
	}

	retvals := make([]dtos.MongoTestRecord, 0)
	for _, v := range foo {
		retvals = append(retvals, v)
	}

	return retvals
}

func (client MongoDatastore) InsertPromotion(record dtos.MongoPromotionRecord) error {
	_, err := client.Session.Database("testing").Collection("football").InsertOne(context.TODO(), record)
	return err
}

func (client MongoDatastore) InsertPromotions(records dtos.MongoPromotionRecords) error {
	err := client.DeleteAll()
	if err != nil {
		return err
	}
	
	promotions := aggregatePromotions(records)
	newValue := make([]interface{}, len(promotions))

	for i, v := range promotions {
		newValue[i] =  v
	}

	_, err = client.Session.Database("testing").Collection("football").InsertMany(context.TODO(), newValue)
	return err
}

func (client MongoDatastore) DeleteAll() error {
	return client.db.Drop(context.TODO())
}