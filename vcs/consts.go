package vcs

const (
	// RootDir is the directory where the VCS files will be stored
	RootDir = ".mirage/"
	// LocalDir is the directory that contains all of the local repository files
	LocalDir = RootDir + "local/"
	// ConfigPath is the VCS config path
	ConfigPath = RootDir + "config.json"
)
