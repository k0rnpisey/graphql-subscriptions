package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-subscriptions/graph/model"
	"time"
)

// Placeholder is the resolver for the placeholder field.
func (r *queryResolver) Placeholder(ctx context.Context) (*string, error) {
	str := "Hello World"
	return &str, nil
}

// CurrentTime is the resolver for the currentTime field.
func (r *subscriptionResolver) CurrentTime(ctx context.Context) (<-chan *model.Time, error) {
	ch := make(chan *model.Time)

	go func() {
		defer close(ch)

		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Tick")

			currentTime := time.Now()

			t := &model.Time{
				UnixTime:  int(currentTime.Unix()),
				TimeStamp: currentTime.Format(time.RFC3339),
			}

			select {
			case <-ctx.Done():
				// Exit on cancellation
				fmt.Println("Subscription closed.")
				return

			case ch <- t:
				// Our message went through, do nothing
			}

		}
	}()
	return ch, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
