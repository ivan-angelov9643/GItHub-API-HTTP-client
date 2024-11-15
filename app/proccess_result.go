package app

import (
	"fmt"
	"log"
)

func GenerateReport(usernames []string) {
	for _, username := range usernames {
		user, err := fetchUser(username)
		if err != nil {
			log.Printf("Failed to fetch user %s: %v", username, err)
			continue
		}

		repos, err := fetchRepositories(username)
		if err != nil {
			log.Printf("Failed to fetch repositories for user %s: %v", username, err)
			continue
		}

		totalForks := 0
		activityByYear := make(map[int]int)
		languageCounts := make(map[string]int)
		totalLanguageUsage := 0

		for _, repo := range repos {
			totalForks += repo.Forks

			activityByYear[repo.CreatedAt.Year()]++
			activityByYear[repo.UpdatedAt.Year()]++

			languages, err := fetchLanguages(username, repo.Name)
			if err != nil {
				log.Printf("Failed to fetch languages for repo %s: %v", repo.Name, err)
				continue
			}

			for lang, count := range languages {
				languageCounts[lang] += count
				totalLanguageUsage += count
			}
		}

		var activityStr string
		for year, count := range activityByYear {
			activityStr += fmt.Sprintf("%d:%d ", year, count)
		}

		var langStr string
		for lang, count := range languageCounts {
			percentage := (float64(count) / float64(totalLanguageUsage)) * 100
			langStr += fmt.Sprintf("%s: %.2f%% ", lang, percentage)
		}

		fmt.Printf("Username: %s\n", user.Login)
		fmt.Printf("Followers: %d\n", user.Followers)
		fmt.Printf("Number of Repos: %d\n", len(repos))
		fmt.Printf("Languages: %s\n", langStr)
		fmt.Printf("Total Forks: %d\n", totalForks)
		fmt.Printf("Activity by year (<year>:<number of actions>): %s\n", activityStr)
		fmt.Println("---------------------------------------------------")
	}
}
