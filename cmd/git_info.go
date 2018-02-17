package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andrexus/git-info/model"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var gitInfoCmd = cobra.Command{
	Run:   getGitInfo,
	Use:   "info",
	Short: "Get git info",
	Long:  "Show current git state",
}

//noinspection GoUnusedParameter
func getGitInfo(cmd *cobra.Command, args []string) {
	path, err := os.Getwd()
	checkError(err)

	r, err := git.PlainOpen(path)
	checkError(err)
	ref, err := r.Head()
	checkError(err)
	commit, err := r.CommitObject(ref.Hash())
	checkError(err)

	mode, err := cmd.PersistentFlags().GetString("mode")
	checkError(err)

	if mode != "short" {
		mode = "full"
	}

	i := gitInfo(mode, ref, commit)

	format, err := cmd.PersistentFlags().GetString("out")
	checkError(err)

	var data []byte
	if format == "yaml" {
		data, err = i.GetYAML()
		checkError(err)
	} else {
		data, err = i.GetJSON()
		checkError(err)
	}

	f, err := cmd.PersistentFlags().GetString("file")
	checkError(err)

	if f != "" {
		writeFile(f, data)
	} else {
		fmt.Printf("%s\n", data)
	}
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func gitInfo(mode string, ref *plumbing.Reference, commit *object.Commit) model.GitInfo {
	i := model.GitInfo{
		Branch: ref.Name().Short(),
		Commit: model.Commit{
			ID:   commit.Hash.String(),
			Time: commit.Author.When,
		},
	}

	if mode == "full" {
		i.Commit.IDAbbrev = commit.Hash.String()[:7]
		i.Commit.User = &model.CommitUser{
			Name:  commit.Author.Name,
			Email: commit.Author.Email,
		}
		i.Commit.Message = &model.CommitMessage{
			Short: strings.Split(commit.Message, "\n")[0],
			Full:  commit.Message,
		}
	}
	return i
}

func writeFile(path string, data []byte) {
	f, err := os.Create(path)
	checkError(err)
	defer f.Close()

	_, err = f.Write(data)
	checkError(err)
	f.Sync()
}
