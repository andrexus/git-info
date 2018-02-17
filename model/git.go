package model

import (
	"encoding/json"
	"time"

	"gopkg.in/yaml.v2"
)

type GitInfo struct {
	Branch string `json:"branch"`
	Commit Commit `json:"commit"`
}

type Commit struct {
	ID       string         `json:"id"`
	IDAbbrev string         `json:"idabbrev,omitempty" yaml:",omitempty"`
	Message  *CommitMessage `json:"message,omitempty" yaml:",omitempty"`
	User     *CommitUser    `json:"user,omitempty" yaml:",omitempty"`
	Time     time.Time      `json:"time,omitempty"`
}

type CommitMessage struct {
	Short string `json:"short"`
	Full  string `json:"full"`
}

type CommitUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *GitInfo) GetJSON() ([]byte, error) {
	bytes, err := json.MarshalIndent(c, "", "  ")
	return bytes, err
}

func (c *GitInfo) GetYAML() ([]byte, error) {
	bytes, err := yaml.Marshal(&c)
	return bytes, err
}
