package score

import (
	"strings"

	"github.com/g-hyoga/CP5I/server/model"
	"github.com/k0kubun/pp"
)

// mockとしてはこれで良い
var difficultyWords = []string{
	"揚げる",
	"フライパン",
	"炒める",
}

var dictionary = map[string]int{
	"揚げる":   20,
	"フライパン": 10,
	"炒める":   10,
}

func Contains(array []string, target string) bool {
	for _, e := range array {
		if e == target {
			pp.Println(e, target)
			return true
		}
	}
	return false
}

func Score(method []model.Method) (int, error) {
	scoreResult := 0
	foundWords := []string{}

	for _, m := range method {
		for _, difficultyWord := range difficultyWords {
			if strings.Contains(m.Content, difficultyWord) && !Contains(foundWords, difficultyWord) {
				foundWords = append(foundWords, difficultyWord)
				scoreResult += dictionary[difficultyWord]
			}
		}
	}

	return scoreResult, nil
}
