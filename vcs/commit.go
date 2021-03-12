package vcs

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/shomali11/util/xhashes"
)

// CreateCommit creates a commit from the specified data and stores it in the
// local state.
func CreateCommit(title string, description string, author string) error {
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
		Title:       title,
		Dsecription: description,
		Author:      author,
		Hash:        h,
	})

	fs, _ := json.MarshalIndent(s, "", "  ")
	ioutil.WriteFile(StatePath, fs, 0644)

	return nil
}

// Commit is a snapshot of the database state.
type Commit struct {
	Title       string `json:"Title"`
	Dsecription string `json:"Description"`
	Author      string `json:"Author"`
	Hash        string `json:"Hash"`
}

type revertMethod int

const (
	// Soft sets the HEAD pointer to point the commit SHA.
	Soft revertMethod = iota
	// Hard sets the HEAD pointer to point the commit SHA and reverts changes.
	Hard
)

func (r revertMethod) String() string {
	return [...]string{"soft", "hard"}[r]
}

// RevertTo reverts to a commit.
func (c Commit) RevertTo(m revertMethod) error {
	switch m {
	case Soft:
		return nil
	case Hard:
		return nil
	default:
		return errors.New("Unrecognized revertMethod")
	}
}
