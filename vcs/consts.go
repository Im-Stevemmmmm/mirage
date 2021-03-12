package vcs

const (
	// RootDir is the directory where the version control files are stored.
	RootDir = ".mirage/"
	// LocalDir is the directory where all of the files required for local operation are stored.
	LocalDir = RootDir + "local/"
	// ConfigPath is where the version control's config file is located at.
	ConfigPath = RootDir + "config.json"
)
