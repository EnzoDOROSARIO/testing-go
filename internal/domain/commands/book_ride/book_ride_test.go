package book_ride_test

import (
	"testing"

	"github.com/EnzoDOROSARIO/uber-go/internal/domain/commands/book_ride"
	"github.com/EnzoDOROSARIO/uber-go/internal/domain/model/ride_booking"
	"github.com/stretchr/testify/assert"
)

func TestBasicRideBooking(t *testing.T) {
	bookRide := book_ride.NewRideBooker()

	actualRide := bookRide.Execute("rideId", "riderId", "PARIS_ADDRESS", "PARIS_ADDRESS")

	expectedSnapshot := ride_booking.RideSnapshot{
		ID:        "rideId",
		RiderId:   "riderId",
		Departure: "PARIS_ADDRESS",
		Arrival:   "PARIS_ADDRESS",
		Price:     30,
		Status:    "WAITING_FOR_DRIVER",
	}

	assert.Equal(t, expectedSnapshot, actualRide.ToSnapshot())
}
