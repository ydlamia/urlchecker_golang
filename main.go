package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	//1
	// var results map[string]string
	// results["hello"] = "Hello"	//panic -> results 초기화를 해주지 않아서
	//solve -> 2, 3번과 같이 초기화를 해준뒤 값을넣어야 함

	//2
	// var results = map[string]string{}
	// results["hello"] = "Hello"

	//3
	// var results = make(map[string]string)	//make함수는 안에 선언된 map의 타입에 맞춰 map을 생성하고 빈 상태로 초기화해줌
	// results["hello"] = "Hello"

	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.naver.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co/",
		"https://www.abc.com",
		"https://workingscorpion.com",
	}

	for _, url := range urls {
		// fmt.Print("Checking : ", url, " ")
		go hitURL(url, c)
		// if err == errRequestFailed {
		// 	results[url] = "DOWN"
		// } else {
		// 	results[url] = "OK"
		// }
		// fmt.Println(results[url])
	}

	fmt.Println("===============")

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "Failed"
	}
	c <- result{
		url:    url,
		status: status,
	}

}

//send Only
// func hitURL(url string, c chan<- result) {
// 	fmt.Println(<-c)
// }
