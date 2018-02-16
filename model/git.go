package model

import "time"

type GitInfo struct {
	Branch string `json:"branch"`
	Commit Commit `json:"commit"`
}

type Commit struct {
	ID       string        `json:"id"`
	IDAbbrev string        `json:"id.abbrev,omitempty"`
	Message  *CommitMessage `json:"message,omitempty"`
	User     *CommitUser    `json:"user,omitempty"`
	Time     time.Time     `json:"time"`
}

type CommitMessage struct {
	Short string `json:"short"`
	Full  string `json:"full"`
}

type CommitUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
