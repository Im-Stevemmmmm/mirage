package cli

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	statePath = rootDir + "state.json"
)

// CreateBranch creates a branch
func CreateBranch(name string) error {
	sf, err := ioutil.ReadFile(statePath)
	if err != nil {
		return err
	}

	var s state
	json.Unmarshal(sf, &s)

	s.Branches = append(s.Branches, branch{
		Name:    name,
		Commits: make([]commit, 0),
	})

	fs, _ := json.MarshalIndent(s, "", "  ")
	ioutil.WriteFile(statePath, fs, 0644)

	return nil
}

// CheckoutBranch switches branch
func CheckoutBranch(name string) error {
	sf, err := ioutil.ReadFile(statePath)
	if err != nil {
		return err
	}

	var s state
	json.Unmarshal(sf, &s)

	s.CurrentBranch = name

	fs, _ := json.MarshalIndent(s, "", "  ")
	ioutil.WriteFile(statePath, fs, 0644)

	return nil
}

func validateState() bool {
	os.Create(rootDir)
	return true
}
