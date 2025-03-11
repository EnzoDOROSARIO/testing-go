package book_ride_test

import (
	"github.com/EnzoDOROSARIO/testing-go/internal/adapters/secondary/rider_repository"
	"testing"

	"github.com/EnzoDOROSARIO/testing-go/internal/adapters/secondary/trip_scanner"
	"github.com/EnzoDOROSARIO/testing-go/internal/domain/commands/book_ride"
	"github.com/EnzoDOROSARIO/testing-go/internal/domain/model/ride_booking"
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
		{
			name:          "Paris -> External | 0km",
			departure:     "PARIS_ADDRESS",
			arrival:       "EXTERNAL_ADDRESS",
			distance:      0,
			expectedPrice: 20.0,
		},
		{
			name:          "External -> Paris | 0km",
			departure:     "EXTERNAL_ADDRESS",
			arrival:       "PARIS_ADDRESS",
			distance:      0,
			expectedPrice: 50.0,
		},
		{
			name:          "External -> External | 0km",
			departure:     "EXTERNAL_ADDRESS",
			arrival:       "EXTERNAL_ADDRESS",
			distance:      0,
			expectedPrice: 100.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sampleRider := ride_booking.NewRider("riderId")

			riderRepository := rider_repository.NewFakeRiderRepository()
			tripScanner := trip_scanner.NewFakeTripScanner(tc.distance)
			bookRide := book_ride.NewRideBooker(tripScanner, riderRepository)

			riderRepository.Save(sampleRider)
			bookRide.Execute("riderId", tc.departure, tc.arrival)

			rider := riderRepository.ById("riderId")
			expectedSnapshot := ride_booking.RideSnapshot{
				ID:        "123abc",
				RiderId:   "riderId",
				Departure: tc.departure,
				Arrival:   tc.arrival,
				Price:     tc.expectedPrice,
				Status:    "WAITING_FOR_DRIVER",
			}
			assert.Equal(t, expectedSnapshot, rider.ToSnapshot().CurrentRide)
		})
	}
}
