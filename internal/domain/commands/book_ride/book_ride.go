package book_ride

import "github.com/EnzoDOROSARIO/uber-go/internal/domain/model/ride_booking"

func NewRideBooker(tripScanner TripScanner) *RideBooker {
	return &RideBooker{tripScanner: tripScanner}
}

type TripScanner interface {
	DistanceBetween(departure string, arrival string) int
}

type RideBooker struct {
	tripScanner TripScanner
}

func (b *RideBooker) Execute(
	rideId string,
	riderId string,
	departure string,
	arrival string,
) *ride_booking.Ride {
	distance := b.tripScanner.DistanceBetween(departure, arrival)
	price := 30.0 + float64(distance)*0.5
	return ride_booking.Book(rideId, riderId, departure, arrival, price)
}
