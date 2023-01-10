package channels

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func Ponger(c chan<- string) {
	for i := 0; ;i++ {
		c <- "pong"
	}
}

func PingPonger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping pong"
	}
}

func Printer(c <-chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

type HomePageSize struct {
	URL string
	Size int
}

func BufferChannels() {
	urls := []string {
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func (url string)  {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL: url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("THe biggest homepage url is ",biggest.URL)
}