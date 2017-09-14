package score

import (
	"strings"

	"github.com/g-hyoga/CP5I/server/model"
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

func Score(method []model.Method) (int, error) {
	scoreResult := 0

	for _, m := range method {
		for _, difficultyWord := range difficultyWords {
			if strings.Contains(m.Content, difficultyWord) {
				scoreResult += dictionary[difficultyWord]
			}
		}
	}

	return scoreResult, nil
}
