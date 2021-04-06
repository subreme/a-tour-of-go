package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type mutexCache struct {
	mux   sync.Mutex
	sites map[string]bool
}

func (cache *mutexCache) visited(url string) bool {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	if cache.sites[url] {
		return true
	}
	cache.sites[url] = true
	return false
}

var cache = mutexCache{sites: make(map[string]bool)}

func crawlParallel(url string, depth int, fetcher Fetcher, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	if depth <= 0 {
		return
	}
	if cache.visited(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		waitGroup.Add(1)
		go crawlParallel(u, depth-1, fetcher, waitGroup)
	}
	return
}

func Crawl(url string, depth int, fetcher Fetcher) {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	go crawlParallel(url, depth, fetcher, waitGroup)
	waitGroup.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
	return "", nil, fmt.Errorf("not found: %s", url)
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
