package main

import (
	"fmt"
	"log"
	"notify/notify"

	"github.com/spf13/viper"
)

type Config struct {
	DingDingURL    string
	Sign           string
	RepoURL        string
	MilestoneURL   string
	GithubToken    string
	NotifyRepoList []string
}

var c Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.Unmarshal(&c)
}

func main() {
	fmt.Println(`
	___ ____ ____  _   _ _____   _   _  ___  _____ ___ ____    _  _____ ___ ___  _   _ 
	|_ _/ ___/ ___|| | | | ____| | \ | |/ _ \|  ___|_ _/ ___|  / \|_   _|_ _/ _ \| \ | |
	 | |\___ \___ \| | | |  _|   |  \| | | | | |_   | | |     / _ \ | |  | | | | |  \| |
	 | | ___) |__) | |_| | |___  | |\  | |_| |  _|  | | |___ / ___ \| |  | | |_| | |\  |
	|___|____/____/ \___/|_____| |_| \_|\___/|_|   |___\____/_/   \_\_| |___\___/|_| \_|
				  
	`)

	github := notify.NewGithub(c.GithubToken, c.NotifyRepoList)
	github.Rank(c.RepoURL, c.MilestoneURL)
}
