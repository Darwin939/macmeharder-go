package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"io"
	"os"
	"log"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Darwin939/macmeharder-go/graph/generated"
	"github.com/Darwin939/macmeharder-go/graph/model"
	"github.com/Darwin939/macmeharder-go/internal/apps"
)

func (r *mutationResolver) UploadAppAvatar(ctx context.Context, file graphql.Upload) (*model.AppAvatar, error) {
	filepath := "./public/avatars/"+file.Filename
	newfile, err := os.Create(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer newfile.Close()
	_, err = io.Copy(newfile, file.File)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Success!")
	
	return &model.AppAvatar{},nil

}

func (r *mutationResolver) CreateApp(ctx context.Context, input model.NewApp) (*model.App, error) {
	category, _ := strconv.Atoi(input.CategoryID)
	app := apps.App{
		Title:         input.Title,
		Language:      input.Language,
		LanguageCount: input.LanguageCount,
		Description:   input.Description,
		Size:          input.Size,
		Award:         input.Award,
		Developer:     input.Developer,
		Chart:         input.Chart,
		Version:       input.Version,
		Compatibility: input.Compatibility,
		DownloadURL:   input.DownloadURL,
		Category:      category,
		Age: input.Age,
	}
	app.Save()
	
	return &model.App{ Title: app.Title,
					Language:      app.Language,
					LanguageCount: app.LanguageCount,
					Description:   app.Description,
					Size:          app.Size,
					Award:         app.Award,
					Developer:     app.Developer,
					Chart:         app.Chart,
					Version:       app.Version,
					Compatibility: app.Compatibility,
					DownloadURL:   app.DownloadURL,
					Age: app.Age,
					}, nil
}

func (r *queryResolver) Apps(ctx context.Context) ([]*model.App, error) {
	var (
		resultApps []*model.App
		dbApps []apps.App
	)
	dbApps = apps.GetAll()
	for _,app := range dbApps {
		resultApps = append(resultApps,&model.App{
			ID: app.ID,
			Title: app.Title,
			Language:      app.Language,
			LanguageCount: app.LanguageCount,
			Description:   app.Description,
			Size:          app.Size,
			Award:         app.Award,
			Developer:     app.Developer,
			Chart:         app.Chart,
			Version:       app.Version,
			Compatibility: app.Compatibility,
			DownloadURL:   app.DownloadURL,
			Age: app.Age,} )
	}
	return resultApps,nil
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
