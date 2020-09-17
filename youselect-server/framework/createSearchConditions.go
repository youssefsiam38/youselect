package framework

import (
	"fmt"
	"strings"
)

func CreateSearchConditions(searchQuery *string, exact bool) *string {

	words := strings.Split(*searchQuery, " ")
	searchConditions := ""
	if exact {
		for i := 0; i < len(words); i++ {
			if words[i] == "" {
				continue
			}
			searchConditions = fmt.Sprintf("%s and products.title like '%%%s%%'", searchConditions, words[i])
		}
	} else {
		for i := 0; i < len(words); i++ {
			if words[i] == "" {
				continue
			}
			searchConditions = fmt.Sprintf("%s or products.title like '%%%s%%'", searchConditions, words[i])
		}
	}
	
	if searchConditions == "" {
		falseCond := "and 1=0"
		return &falseCond
	}

	searchConditions = fmt.Sprintf("and (%s)", searchConditions)

	if exact {
		searchConditions = strings.Replace(searchConditions, "( and", "(", 1)
	} else {
		searchConditions = strings.Replace(searchConditions, "( or", "(", 1)
	}

	return &searchConditions
}

// " and title like '%new%' and title like '%laptop%'"
