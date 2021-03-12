package vcs

const (
	// StatePath is the location of the state file.
	StatePath = LocalDir + "state.json"
)

// State is the keeps track of commits and the data required to operate the
// version control system's branching and history features.
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
