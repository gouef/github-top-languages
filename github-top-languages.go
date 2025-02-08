package main

import (
	"fmt"
	"github.com/gouef/githubtoplanguages/requests"
	"github.com/gouef/utils"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sort"
	"strconv"
)

/*
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"repositories": &graphql.Field{
				Type: graphql.NewList(repositoryType),
			},
		},
	},
)*/
/*
var repositoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Repository",
		Fields: graphql.Fields{
			"name":      &graphql.Field{Type: graphql.String},
			"fullName":  &graphql.Field{Type: graphql.String},
			"languages": &graphql.Field{Type: graphql.NewList(languageType)},
		},
	},
)*/
/*
var languageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Language",
		Fields: graphql.Fields{
			"name": &graphql.Field{Type: graphql.String},
			"size": &graphql.Field{Type: graphql.Int},
		},
	},
)*/
/*
type GraphQLResponse struct {
	Data struct {
		Viewer struct {
			Repositories struct {
				Nodes    []organizations.Repository `json:"nodes"`
				PageInfo toplanguages.PageInfo      `json:"pageInfo"`
			} `json:"repositories"`
			Organizations struct {
				Nodes []struct {
					Login        string `json:"login"`
					Repositories struct {
						Nodes    []organizations.Repository `json:"nodes"`
						PageInfo toplanguages.PageInfo      `json:"pageInfo"`
					} `json:"repositories"`
				} `json:"nodes"`
				PageInfo toplanguages.PageInfo `json:"pageInfo"`
			} `json:"organizations"`
		} `json:"viewer"`
	} `json:"data"`
}
*/
/*
func fetchRepositories(token string, query string) (*GraphQLResponse, error) {
	payload := struct {
		Query string `json:"query"`
	}{
		Query: query,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", githubGraphQLAPI, bytes.NewBuffer(payloadJSON)) // Use bytes.NewBuffer
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json") // Important: Set Content-Type header

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body) // Read the error body
		return nil, fmt.Errorf("GitHub API returned non-200 status: %d, Body: %s", resp.StatusCode, string(bodyBytes))
	}

	var result GraphQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}*/
/*
func fetchAllRepositories(token string) ([]Repository, map[string]int, error) {
	var repositories []Repository
	languageStats := make(map[string]int)

	query := `{
		viewer {
			repositories(first: 100, isFork: false) {
				nodes {
					name
					nameWithOwner
					languages(first: 10) {
						edges {
							node { name }
							size
						}
					}
				}
				pageInfo { hasNextPage endCursor }
			}
		}
	}`

	response, err := fetchRepositories(token, query)
	if err != nil {
		return nil, nil, err
	}

	repositories = append(repositories, response.Data.Viewer.Repositories.Nodes...)

	for _, repo := range repositories {
		for _, edge := range repo.Languages.Edges {
			languageStats[edge.Node.Name] += edge.Size
		}
	}

	return repositories, languageStats, nil
}*/

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}
	user := os.Getenv("GITHUB_USERNAME")
	if user == "" {
		log.Fatal("GITHUB_USERNAME is not set")
	}

	ignoredOrganizationsEnv := os.Getenv("GITHUB_IGNORE_ORGANIZATIONS")
	ignoredRepositoriesEnv := os.Getenv("GITHUB_IGNORE_REPOS")

	ignoredOrganizations := utils.Explode(",", ignoredOrganizationsEnv)
	ignoredRepositories := utils.Explode(",", ignoredRepositoriesEnv)
	ignored := ignoredOrganizations
	ignored = append(ignored, ignoredRepositories...)

	limitEnv := os.Getenv("GITHUB_TOP_LIMIT")
	limit := 10

	if limitEnv != "" {
		limit, _ = strconv.Atoi(limitEnv)
	}

	result, err := requests.FetchOrganizations(user, token, ignored...)

	if err != nil {
		log.Fatalf("Failed to fetch organizations: %v", err)
	}

	var repositories []string
	languages := make(map[string]int)

	for _, repoList := range result.Repositories {
		repositories = append(repositories, repoList.Name)

		for _, lang := range repoList.Languages {
			languages[lang.Name] += lang.Size
		}
	}

	result, err = requests.FetchUser(user, token, ignored...)

	if err != nil {
		log.Fatalf("Failed to fetch user: %v", err)
	}

	for _, repoList := range result.Repositories {
		if utils.InArray(repoList.Name, repositories) {
			continue
		}

		repositories = append(repositories, repoList.Name)

		for _, lang := range repoList.Languages {
			languages[lang.Name] += lang.Size
		}
	}

	languages, languagesPercentage := sortLanguages(languages, limit)

	//languagesPercentage := make(map[string]float64)

	for percentage, lang := range languagesPercentage {
		fmt.Printf("%s: %.2f\n", lang, percentage)
	}

	/*
		for lang, size := range languages {
			languagesPercentage[lang] = (float64(size) / float64(languagesSize)) * 100
			fmt.Printf("%s: %.2f\n", lang, languagesPercentage[lang])
		}*/

}

func sortLanguages(languages map[string]int, limit int) (map[string]int, map[float64]string) {
	type kv struct {
		Key   string
		Value int
	}

	var sorted []kv
	for k, v := range languages {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	if limit > len(sorted) {
		limit = len(sorted)
	}

	sortedMap := make(map[string]int, limit)

	languagesPercentage := make(map[string]float64)
	var languagesSize int
	for i := 0; i < limit; i++ {
		sortedMap[sorted[i].Key] = sorted[i].Value
		languagesSize += sorted[i].Value
	}

	for key, size := range sortedMap {
		languagesPercentage[key] = (float64(size) / float64(languagesSize)) * 100
	}

	type kvp struct {
		Key   string
		Value float64
	}

	var sortedp []kvp
	for k, v := range languagesPercentage {
		sortedp = append(sortedp, kvp{k, v})
	}

	sort.Slice(sortedp, func(i, j int) bool {
		return sortedp[i].Value > sortedp[j].Value
	})

	sortedpMap := make(map[float64]string, limit)

	for i := 0; i < limit; i++ {
		sortedpMap[sortedp[i].Value] = sortedp[i].Key
	}

	return sortedMap, sortedpMap
}
