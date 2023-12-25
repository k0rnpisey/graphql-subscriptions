package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"github.com/edgedb/edgedb-go"
	edgedbmodel "gqlgen-subscriptions/dbschema/model"
	"gqlgen-subscriptions/graph/model"
	"strconv"
)

// UpsertUser is the resolver for the upsertUser field.
func (r *mutationResolver) UpsertUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	var user model.User
	n := len(r.Resolver.UserStore)
	id := input.ID

	if id != nil {
		u, ok := r.Resolver.UserStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		u.Name = input.Name
		user = *u
	} else {
		// generate unique id
		nid := strconv.Itoa(n + 1)
		user.ID = nid
		user.Name = input.Name
		r.Resolver.UserStore[nid] = &user
	}
	return &user, nil
}

// FollowUser is the resolver for the followUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, userID string, followingUserID string) (*model.User, error) {
	// Find the user and following in the data store.
	user, userOk := r.UserStore[userID]
	following, followingOk := r.UserStore[followingUserID]

	// If the user or following was not found, return an error.
	if !userOk {
		return nil, fmt.Errorf("user with ID %s not found", userID)
	}
	if !followingOk {
		return nil, fmt.Errorf("following with ID %s not found", followingUserID)
	}

	// Add the following to the user's list of friends.
	user.Following = append(user.Following, following)

	// Update the user in the data store.
	r.UserStore[userID] = user

	listener := r.NotificationSubscription[followingUserID]
	notification := model.Notification{
		ID:      strconv.Itoa(len(r.NotificationStore) + 1),
		Type:    model.NotificationTypeFollower,
		Message: fmt.Sprintf("%s is now following you", user.Name),
	}
	// append to the notification store for the user
	r.NotificationStore[followingUserID] = append(r.NotificationStore[followingUserID], &notification)
	//r.NotificationStore[followingUserID] = &notification

	if listener != nil {
		listener <- &notification
	}
	// Return the updated user.
	return user, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.Post, error) {
	var inserted struct{ id edgedb.UUID }
	query := fmt.Sprintf(`
    INSERT Post {
        title := '%s',
        content := '%s'
    }
`, input.Title, input.Content)
	err := r.Db.QuerySingle(ctx, query, &inserted)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		ID:      inserted.id.String(),
		Title:   input.Title,
		Content: input.Content,
	}, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.Post, error) {
	var post edgedbmodel.Post
	query := `SELECT Post { id, title, content } FILTER .id = <uuid>$0`
	uuid, _ := edgedb.ParseUUID(input.ID)
	err := r.Db.QuerySingle(ctx, query, &post, uuid)
	if err != nil {
		return nil, err
	}
	post.Title = input.Title
	post.Content = input.Content
	query = `UPDATE Post filter .id = <uuid>$0 SET { title := <str>$1, content := <str>$2 };`
	err = r.Db.Execute(ctx, query, uuid, input.Title, input.Content)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		ID:      post.Id.String(),
		Title:   post.Title,
		Content: post.Content,
		Author:  &model.User{},
	}, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	query := `DELETE Post FILTER .id = <uuid>$0;`
	uuid, _ := edgedb.ParseUUID(id)
	err := r.Db.Execute(ctx, query, uuid)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Placeholder is the resolver for the placeholder field.
func (r *queryResolver) Placeholder(ctx context.Context) (*string, error) {
	str := "Hello World"
	return &str, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	if len(r.Resolver.UserStore) > 0 {
		users := make([]*model.User, 0)
		// eager load all users and their friends (if any) from the data store and return them.
		for idx := range r.Resolver.UserStore {
			user := r.Resolver.UserStore[idx]
			users = append(users, user)
		}
		return users, nil
	}
	return []*model.User{}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, email string, password string) (*model.User, error) {
	if len(r.Resolver.UserStore) > 0 {
		// eager load all users and their friends (if any) from the data store and return them.
		for idx := range r.Resolver.UserStore {
			user := r.Resolver.UserStore[idx]
			if user.Name == email {
				return user, nil
			}
		}
	}
	return nil, nil
}

// UserNotifications is the resolver for the userNotifications field.
func (r *queryResolver) UserNotifications(ctx context.Context, userID string) ([]*model.Notification, error) {
	store := r.Resolver.NotificationStore[userID]
	if store != nil {
		return store, nil
	}
	return []*model.Notification{}, nil
}

// Notifications is the resolver for the notifications field.
func (r *queryResolver) Notifications(ctx context.Context) ([]*model.Notification, error) {
	if len(r.Resolver.NotificationStore) > 0 {
		notifications := make([]*model.Notification, 0)
		// eager load all users and their friends (if any) from the data store and return them.
		for idx := range r.Resolver.NotificationStore {
			notification := r.Resolver.NotificationStore[idx]
			notifications = append(notifications, notification...)
		}
		return notifications, nil
	}
	return []*model.Notification{}, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var posts []edgedbmodel.Post
	query := `SELECT Post { id, title, content }`
	err := r.Db.Query(ctx, query, &posts)
	if err != nil {
		return nil, err
	}
	var out []*model.Post
	for _, post := range posts {
		o := model.Post{
			ID:      post.Id.String(),
			Title:   post.Title,
			Content: post.Content,
			Author:  &model.User{},
		}
		out = append(out, &o)
	}
	return out, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	var post edgedbmodel.Post
	query := `SELECT Post { id, title, content } FILTER .id = <uuid>$0`
	uuid, _ := edgedb.ParseUUID(id)
	err := r.Db.QuerySingle(ctx, query, &post, uuid)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		ID:      post.Id.String(),
		Title:   post.Title,
		Content: post.Content,
		Author:  &model.User{},
	}, nil
}

// Notification is the resolver for the notification field.
func (r *subscriptionResolver) Notification(ctx context.Context, userID string) (<-chan *model.Notification, error) {
	updates := make(chan *model.Notification, 10) // FIXME: 10 is arbitrary
	c := make(chan *model.Notification, 10)
	r.Resolver.NotificationSubscription[userID] = c
	go func() {
		defer close(updates)
		for {
			select {
			case <-ctx.Done():
				// If the client disconnects, stop the goroutine.
				return
			case followerEvent := <-c:
				// We received a user update event. Send the update to the client.
				updates <- followerEvent
			}
		}
	}()
	return updates, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
