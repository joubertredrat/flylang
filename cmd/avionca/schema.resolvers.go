package avionca

import (
	"context"
)

func (r *queryResolver) Flights(ctx context.Context, origin string, destination string, date string) ([]*Flight, error) {
	return []*Flight{
		{
			FlightNumber: "1",
			Origin:       origin,
			Destination:  destination,
			Departure:    date,
			Duration:     300,
		},
	}, nil
}

func (r *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	return "OK", nil
}

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
