package book_ride_test

import (
	"testing"

	"github.com/EnzoDOROSARIO/uber-go/internal/adapters/secondary/ride_repository"
	"github.com/EnzoDOROSARIO/uber-go/internal/adapters/secondary/trip_scanner"
	"github.com/EnzoDOROSARIO/uber-go/internal/domain/commands/book_ride"
	"github.com/EnzoDOROSARIO/uber-go/internal/domain/model/ride_booking"
	"github.com/stretchr/testify/assert"
)

func TestBasicRideBooking(t *testing.T) {
	testCases := []struct {
		name          string
		departure     string
		arrival       string
		distance      int
		expectedPrice float64
	}{
		{
			name:          "Paris -> Paris | 0km",
			departure:     "PARIS_ADDRESS",
			arrival:       "PARIS_ADDRESS",
			distance:      0,
			expectedPrice: 30.0,
		},
		{
			name:          "Paris -> Paris | 1km",
			departure:     "PARIS_ADDRESS",
			arrival:       "PARIS_ADDRESS",
			distance:      1,
			expectedPrice: 30.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rideRepository := ride_repository.NewFakeRideRepository()
			tripScanner := trip_scanner.NewFakeTripScanner(tc.distance)
			bookRide := book_ride.NewRideBooker(tripScanner, rideRepository)

			bookRide.Execute("rideId", "riderId", tc.departure, tc.arrival)

			actualRide := rideRepository.Rides[0]

			expectedSnapshot := ride_booking.RideSnapshot{
				ID:        "rideId",
				RiderId:   "riderId",
				Departure: tc.departure,
				Arrival:   tc.arrival,
				Price:     tc.expectedPrice,
				Status:    "WAITING_FOR_DRIVER",
			}
			assert.Equal(t, expectedSnapshot, actualRide.ToSnapshot())
		})
	}
}
