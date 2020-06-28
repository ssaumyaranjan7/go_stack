package models

import (
	"fmt"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ride struct {
	mgm.DefaultModel `bson:"inline"`
	UserID           primitive.ObjectID `json:"user_id" bson:"userId"`
	CabID            primitive.ObjectID `json:"cab_id" bson:"cabId"`
	DriverID         primitive.ObjectID `json:"driver_id" bson:"driverId"`
	FromAddress      string             `json:"from_address" bson:"fromAddress"`
	ToAddress        string             `json:"to_address" bson:"toAddress"`
	FromLocation     []float64          `json:"from_location" bson:"fromLocation"`
	ToLocation       []float64          `json:"to_location" bson:"toLocation"`
	TotalPrice       int64              `json:"total_price" bson:"totalprice"`
	InvoiceURL       string             `json:"invoice_url" bson:"invoiceUrl"`
	RideCity         string             `json:"ride_city" bson:"rideCity"`
}

func CreateRide(r *Ride) *Ride {

	err := mgm.Coll(r).Create(r)
	if err != nil {
		return nil
	}
	return r
}

func GetRideByRideID(id string) *Ride {
	ride := &Ride{}
	if err := mgm.Coll(ride).FindByID(id, ride); err != nil {
		fmt.Println(err)
		return nil
	}
	return ride
}

func GetRideByUserID(id string) []Ride {
	coll := mgm.Coll(&Ride{})
	rides := []Ride{}
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err := coll.SimpleFind(&rides, bson.M{"userId": ID}); err != nil {
		fmt.Println(err)
		return nil
	}
	return rides
}
