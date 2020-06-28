package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Kamva/mgm"
)

type Cab struct {
	mgm.DefaultModel   `bson:"inline"`
	Model              string `json:"model" bson:"model"`
	Brand              string `json:"brand" bson:"brand"`
	RegistrationNumber string `json:"registration_number" bson:"registrationNumber"`
	CarType            string `json:"car_type" bson:"carType"`
	Color              string `json:"color" bson:"color"`
	RegistrationDate   string `json:"registration_date" bson:"registrationDate"`
	IsAvailable        bool   `json:"is_available" bson:"isAvailable"`
	CurrentLocation    struct {
		Type        string    `json:"type" bson:"type"`
		Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	}
}

// CreateCab is used for creating Cab
func CreateCab(c *Cab) *Cab {
	err := mgm.Coll(c).Create(c)
	if err != nil {
		return nil
	}
	return c
}

func GetCabByID(ID string) *Cab {
	cab := &Cab{}
	if err := mgm.Coll(cab).FindByID(ID, cab); err != nil {
		fmt.Println(err)
		return nil
	}
	return cab
}

func FindNearByCab(long float64, lat float64) []Cab {

	coll := mgm.Coll(&Cab{})
	cabs := []Cab{}
	query := bson.D{{"CurrentLocation", bson.D{{"$near", bson.D{{"$geometry", bson.D{{"type", "Point"}, {"coordinates", bson.A{long, lat}}}}, {"$maxDistance", 6000}, {"$minDistance", 500}}}}}}
	err := coll.SimpleFind(&cabs, query)
	if err != nil {
		fmt.Println(err)
	}
	return cabs
}
