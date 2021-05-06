package notify

import (
	"fmt"
	"log"

	"github.com/bitly/go-simplejson"
)

type Github struct {
	GithubToken string
}

func NewGithub(githubToken string) *Github {
	return &Github{GithubToken: githubToken}
}

func contains(repo_name string, repo_list []string) bool {
	for _, element := range repo_list {
		if repo_name == element {
			return true
		}
	}
	return false
}

func (g *Github) GetRepoList(url string, notifyRepoList []string) []string {
	repoList := []string{}
	res := HTTPRequest(url, g.GithubToken)
	arr, err := res.Array()
	if err != nil {
		log.Fatal(err)
	}
	if arr == nil {
		log.Fatal("something wrong when call Get and Array")
	}
	for _, content := range arr { 
		newdi, _ := content.(map[string]interface{})
		fullName := fmt.Sprintf("%s", newdi["full_name"])
		if contains(fullName, notifyRepoList) {
			repoList = append(repoList, fullName)
		}
	}
	return repoList
}

func (g *Github) GetLatestMilestone(mileStoneURL string, repoNameList []string) {
	mileStoneList := []string{}
	for _, repo := range repoNameList {
		mileStoneList = append(mileStoneList, fmt.Sprintf(mileStoneURL, repo))
	}
	for _, m := range mileStoneList {
		arr, err := g.GetMileStoneInfo(m).Array()
		if err != nil {
			fmt.Println(err)
		}
		for _, x := range arr {
			fmt.Println(x)
		}
	}

}

func (g *Github) GetMileStoneInfo(mileStoneURL string) *simplejson.Json {
	res := HTTPRequest(mileStoneURL, g.GithubToken)
	return res
}

// def get_latest_milestone(repo_name_list):
//     milestone_list = []
//     for repo_name in repo_name_list:
//         milestone_list.extend(get_milestone_info(repo_name))

//     milestone_list.sort(key=lambda x: x.get("created_at"), reverse=True)

//     print("The latest milestone is: {}".format(milestone_list[0].get("title")))

//     todo_milestone = [milestone for milestone in milestone_list if
//                       milestone.get("title") == milestone_list[0].get("title")]

//     return milestone_list[0], todo_milestone
