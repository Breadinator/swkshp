package main

import (
	"github.com/breadinator/swkshp/config"
	"github.com/breadinator/swkshp/utils"
	"github.com/tcnksm/go-latest"
)

func checkUpdate() {
	// Check for if there is a new version
	githubTag := &latest.GithubTag{
		Owner:             "breadinator",
		Repository:        "swkshp",
		FixVersionStrFunc: latest.DeleteFrontV(),
	}
	res, err := latest.Check(githubTag, VERSION[1:])
	if err != nil {
		utils.Err(err)
	} else if res.Outdated {
		utils.Warn("%s is not the latest version, you should upgrade to v%s", VERSION, res.Current)
	}

	// updates the config version to be what is listed in main.go
	if config.Conf.Main.Version != VERSION {
		config.Conf.Main.Version = VERSION
		config.SaveConfig(config.Conf)
	}
}
