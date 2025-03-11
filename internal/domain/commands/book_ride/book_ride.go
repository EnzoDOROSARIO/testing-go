package book_ride

import "github.com/EnzoDOROSARIO/testing-go/internal/domain/model/ride_booking"

func NewRideBooker(
	tripScanner TripScanner,
	riderRepository RiderRepository,
) *RideBooker {
	return &RideBooker{
		tripScanner:     tripScanner,
		riderRepository: riderRepository,
	}
}

type TripScanner interface {
	DistanceBetween(departure string, arrival string) int
	InParis(address string) bool
}

type RiderRepository interface {
	Save(ride *ride_booking.Rider)
	ById(id string) *ride_booking.Rider
}

type RideBooker struct {
	tripScanner     TripScanner
	riderRepository RiderRepository
}

func (b *RideBooker) Execute(
	riderId string,
	departure string,
	arrival string,
) {
	rider := b.riderRepository.ById(riderId)
	distance := b.tripScanner.DistanceBetween(departure, arrival)
	departureInParis := b.tripScanner.InParis(departure)
	arrivalInParis := b.tripScanner.InParis(arrival)
	basePrice := calculateBasePrice(departureInParis, arrivalInParis)
	price := basePrice + float64(distance)*0.5
	rider.Book("123abc", departure, arrival, price)
	b.riderRepository.Save(rider)
}

func calculateBasePrice(departureInParis bool, arrivalInParis bool) float64 {
	const (
		ParisToParisPrice       = 30.0
		ParisToExternalPrice    = 20.0
		ExternalToParisPrice    = 50.0
		ExternalToExternalPrice = 100.0
	)

	switch {
	case departureInParis && arrivalInParis:
		return ParisToParisPrice
	case departureInParis:
		return ParisToExternalPrice
	case arrivalInParis:
		return ExternalToParisPrice
	default:
		return ExternalToExternalPrice
	}
}
