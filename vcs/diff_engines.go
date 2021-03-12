package vcs

func (l lightway) Diff() (forward []string, backward []string) {
	return nil, nil
}

func (v voyage) Diff() (forward []string, backward []string) {
	return nil, nil
}

// DiffEngine is an engine that controls how commits are created and how
// differences between commits are viewed. Each engine is designed to support a
// different workload.
type DiffEngine interface {
	Diff(currentDump string, newDump string) (forward []string, backward []string)
}

type lightway struct{}
type voyage struct{}
