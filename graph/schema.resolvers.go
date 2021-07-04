package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Darwin939/macmeharder-go/graph/generated"
	"github.com/Darwin939/macmeharder-go/graph/model"
)

func (r *mutationResolver) UploadAppAvatar(ctx context.Context, file graphql.Upload) (*model.AppAvatar, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateApp(ctx context.Context, input model.NewApp) (*model.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Apps(ctx context.Context) ([]*model.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }


