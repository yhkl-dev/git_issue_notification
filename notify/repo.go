package notify

import (
	"fmt"
	"log"
)

var (
	REPO_URL         = "https://api.github.com/user/repos?page=1&per_page=1000"
	GITHUB_TOKEN     = "419ed57e343e23a1dae83799e497fe0392c0ef22"
	MILESTONE_URL    = "https://api.github.com/repos/%s/milestones?state=open"
	NOTIFY_REPO_LIST = []string{
		"YicunWendyWu/fire-risk-client",
		"YicunWendyWu/fire-inspection-client",
	}
)

func contains(repo_name string, repo_list []string) bool {
	for _, element := range repo_list {
		if repo_name == element {
			return true
		}
	}
	return false
}

func GetRepoList(url string) []string {
	repoList := []string{}
	res := HTTPRequest(url, GITHUB_TOKEN)
	arr, err := res.Array()
	if err != nil || arr == nil {
		log.Fatal("something wrong when call Get and Array")
	}
	for _, content := range arr { //就在这里对di进行类型判断
		newdi, _ := content.(map[string]interface{})
		fullName := fmt.Sprintf("%s", newdi["full_name"])
		if contains(fullName, NOTIFY_REPO_LIST) {
			repoList = append(repoList, fullName)
		}
	}
	return repoList
}

func GetMileStoneInfo(repoName string) {
	mileStoneURL := fmt.Sprintf(MILESTONE_URL, repoName)
	fmt.Println(mileStoneURL)
}

func GetLatestMilestone(repoNameList []string) {
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