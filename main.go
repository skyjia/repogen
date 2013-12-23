/*
 Author: Sky Jia
 GitHub: https://github.com/skyjia
 Web: http://www.skyjia.com
*/

package main

import (
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"os"
	"time"
)

var (
	username = flag.String("u", "", "GitHub username. (Required)")
)

func init() {
	flag.Parse()

	if *username == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Init var
	langRepoMap = map[string][]github.Repository{}
	languageList = []string{}
}

const PER_PAGE int = 100

var langRepoMap map[string][]github.Repository
var languageList []string

func main() {
	fetchGitHubData()
	printDocument()
}

func fetchGitHubData() {
	client := github.NewClient(nil)
	opt := &github.ActivityListStarredOptions{}
	opt.ListOptions.PerPage = PER_PAGE

	page_idx := 1
	for {
		opt.ListOptions.Page = page_idx

		reps, _, _ := client.Activity.ListStarred(*username, opt)
		addListToMap(reps)

		if len(reps) != PER_PAGE {
			break
		}

		page_idx++
	}
}

func addListToMap(list []github.Repository) {
	if len(list) == 0 {
		return
	}

	for _, r := range list {
		lang := "Unknown"
		if r.Language != nil {
			lang = *r.Language
		}

		langList, ok := langRepoMap[lang]
		if !ok {
			langList = []github.Repository{}
			languageList = append(languageList, lang)
		}
		// FIXME: Memory problem here
		langList = append(langList, r)
		langRepoMap[lang] = langList
	}
}

func printDocument() {
	printHeader()
	printLanguageList()
	printRepositoriesByLanguage()
	printFooter()
}

func printHeader() {
	fmt.Println("# Starred Repositories")
	fmt.Println("===============")
	fmt.Println()
	fmt.Printf("__[%s](https://github.com/%s)__ on GitHub.\r\n", *username, *username)
	fmt.Println()
}

func printLanguageList() {
	fmt.Println("## Languages")
	fmt.Println()

	for _, lang := range languageList {
		fmt.Printf("- [%s](#%s)\r", lang, lang)
	}

	fmt.Println()
}

func printRepositoriesByLanguage() {
	fmt.Println("# Index")
	fmt.Println()

	for _, lang := range languageList {
		list, _ := langRepoMap[lang]
		fmt.Printf("## [%s](id:%s) (%d)\r\n", lang, lang, len(list))
		fmt.Println()

		printRepositoryList(list)
	}
}

const layout = "Jan 2, 2006 3:04PM (MST)"

func printRepositoryList(list []github.Repository) {
	for _, r := range list {
		fmt.Printf("### %s/%s\r\n", *r.Owner.Login, *r.Name)
		fmt.Println()
		fmt.Printf("%s\r\n", *r.Description)
		fmt.Println()
		fmt.Printf("- URL: <%s>\r\n", *r.HTMLURL)
		fmt.Println()
		if r.Homepage != nil && *r.Homepage != "" {
			fmt.Printf("- Site: <%s>\r\n", *r.Homepage)
			fmt.Println()
		}

		fmt.Printf("- Pushed at: _%s_\r\n", r.PushedAt.Format(layout))
		fmt.Println()
	}
}

func printFooter() {
	fmt.Println("---")
	fmt.Println("Generated at:", time.Now().UTC().Format(layout))
	fmt.Println()
	fmt.Println("_Get generator on [GitHub](https://github.com/skyjia/github-repo-gen)_")
}
