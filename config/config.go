package config

var (
	Version       = "0.1.0"
	WireCmd       = "github.com/google/wire/cmd/wire@latest"
	TruxCmd       = "github.com/truxcoder/trux@latest"
	RepoBase      = "https://github.com/truxcoder/trux-layout-basic.git"
	RepoAdvanced  = "https://github.com/truxcoder/trux-layout-advanced.git"
	RunExcludeDir = ".git,.idea,tmp,vendor"
	RunIncludeExt = "go,html,yaml,yml,toml,ini,json,xml,tpl,tmpl"
	Logo          = `
________________ _____  ______  __
___  __/___  __ \__  / / /__  |/ /
__  /   __  /_/ /_  / / / __    / 
_  /    _  _, _/ / /_/ /  _    |  
/_/     /_/ |_|  \____/   /_/|_|  
`
)
