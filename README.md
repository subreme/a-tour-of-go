# A Tour Of Go Exercises
As I progressed through [A Tour of Go](https://tour.golang.org/) while learning Golang, I was slightly annoyed by how my edits would be reset after a while, so I decided to save my answers to the exercises on this repo. Some answers are quite similar to the [official ones](https://github.com/golang/tour/tree/master/solutions), in part as a coincidence and in part because i occasionally had a look when I was stuck. Nevertheless, I recommend trying them out for yourselves if you're also interested in learning the language, and hope you all have a great day!

**Note**: Some of my solutions aren't that great and lack annotations, as I removed most comments. Anyone is more than welcome to submit a pull request to improve the code or add explanations to how it works, although I don't expect any as the main reason I'm posting this is that I find it pretty sad not to have any public repositories and I don't have any good projects to share [yet].

## Exercises
#### Loops and Functions
The goal for this [exercise](https://tour.golang.org/flowcontrol/8) was to create a function using [Newton's method](https://en.wikipedia.org/wiki/Newton%27s_method) to approximate the square root of 2.
```Go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, y := x, 0.0
	for math.Abs(y-z) > 1e-6 {
		y, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println("Sqrt()'s result: ", Sqrt(2))
	fmt.Println("math.Sqrt()'s result: ", math.Sqrt(2))
}
```

#### Slices
In this [exercise](https://tour.golang.org/moretypes/18), the `Pic()` function is used to generate `uint8/byte` values in a slice within a slice (used as x and y coordinates) based on the indices of the two slices.
```Go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for x := range pic {
		pic[x] = make([]uint8, dx)
		for y := range pic[x] {
			// pic[x][y] = uint8((x+y)/2)
			// pic[x][y] = uint8(x*y)
			pic[x][y] = uint8(x ^ y)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
```

#### Maps
This [task](https://tour.golang.org/moretypes/23) involved creating a function that returns a map containing all words within the string `s` as keys to a value corresponding to their frequency.
```Go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```

#### Fibonacci closure
As the name suggests, this [exercise](https://tour.golang.org/moretypes/26) involved creating a function that returned consecutive elements of the [Fibonacci sequence](https://en.wikipedia.org/wiki/Fibonacci_number).
```Go
package main

import "fmt"

func fibonacci() func() int {
	x, y := 1, 0
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

#### Stringers
Given an array of 4 bytes, the goal of this [exercise](https://tour.golang.org/methods/18) was to join the bytes into a single string, separating them with a dot in order to format them like an IP Address.
```Go
package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	// Simple solution:
	// return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2],ip[3])

	// More interesting (needlessly "complicated") solution:
	s := make([]string, len(ip))
	for i, val := range ip {
		s[i] = fmt.Sprint(int(val))
	}
	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```
 
#### Errors
Adding on to the previous `Sqrt()` function, this [exercise](https://tour.golang.org/methods/20) involved modifying the original code to return an error message instead of a [complex number](https://en.wikipedia.org/wiki/Complex_number) when given a negative integer as a parameter.
 ```Go
 package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintln("cannot Sqrt negative number:", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z, y := x, 0.0
	for math.Abs(y-z) > 1e-6 {
		y, z = z, z-(z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

#### Readers
This [task](https://tour.golang.org/methods/22) was rather simple: to output an infinite stream of the ASCII character `'A'`.
```Go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 'A'
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
```

#### rot13 Reader
The goal of this [exercise](https://tour.golang.org/methods/23) was to create a function that "deciphers" a message encoded using the [ROT13](https://en.wikipedia.org/wiki/ROT13) substitution cypher.
```Go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, e error) {
	n, e = r.r.Read(b)
	if e == nil {
		// This is messy but it works...
		for i := 0; i < n; i++ {
			if 'A' <= b[i] && b[i] <= 'Z' || 'a' <= b[i] && b[i] <= 'z' {
				b[i] += 13
				if b[i] > 'z' {
					b[i] -= 26
				}
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

#### Images
This [exercise](https://tour.golang.org/methods/25) also involved the improvement of a previous function, adapting the code from the "Slices" lesson to have it's own custom interface.
```Go
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
	// v := uint8((x+y)/2)
	// v := uint8(x*y)
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
```

#### Equivalent Binary Trees
Explained both [here](https://tour.golang.org/concurrency/7) and [here](https://tour.golang.org/concurrency/8), this lesson's task was to create a function that compares [binary trees](https://en.wikipedia.org/wiki/Binary_tree) and determines whether or not they are equivalent to each other.
```Go
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
    if !ok1 || !ok2 {
			break
		}
		if v1 != v2 || ok1 != ok2 {
			return false
		}
	}
	return true
}

func main() {
	test1, test2 := Same(tree.New(1), tree.New(1)), Same(tree.New(1), tree.New(2))
	fmt.Println("Same(tree.New(1), tree.New(1)):", test1)
	fmt.Println("Same(tree.New(1), tree.New(2)):", test2)
	if test1 && !test2 {
		fmt.Println("Success!")
	} else {
		fmt.Println("Something's not right...")
	}
}
```

#### Web Crawler
This final [exercise](https://tour.golang.org/concurrency/10) involved modifying a given [web crawler](https://en.wikipedia.org/wiki/Web_crawler) to run in parallel by using Go Routines while caching URLs in order to avoid visiting a page twice.
```Go
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
```
