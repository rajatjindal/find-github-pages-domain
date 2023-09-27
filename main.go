package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v55/github"
)

func main() {
	token := "<token>"
	username := "rajatjindal"

	client := github.NewClient(nil).WithAuthToken(token)

	for i := 0; i <= 5; i++ {
		repos, _, err := client.Repositories.List(context.Background(), username, &github.RepositoryListOptions{ListOptions: github.ListOptions{Page: i, PerPage: 100}})
		if err != nil {
			panic(err)
		}

		fmt.Println(len(repos))
		for _, repo := range repos {
			fmt.Println(repo.GetName())
			pages, resp, err := client.Repositories.GetPagesInfo(context.TODO(), username, repo.GetName())
			if resp.StatusCode == http.StatusNotFound {
				continue
			}

			if err != nil {
				panic(err)
			}

			fmt.Printf("repo: %s, domain: %s\n", repo.GetName(), pages.GetCNAME())
		}
	}

}
