package ride_repository

import "github.com/EnzoDOROSARIO/testing-go/internal/domain/model/ride_booking"

type FakeRideRepository struct {
	Rides []ride_booking.Ride
}

func NewFakeRideRepository() *FakeRideRepository {
	return &FakeRideRepository{
		Rides: []ride_booking.Ride{},
	}
}

func (r *FakeRideRepository) Save(ride *ride_booking.Ride) {
	r.Rides = append(r.Rides, *ride)
}
