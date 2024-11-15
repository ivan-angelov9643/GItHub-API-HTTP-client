package app

import (
	"fmt"
	"log"
)

import (
	"sort"
)

func GenerateReport(usernames []string) {
	fmt.Printf("%-20s %-15s %-10s %-20s %-15s %-20s\n", "Username", "Followers", "Repos", "Languages", "Total Forks", "Activity")
	fmt.Printf("%-20s %-15s %-10s %-20s %-15s %-20s\n", "--------", "---------", "-----", "---------", "-----------", "--------")

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

		languageCounts, totalForks, activityByYear := collectRepoData(repos, username)

		langLines := calculateLanguagePercentages(languageCounts)
		activityLines := calculateActivityByYear(activityByYear)

		printReport(user, repos, totalForks, langLines, activityLines)
	}
}

func collectRepoData(repos []Repository, username string) (map[string]int, int, map[int]int) {
	totalForks := 0
	activityByYear := make(map[int]int)
	languageCounts := make(map[string]int)

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
		}
	}

	return languageCounts, totalForks, activityByYear
}

type LangLine struct {
	Language   string
	Percentage float64
}

func calculateLanguagePercentages(languageCounts map[string]int) []LangLine {
	var langLines []LangLine

	totalLanguageUsage := 0
	for _, count := range languageCounts {
		totalLanguageUsage += count
	}

	for lang, count := range languageCounts {
		percentage := (float64(count) / float64(totalLanguageUsage)) * 100
		langLines = append(langLines, LangLine{lang, percentage})
	}

	sort.Slice(langLines, func(i, j int) bool {
		return langLines[i].Percentage > langLines[j].Percentage
	})

	return langLines
}

type ActivityLine struct {
	Year  int
	Count int
}

func calculateActivityByYear(activityByYear map[int]int) []ActivityLine {
	var activityLines []ActivityLine

	for year, count := range activityByYear {
		activityLines = append(activityLines, ActivityLine{year, count})
	}

	sort.Slice(activityLines, func(i, j int) bool {
		return activityLines[i].Year > activityLines[j].Year
	})

	return activityLines
}

func printReport(user User, repos []Repository, totalForks int, langLines []LangLine, activityLines []ActivityLine) {
	maxLines := len(langLines)
	if len(activityLines) > maxLines {
		maxLines = len(activityLines)
	}

	for i := 0; i < maxLines; i++ {
		usernameField := ""
		followersField := ""
		reposField := ""
		totalForksField := ""
		if i == 0 {
			usernameField = user.Login
			followersField = fmt.Sprintf("%d", user.Followers)
			reposField = fmt.Sprintf("%d", len(repos))
			totalForksField = fmt.Sprintf("%d", totalForks)
		}

		langField := ""
		if i < len(langLines) {
			langField = fmt.Sprintf("%s: %.2f%%", langLines[i].Language, langLines[i].Percentage)
		}

		activityField := ""
		if i < len(activityLines) {
			activityField = fmt.Sprintf("%d: %d", activityLines[i].Year, activityLines[i].Count)
		}

		fmt.Printf("%-20s %-15s %-10s %-20s %-15s %-20s\n",
			usernameField, followersField, reposField, langField, totalForksField, activityField)
	}

	fmt.Println("---------------------------------------------------")
}
