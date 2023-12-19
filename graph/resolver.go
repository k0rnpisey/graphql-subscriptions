package graph

import "gqlgen-subscriptions/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserStore                map[string]*model.User
	NotificationStore        map[string][]*model.Notification
	NotificationSubscription map[string]chan *model.Notification
}
