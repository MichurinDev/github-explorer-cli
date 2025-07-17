package main

import (
	"flag"
	"fmt"
	"github-explorer-cli/internal/app"
	"strings"
)

func main() {
	lang := flag.String("lang", "go", "Язык программирования для фильтрации")
	period := flag.String("period", "week", "Период тренда (day, week, month)")
	top := flag.Int("top", 10, "Количество топовых репозиториев для вывода")

	flag.Parse()

	repos, err := app.FetchTrendingRepos(*lang, *period, *top)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("\033[32m ------ Top-\033[36m%d\033[32m repo with \033[36m%s\033[32m by \033[36m%s\033[32m -------\033[0m\n", *top, *lang, *period)
	if len(repos) == 0 {
		fmt.Println(
			"Нет репозиториев для выбранного языка и периода.",
			"Попробуйте изменить параметры.",
		)
	}

	for i, repo := range repos {
		fmt.Printf("\033[36m%2d. %s ⭐ %d  Issues: %d\n   \033[35m%s\033[0m\n   URL: %s\n\n",
			i+1, repo.FullName, repo.Stars, repo.OpenIssues,
			strings.Trim(repo.Description, "\n"), repo.RepoURL)
	}

	fmt.Println("\033[32m------------- END -------------------\033[0m")
}
