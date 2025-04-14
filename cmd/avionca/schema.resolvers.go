package avionca

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

func (r *queryResolver) Flights(ctx context.Context, origin string, destination string, date string) ([]*Flight, error) {
	simulatedDelay := time.Duration(rand.Intn(600)+1120) * time.Millisecond
	time.Sleep(simulatedDelay)

	departureDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	today := time.Now().Truncate(24 * time.Hour)
	if !departureDate.After(today) {
		return []*Flight{}, nil
	}

	if origin == "" || destination == "" {
		return nil, errors.New("invalid origin and/or destination")
	}

	if origin == destination {
		return nil, errors.New("not supported same origin and destination")
	}

	if (origin != MIA && origin != SCL) || (destination != MIA && destination != SCL) {
		return []*Flight{}, nil
	}

	return flights(origin, destination, departureDate), nil
}

func (r *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	return "OK", nil
}

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
