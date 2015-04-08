package main

import (
	"fmt"
	//"log"
)

type Tweet struct {
	User     string
	Message  string
	Retweets int
}

//
func mapper(object string) (interface{}, error) {

	fmt.Println("Testing mapper")
	tweet := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	return tweet, nil
}

func processor(object interface{}) error {

	fmt.Println(object.(Tweet).User)
	fmt.Println(object.(Tweet).Message)
	fmt.Println(object.(Tweet).Retweets)
	return nil
}

func main() {
	content, _ := mapper("8453753745375483 INFO rereterte  erter ")


	///Queue


	processor(content)
}
