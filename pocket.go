package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/motemen/go-pocket/api"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/motemen/go-pocket/auth"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
)

var cmdPocket = &cobra.Command{
	Use:   "_pocket",
	Short: "Send to Pocket",
	Long:  "Send articles to your Pocket",
}

func init() {
	cmdPocket.Run = Pocket
}

func Pocket(
	cmd *cobra.Command, args []string) {

	const consumerKey = "41271-5803fea25a99fa750709a512"

	// todo: restore access token from saved config file, rather
	//       than requiring authorization every time
	accessToken, err := obtainAccessToken(consumerKey)
	if err != nil {
		panic(err)
	}

	client := api.NewClient(consumerKey, accessToken.AccessToken)

	posts, err := ReadPosts(os.Stdin)

	fmt.Println("Adding these articles to Pocket...")
	for _, post := range posts {
		// append berliner tag for easy Pocket searching later
		tags := append(post.Tags, "berliner")
		tagsStr := strings.Join(tags, ",")

		options := &api.AddOption{
			URL:   post.Permalink,
			Title: post.Title,
			Tags:  tagsStr,
		}

		err := client.Add(options)
		if err != nil {
			panic(err)
		}

		fmt.Println(post.Title)
	}
}

func obtainAccessToken(consumerKey string) (*auth.Authorization, error) {
	ch := make(chan struct{})
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/favicon.ico" {
				http.Error(w, "Not Found", 404)
				return
			}

			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintln(w, "Authorized. Please return to your shell.")
			ch <- struct{}{}
		}))
	defer ts.Close()

	redirectURL := ts.URL

	requestToken, err := auth.ObtainRequestToken(consumerKey, redirectURL)
	if err != nil {
		return nil, err
	}

	url := auth.GenerateAuthorizationURL(requestToken, redirectURL)

	fmt.Println("If your web browser doesn't automatically open, copy paste this URL to authorize Pocket access.")
	fmt.Println(url)

	// Use open command to automatically open authorization page
	openCmd := exec.Command("open", url)
	openCmd.Start()

	<-ch

	return auth.ObtainAccessToken(consumerKey, requestToken)
}
