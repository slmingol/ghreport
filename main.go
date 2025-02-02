package main

import (
	"fmt"
	"os"
)

func main() {
	subscribedRepos, envToken, err := getEnvVariables()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := createGithubClient(envToken)

	// List of repos to watch
	for _, ownerRepo := range subscribedRepos {
		owner, repo, err := getOwnerAndRepo(ownerRepo)
		if err != nil {
			fmt.Println(err)
			continue
		}

		pullRequests, err := getPrFromRepo(client, owner, repo)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, pr := range pullRequests {
			fmt.Printf("%s: createdAt %s\n", pr.URL, pr.CreatedAt)
		}
	}
}
