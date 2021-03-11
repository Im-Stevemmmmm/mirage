package vcs

import (
	"encoding/json"
	"io/ioutil"
)

// CreateBranch creates a branch
func CreateBranch(name string) error {
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

	return nil
}

// CheckoutBranch switches branch
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

// Branch represents a branch
type Branch struct {
	Name    string   `json:"Name"`
	Commits []Commit `json:"Commits"`
}
