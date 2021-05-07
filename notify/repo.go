package notify

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
)

type Github struct {
	Token          string
	NotifyRepoList []string
}

func NewGithub(token string, repoList []string) *Github {
	return &Github{Token: token, NotifyRepoList: repoList}
}

func contains(repo_name string, repo_list []string) bool {
	for _, element := range repo_list {
		if repo_name == element {
			return true
		}
	}
	return false
}

func (g *Github) GetRepoList(url string) []string {
	repoList := []string{}
	res := HTTPRequest(url, g.Token)
	arr, err := res.Array()
	if err != nil || arr == nil {
		log.Fatal("something wrong when call Get and Array")
	}
	for _, content := range arr { //就在这里对di进行类型判断
		newdi, _ := content.(map[string]interface{})
		fullName := fmt.Sprintf("%s", newdi["full_name"])
		if contains(fullName, g.NotifyRepoList) {
			repoList = append(repoList, fullName)
		}
	}
	return repoList
}

func (g *Github) GetLatestMilestone(mileStoneURL string, repoNameList []string) ([]interface{}, interface{}) {
	var milestoneList []interface{}
	for _, repo := range repoNameList {
		res := HTTPRequest(fmt.Sprintf(mileStoneURL, repo), g.Token)
		arr, err := res.Array()
		if err != nil || arr == nil {
			log.Fatal("something wrong when call Get and Array")
		}
		milestoneList = append(milestoneList, arr...)
	}

	toDoList := []interface{}{}
	head, _ := milestoneList[0].(map[string]interface{})
	for _, content := range milestoneList {
		newdi, _ := content.(map[string]interface{})
		if head["title"] == newdi["title"] {
			toDoList = append(toDoList, content)
		}
	}
	return toDoList, toDoList[0]
}

func (g *Github) GetMileStoneIssues(status string, milestone string, issueURL string, label string) {

}

func convertJSONtoInt(x json.Number) float64 {
	res, err := x.Float64()
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (g *Github) Rank(repoURL string, milestoneURL string) {
	milestoneList, LatestMileStone := g.GetLatestMilestone(milestoneURL, g.GetRepoList(repoURL))
	x, _ := LatestMileStone.(map[string]interface{})
	var openIssues float64
	var closedIssues float64
	var totalIssues float64
	for _, milestone := range milestoneList {
		m, _ := milestone.(map[string]interface{})
		openIssues += convertJSONtoInt(m["open_issues"].(json.Number))
		closedIssues += convertJSONtoInt(m["closed_issues"].(json.Number))
		totalIssues += convertJSONtoInt(m["open_issues"].(json.Number)) + convertJSONtoInt(m["closed_issues"].(json.Number))
	}
	closePercent := math.Ceil(closedIssues / totalIssues * 100)
	fmt.Println(x["title"])
	fmt.Println(openIssues, closedIssues, totalIssues, fmt.Sprintf("%s%%", strconv.FormatFloat(closePercent, 'f', 1, 64)))

}

// func GetLatestMilestone(repoNameList []string) {
// }

// def get_latest_milestone(repo_name_list):
//     milestone_list = []
//     for repo_name in repo_name_list:
//         milestone_list.extend(get_milestone_info(repo_name))

//     milestone_list.sort(key=lambda x: x.get("created_at"), reverse=True)

//     print("The latest milestone is: {}".format(milestone_list[0].get("title")))

//     todo_milestone = [milestone for milestone in milestone_list if
//                       milestone.get("title") == milestone_list[0].get("title")]

//     return milestone_list[0], todo_milestone
