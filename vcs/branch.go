package vcs

import (
	"encoding/json"
	"io/ioutil"
)

// CreateBranch creates a new branch.
func CreateBranch(name string, checkout bool) error {
	sf, err := ioutil.ReadFile(StatePath)
	if err != nil {
		return err
	}

	var s State
	json.Unmarshal(sf, &s)

	s.Branches = append(s.Branches, Branch{
		Name:    name,
		Commits: make([]Commit, 0),
	})

	fs, _ := json.MarshalIndent(s, "", "  ")
	ioutil.WriteFile(StatePath, fs, 0644)

	if checkout {
		CheckoutBranch(name)
	}

	return nil
}

// CheckoutBranch switches the user's active branch.
func CheckoutBranch(name string) error {
	sf, err := ioutil.ReadFile(StatePath)
	if err != nil {
		return err
	}

	var s State
	json.Unmarshal(sf, &s)

	s.CurrentBranch = name

	fs, _ := json.MarshalIndent(s, "", "  ")
	ioutil.WriteFile(StatePath, fs, 0644)

	return nil
}

// Branch is a group of independent commits.
type Branch struct {
	Name    string   `json:"Name"`
	Commits []Commit `json:"Commits"`
}

// Merge merges a branch with another branch.
func (b Branch) Merge(target Branch) {

}
