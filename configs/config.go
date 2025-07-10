package configs

var (
	MainDir = ".trace"
	SubDirs = []string{
		"objects", "refs",
	}
	IgnoredDirs = map[string]bool{
		".git":   true,
		".trace": true,
	}
)
