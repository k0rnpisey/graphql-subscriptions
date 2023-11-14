package graph

import "gqlgen-subscriptions/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserStore map[string]*model.User
	// 2 is a channel that will be used to push new users to subscriptions
	UserUpdateEvents chan *model.User
}
