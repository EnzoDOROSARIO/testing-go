package book_ride

import "github.com/EnzoDOROSARIO/uber-go/internal/domain/model/ride_booking"

func NewRideBooker() *RideBooker {
	return &RideBooker{}
}

type RideBooker struct{}

func (b *RideBooker) Execute(
	rideId string,
	riderId string,
	departure string,
	arrival string,
) *ride_booking.Ride {
	return ride_booking.Book(rideId, riderId, departure, arrival)
}
