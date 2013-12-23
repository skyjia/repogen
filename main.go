package main

import (
	"fmt"
	"github.com/google/go-github/github"
)

const USERNAME string = "skyjia"
const PER_PAGE int = 100

var languageMap map[string][]github.Repository
var languageList []string

func main() {

	languageMap = map[string][]github.Repository{}
	languageList = []string{}

	client := github.NewClient(nil)
	opt := &github.ActivityListStarredOptions{}
	opt.ListOptions.PerPage = PER_PAGE
	page_idx := 1

	for {
		opt.ListOptions.Page = page_idx

		reps, _, _ := client.Activity.ListStarred(USERNAME, opt)
		if len(reps) == 0 {
			break
		}

		addListToMap(reps)

		page_idx++
	}

	printDocument()
}

func addListToMap(list []github.Repository) {
	for _, r := range list {
		lang := "Unknown"
		if r.Language != nil {
			lang = *r.Language
		}

		langList, ok := languageMap[lang]
		if !ok {
			langList = []github.Repository{}
			languageList = append(languageList, lang)
		}
		// FIXME: Memory problem here
		langList = append(langList, r)
		languageMap[lang] = langList
	}
}

func printDocument() {
	fmt.Println("# Starred Repositories")
	fmt.Println()

	printLanguageList()
	printRepositoriesByLanguage()
}

func printLanguageList() {
	fmt.Println("## Languages")
	fmt.Println()

	for _, lang := range languageList {
		fmt.Printf("- [%s](#%s)\r\n", lang, lang)
		fmt.Println()
	}
}

func printRepositoriesByLanguage() {
	for _, lang := range languageList {
		list, _ := languageMap[lang]
		fmt.Printf("## [%s](id:%s) (%d)\r\n", lang, lang, len(list))
		fmt.Println()

		printRepositoryList(list)
	}
}

const layout = "Jan 2, 2006 at 3:04PM (MST)"

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

		fmt.Printf("_%s_\r\n", r.PushedAt.Format(layout))
		fmt.Println()
	}
}
