package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aosasona/interim"
)

func main() {
	ss := flag.Int("sample-size", 1000, "number of items insertion, fetching and removal to do")
	cs := flag.Int("cache-size", 16, "LRU Cache size")

	flag.Parse()

	sampleSize := *ss
	cacheSize := *cs

	db := interim.New(interim.Config{CacheSize: cacheSize})

	values := []string{
		"Quis aliquet odio various ut phasellus sit amet aliquam consectetur adipiscing elit ut aliquam sapien.",
		"Maecenas consectetur diam sed diam viverra dignissim ut at lorem quisque dignissim sagittis aenean euismod elementum.",
		"Ut enim ad minim veniam quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
		"Etiam auctor nibh ut hendrerit consectetur nisi lectus eget posuere mi hendrerit.",
		"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium totam rem aperiam.",
		"Name fermentum augue vel turpis convallis ut interdum diam gravida nec ut enim.",
		"Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae donec velit neque auctor sit amet.",
		"Quisque cursus ante at mauris commodo euismod name semper justo quis risus suscipit scelerisque.",
		"Nunc congue nisi sed justo sollicitudin euismod sed quis lectus sed mi.",
		"Fusce et quam semper dignissim eget ac magna aliquet consectetur adipiscing elit.",
	}

	fmt.Print("==============================================================")
	fmt.Printf("\nSample Size: %d\nCache Size: %d\n", sampleSize, cacheSize)

	// Write
	start := time.Now()
	for i := 0; i < sampleSize; i++ {
		db.Set(fmt.Sprintf("%d", i), values[(i%10)])
	}
	end := time.Now()

	duration := end.Sub(start)
	fmt.Printf("Writing %v items took %v\n", sampleSize, duration)

	// Read
	start = time.Now()
	count := 0
	for i := 0; i < sampleSize; i++ {
		var result string
		err := db.Get(fmt.Sprintf("%d", i), &result)
		if err == nil {
			count++
		}
	}
	duration = time.Since(start)

	fmt.Printf("(forwards) Reading %d/%d items took %v\n", count, sampleSize, duration)

	// Read - from the back
	start = time.Now()
	count = 0
	for i := (sampleSize - 1); i >= 0; i-- {
		var result string
		err := db.Get(fmt.Sprintf("%d", i), &result)
		if err == nil {
			count++
		}
	}
	duration = time.Since(start)

	fmt.Printf("(backwards) Reading %d/%d items took %v\n", count, sampleSize, duration)

	// Delete
	start = time.Now()
	count = 0
	for i := 0; i < sampleSize; i++ {
		err := db.Delete(fmt.Sprintf("%d", i))
		if err == nil {
			count++
		}
	}
	duration = time.Since(start)

	fmt.Printf("Removing %d/%d items took %v\n", count, sampleSize, duration)

	fmt.Print("==============================================================")
}
