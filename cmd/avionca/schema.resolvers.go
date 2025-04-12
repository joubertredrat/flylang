package avionca

import (
	"context"
	"time"
)

func (r *queryResolver) Flights(ctx context.Context, origin string, destination string, date string) ([]*Flight, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	return flights(origin, destination, parsedDate), nil
}

func (r *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	return "OK", nil
}

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
