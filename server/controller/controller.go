package controller

import "github.com/gin-gonic/gin"

type Instructor struct {
	ID      int    `josn:"id"`
	Name    string `json:"name"`
	ImageID int    `json:"image_id"`
}

type Dish struct {
	ID      int    `json:"id"`
	Name    string `json:name`
	ImageID int    `json"image_id"`
}

type Data struct {
	ID             int        `json:"id"`
	InstructorData Instructor `json:"instructor"`
	DishData       Dish       `json:"dish"`
}

func GetInstructMatch(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": []Data{
			Data{
				ID: 1,
				DishData: Dish{
					ID:      1,
					Name:    "超絶品＊簡単＊煮込みハンバーグ♬",
					ImageID: 1,
				},
				InstructorData: Instructor{
					ID:      2,
					Name:    "はじめ",
					ImageID: 2,
				},
			},
			Data{
				ID: 2,
				DishData: Dish{
					ID:      2,
					Name:    "半熟卵とブロッコリーのアボカドサラダ",
					ImageID: 2,
				},
				InstructorData: Instructor{
					ID:      3,
					Name:    "アヤコ",
					ImageID: 3,
				},
			},
		},
	})
}
