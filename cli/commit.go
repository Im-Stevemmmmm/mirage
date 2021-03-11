package cli

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"

	"github.com/shomali11/util/xhashes"
)

// CreateCommit creates a commit
func CreateCommit(title string) error {
	sf, err := ioutil.ReadFile(statePath)
	if err != nil {
		return err
	}

	var s state
	json.Unmarshal(sf, &s)

	b := s.getBranchByName(s.CurrentBranch)

	token := make([]byte, 4)
	rand.Read(token)
	h := xhashes.SHA256(string(token))

	author := "Stevemmmmm"

	b.Commits = append(b.Commits, commit{
		Hash:   h,
		Author: author,
	})

	fs, _ := json.MarshalIndent(s, "", "  ")
	ioutil.WriteFile(statePath, fs, 0644)

	return nil
}
