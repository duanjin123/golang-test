package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main()  {
	// 定义一个Movie结构体的slice
	movies := []Movie{
		{Title:  "战狼1", Year:   2015, Color:  false, Actors: nil,},
		{Title:  "血战长津湖", Year:   2017, Color:  false, Actors: nil,},
		{Title:  "金刚川", Year:   2018, Color:  false, Actors: nil,},
		{Title:  "狙击手", Year:   2020, Color:  false, Actors: nil,},
	}
	fmt.Printf("movies的值为：%v\n", movies)
	// 将其转化为json
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("转化json出错: %v\n", err)
	}
	fmt.Printf("转化后json的值为：%s\n", data)
}
