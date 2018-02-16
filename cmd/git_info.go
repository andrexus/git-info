package cmd

import (
	"encoding/json"
	"log"
	"os"

	"fmt"

	"github.com/andrexus/git-info/model"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"strings"
)

var gitInfoCmd = cobra.Command{
	Run:   getGitInfo,
	Use:   "info",
	Short: "Get git info",
	Long:  "Show current git state",
}

func getGitInfo(cmd *cobra.Command, args []string) {
	mode, err := cmd.PersistentFlags().GetString("mode")
	if mode != "short" {
		mode = "full"
	}
	if err != nil {
		log.Fatalf("%v", err)
	}
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(r)
	}

	ref, err := r.Head()
	if err != nil {
		log.Fatal(r)
	}

	commit, err := r.CommitObject(ref.Hash())

	output := &model.GitInfo{
		Branch: ref.Name().Short(),
		Commit: model.Commit{
			ID:   commit.Hash.String(),
			Time: commit.Author.When,
		},
	}

	if mode == "full" {
		output.Commit.IDAbbrev = commit.Hash.String()[:7]
		output.Commit.User = &model.CommitUser{
			Name:  commit.Author.Name,
			Email: commit.Author.Email,
		}
		output.Commit.Message = &model.CommitMessage{
			Short: strings.Split(commit.Message, "\n")[0],
			Full:  strings.TrimSpace(strings.Replace(commit.Message, "\n", " ", -1)),
		}
	}

	b, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}
