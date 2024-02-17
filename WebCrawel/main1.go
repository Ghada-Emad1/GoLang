package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}
type URLS struct {
	mapping map[string]bool
	mux     sync.Mutex
}

func (u URLS) IsCrawled(url string) bool {
	fmt.Printf("\n👀Checking if %v is Crawled or Not \n", url)
	u.mux.Lock()
	defer u.mux.Unlock()
	if _, ok := u.mapping[url]; ok {
		fmt.Println("it hasn't\t")
		return ok
	}
	fmt.Println("it has\t")
	return false
}
func (u URLS) Crawled(url string) {
	u.mux.Lock()
	u.mapping[url] = true
	wg.Done()
	u.mux.Unlock()
}

var history = URLS{mapping: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 || history.IsCrawled(url) {
		return
	}
	res, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("✅ Found %s %q\n", url, res)
	for _, u := range urls {
		wg.Add(1)
		Crawl(u, depth-1, fetcher, wg)
	}
	return
}
func syncCrawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait()
}
func main() {
	syncCrawl("https://golang.org/", 4, fetcher)
	fmt.Println("Done")

}

type fakeFetcher map[string]*fakeResult
type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("⚠ Not Found :%s", url)

}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
