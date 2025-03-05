package book_ride

import "github.com/EnzoDOROSARIO/uber-go/internal/domain/model/ride_booking"

func NewRideBooker(
	tripScanner TripScanner,
	rideRepository RideRepository,
) *RideBooker {
	return &RideBooker{
		tripScanner:    tripScanner,
		rideRepository: rideRepository,
	}
}

type TripScanner interface {
	DistanceBetween(departure string, arrival string) int
}

type RideRepository interface {
	Save(ride *ride_booking.Ride)
}

type RideBooker struct {
	tripScanner    TripScanner
	rideRepository RideRepository
}

func (b *RideBooker) Execute(
	rideId string,
	riderId string,
	departure string,
	arrival string,
) {
	distance := b.tripScanner.DistanceBetween(departure, arrival)
	price := 30.0 + float64(distance)*0.5
	ride := ride_booking.Book(rideId, riderId, departure, arrival, price)
	b.rideRepository.Save(ride)
}
