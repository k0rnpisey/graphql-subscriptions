package graph

import (
	"github.com/edgedb/edgedb-go"
	"gqlgen-subscriptions/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db                       *edgedb.Client
	NotificationSubscription map[string]chan *model.Notification
}
