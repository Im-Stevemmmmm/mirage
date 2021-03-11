package vcs

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"

	"github.com/shomali11/util/xhashes"
)

// CreateCommit creates a commit
func CreateCommit(title string, author string) error {
	sf, err := ioutil.ReadFile(StatePath)
	if err != nil {
		return err
	}

	var s State
	json.Unmarshal(sf, &s)

	b := s.getBranchByName(s.CurrentBranch)

	token := make([]byte, 4)
	rand.Read(token)
	h := xhashes.SHA256(string(token))

	b.Commits = append(b.Commits, Commit{
		Title:  title,
		Author: author,
		Hash:   h,
	})

	fs, _ := json.MarshalIndent(s, "", "  ")
	ioutil.WriteFile(StatePath, fs, 0644)

	return nil
}

// Commit is a snapshot of the database state
type Commit struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Hash   string `json:"Hash"`
}
