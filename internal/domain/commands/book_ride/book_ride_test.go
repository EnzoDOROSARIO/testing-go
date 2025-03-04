package book_ride_test

import (
	"testing"

	"github.com/EnzoDOROSARIO/uber-go/internal/adapters/secondary/trip_scanner"
	"github.com/EnzoDOROSARIO/uber-go/internal/domain/commands/book_ride"
	"github.com/EnzoDOROSARIO/uber-go/internal/domain/model/ride_booking"
	"github.com/stretchr/testify/assert"
)

func TestBasicRideBooking(t *testing.T) {
	testCases := []struct {
		distance      int
		expectedPrice float64
	}{
		{distance: 0, expectedPrice: 30},
		{distance: 1, expectedPrice: 30.5},
	}

	for _, tc := range testCases {
		tripScanner := trip_scanner.NewFakeTripScanner(tc.distance)
		bookRide := book_ride.NewRideBooker(tripScanner)

		actualRide := bookRide.Execute("rideId", "riderId", "PARIS_ADDRESS", "PARIS_ADDRESS")

		expectedSnapshot := ride_booking.RideSnapshot{
			ID:        "rideId",
			RiderId:   "riderId",
			Departure: "PARIS_ADDRESS",
			Arrival:   "PARIS_ADDRESS",
			Price:     tc.expectedPrice,
			Status:    "WAITING_FOR_DRIVER",
		}

		assert.Equal(t, expectedSnapshot, actualRide.ToSnapshot())
	}
}
