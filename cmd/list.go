/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	// "text/tabwriter"

	"github.com/christianvozar/labelmaker/pkg/labelmaker"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

type label struct {
	Name  string
	Color string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all labels for repository",
	Long: `Retrives all labels associated with a Github repository

Ex: labelmaker list github/semantic`,
	Run: func(cmd *cobra.Command, args []string) {
		queryAllLabels(args[0])
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func queryAllLabels(r string) {
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	if r == "" {
		fmt.Println("No repository specified.")
		return
	}
	o, n := labelmaker.ParseRepository(r)
	variables := map[string]interface{}{
		"repositoryOwner": githubv4.String(o),
		"repositoryName":  githubv4.String(n),
		"commentsCursor":  (*githubv4.String)(nil), // Null after argument to get first page.
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		Repository struct {
			Labels struct {
				Nodes    []label
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"labels(first: 100, after: $commentsCursor)"` // 100 per page.
		} `graphql:"repository(owner: $repositoryOwner, name: $repositoryName)"`
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "\n%s\t%s", "Label", "Color")
	for _, l := range query.Repository.Labels.Nodes {
		// It would be wicked to support coloring the colors with their values using escape codes.
		fmt.Fprintf(w, "\n%s\t%s", l.Name, l.Color)
	}
	return
}
