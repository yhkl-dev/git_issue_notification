package main

import (
	"fmt"
	"log"
	"notify/notify"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	DingdingURL    string
	SIGN           string
	RepoURL        string
	MileStoneURL   string
	GithubToken    string
	NotifyRepoList []string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	config = Config{
		DingdingURL:    viper.GetString("config.dingding_url"),
		SIGN:           viper.GetString("config.sign"),
		RepoURL:        viper.GetString("config.repo_url"),
		MileStoneURL:   viper.GetString("config.milestone_url"),
		GithubToken:    viper.GetString("config.github_token"),
		NotifyRepoList: viper.GetStringSlice("config.notify_repo_list"),
	}
}

func main() {
	fmt.Println(`
	___ ____ ____  _   _ _____   _   _  ___  _____ ___ ____    _  _____ ___ ___  _   _ 
	|_ _/ ___/ ___|| | | | ____| | \ | |/ _ \|  ___|_ _/ ___|  / \|_   _|_ _/ _ \| \ | |
	 | |\___ \___ \| | | |  _|   |  \| | | | | |_   | | |     / _ \ | |  | | | | |  \| |
	 | | ___) |__) | |_| | |___  | |\  | |_| |  _|  | | |___ / ___ \| |  | | |_| | |\  |
	|___|____/____/ \___/|_____| |_| \_|\___/|_|   |___\____/_/   \_\_| |___\___/|_| \_|
				  
	`)

	// notify.SendDingDingMessage("test", viper.Get("config.dingding_url").(string), viper.Get("config.sign").(string))
	fmt.Println(config.RepoURL)
	fmt.Println(config.GithubToken)
	res := notify.GetRepoList(config.RepoURL, config.GithubToken)
	for _, r := range res {
		fmt.Println(r)
	}
}
