package mygh

import (
	"context"
	"fmt"

	"github.com/google/go-github/v52/github"
)

func Repositories(ctx context.Context, name string) ([]*github.Repository, error) {
	client := github.NewClient(nil)

	repos, response, err := client.Repositories.List(ctx, name, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list repositories: %w", err)
	}
	defer response.Body.Close()

	return repos, nil
}
