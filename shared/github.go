package shared

import (
	"os"
	"log"
	"fmt"
	// "bufio"
	// "strings"
	"syscall"
	"context"

	"golang.org/x/term"
	"github.com/joho/godotenv"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func MainGitHub() (context.Context, *github.Client, error){
	//call github utilities
	homeDir,_ := os.UserHomeDir()
	path := homeDir + "/.k3ai/" 
	err := godotenv.Load(path + ".env")
	if err != nil {
		// reader := bufio.NewReader(os.Stdin)
		
		fmt.Println("Missing GitHub authentication token, please paste it here: ")
		bytePassword, _:= term.ReadPassword(int(syscall.Stdin))
		token := string(bytePassword)
		if token != "" {
			ctx,client,_ := login(token)
			return ctx, client, nil
			 }
	}
	if err == nil {
	//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
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	 if token != "" {
		ctx,client,_ := login(token)
		return ctx, client, nil
	 	}
	
	}
	return nil,nil,nil
}

func login(token string) (context.Context, *github.Client, error) {

	if token == "" {
		log.Fatal("please provide a GitHub API token via env variable GITHUB_AUTH_TOKEN")
	}

	ctx, client, err := githubAuth(token)
	if err != nil {
		log.Fatalf("unable to authorize using env GITHUB_AUTH_TOKEN: %v", err)
	}
	return ctx, client, nil
}

// githubAuth returns a GitHub client and context.
func githubAuth(token string) (context.Context, *github.Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return ctx, client, nil
}

