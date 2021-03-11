package cli

type state struct {
	CurrentBranch string   `json:"CurrentBranch"`
	Branches      []branch `json:"Branches"`
}

func (s state) getBranchByName(name string) *branch {
	branches := s.Branches
	for i := range branches {
		b := branches[i]
		if b.Name == name {
			return &s.Branches[i]
		}
	}
	return nil
}

type branch struct {
	Name    string   `json:"Name"`
	Commits []commit `json:"Commits"`
}

type commit struct {
	Hash   string `json:"Hash"`
	Author string `json:"Author"`
	// Time   string `json:"Time"`
}
