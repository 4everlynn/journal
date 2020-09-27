package config

// Git to describe Repo's structure
type Git struct {
	Name    string
	Path    string
	Disable bool
}

// JournalConfig config to describe .journal.yaml
type JournalConfig struct {
	Version    string
	Maintainer string
	Git        map[string]Git
}
