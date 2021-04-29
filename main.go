package main

import (
	"fmt"
	"log"
	"notify/notify"

	"github.com/spf13/viper"
)

var (
	REPO_URL     = "https://api.github.com/user/repos?page=1&per_page=1000"
	GITHUB_TOKEN = "419ed57e343e23a1dae83799e497fe0392c0ef22"
)

func main() {
	fmt.Println(`
	 ___ ____ ____  _   _ _____ 
	|_ _/ ___/ ___|| | | | ____|
	 | |\___ \___ \| | | |  _|  
	 | | ___) |__) | |_| | |___ 
	|___|____/____/ \___/|_____|
								
	 _   _  ___  _____ ___ ____    _  _____ ___ ___  _   _ 
	| \ | |/ _ \|  ___|_ _/ ___|  / \|_   _|_ _/ _ \| \ | |
	|  \| | | | | |_   | | |     / _ \ | |  | | | | |  \| |
	| |\  | |_| |  _|  | | |___ / ___ \| |  | | |_| | |\  |
	|_| \_|\___/|_|   |___\____/_/   \_\_| |___\___/|_| \_|
	
	`)
	res := notify.GetRepoList(REPO_URL)
	for _, r := range res {
		fmt.Println(r)
		notify.GetMileStoneInfo(r)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("protocols: ", viper.Get("config.dingding_url"))
	fmt.Println("protocols: ", viper.Get("config.sign"))
	// notify.SendDingDingMessage("test", viper.Get("config.dingding_url").(string), viper.Get("config.sign").(string))
}
