package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 获取data目录中的json数据
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲通道
	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)

	}

	go func() {
		waitGroup.Wait()

		close(results)
	}()

	Display(results)
}

func Register(feedType string, matcher Matcher) {
    if _, exists := matchers[feedType]; exists {
        log.Fatal(feedType, "Matcher already registered")
    }

    log.Println("Register", feedType, "matcher")
    matchers[feedType] = matcher

}
