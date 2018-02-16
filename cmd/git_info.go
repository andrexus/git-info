package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/andrexus/git-info/model"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"os"
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
	checkError(err)

	if mode != "short" {
		mode = "full"
	}
	path, err := os.Getwd()
	checkError(err)

	r, err := git.PlainOpen(path)
	checkError(err)

	ref, err := r.Head()
	checkError(err)

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

	bytes, err := json.MarshalIndent(output, "", "  ")
	checkError(err)

	out, err := cmd.PersistentFlags().GetString("out")
	checkError(err)

	if out != "" {
		f, err := os.Create(out)
		checkError(err)
		defer f.Close()

		_, err = f.Write(bytes)
		checkError(err)
		f.Sync()
	} else {
		fmt.Printf("%s\n", bytes)
	}
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
