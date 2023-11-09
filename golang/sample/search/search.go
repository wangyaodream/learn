package search


import (
    "log"
    "sync"
)


// 注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)


func Run(searchTerm string) {
    feeds, err := RetrieveFeeds()
    if err != nil {
        log.Fatal(err)
    }

    // 创建无缓冲的通道,接收匹配后的结果
    results := make(chan *Result)

    var waitGroup sync.WaitGroup


    // 为每个数据源启动一个goroutine
    for _, feed := range feeds {
        matcher, exists := matchers[feed.Type]
        if ! exists {
            matcher = matchers["defaults"]
        }

        // 启动一个goroutine来执行搜索
        go func(matcher Matcher, feed *Feed) {
            Match(matcher, feed, searchTerm, results)
            waitGroup.Done()
        } (matcher, feed)
    }


    // 启动一个goroutine来监控是否所有的工作都做完了
    go func() {
        waitGroup.Wait()

        close(results)
    }()

    // 在最后一个结果显示完后返回
    Display(results)

}

