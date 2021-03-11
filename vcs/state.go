package vcs

const (
	// StatePath is where the state file is located
	StatePath = LocalDir + "state.json"
)

// State is the local repository state
type State struct {
	CurrentBranch string   `json:"CurrentBranch"`
	Branches      []Branch `json:"Branches"`
}

func (s State) getBranchByName(name string) *Branch {
	branches := s.Branches
	for i := range branches {
		b := branches[i]
		if b.Name == name {
			return &s.Branches[i]
		}
	}
	return nil
}
