package rider_repository

import "github.com/EnzoDOROSARIO/testing-go/internal/domain/model/ride_booking"

type FakeRiderRepository struct {
	Riders []ride_booking.Rider
}

func NewFakeRiderRepository() *FakeRiderRepository {
	return &FakeRiderRepository{
		Riders: []ride_booking.Rider{},
	}
}

func (r *FakeRiderRepository) Save(rider *ride_booking.Rider) {
	for i, existingRider := range r.Riders {
		if existingRider.ID() == rider.ID() {
			r.Riders[i] = *rider
			return
		}
	}
	r.Riders = append(r.Riders, *rider)
}

func (r *FakeRiderRepository) ById(id string) *ride_booking.Rider {
	for _, rider := range r.Riders {
		if rider.ID() == id {
			return &rider
		}
	}
	return nil
}
