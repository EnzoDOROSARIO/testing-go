package book_ride_test

import (
	"testing"

	"github.com/EnzoDOROSARIO/uber-go/internal/domain/commands/book_ride"
	"github.com/stretchr/testify/assert"
)

func TestBasicRideBooking(t *testing.T) {
	expectedRide := "ride"
	bookRide := book_ride.NewRideBooker()

	actualRide := bookRide.Execute()

	assert.Equal(t, expectedRide, actualRide)
}
