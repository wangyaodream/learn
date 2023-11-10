package search


import (
    "log"
)

// Result 保存搜索的结果
type Result struct {
    Filed   string
    Content string
}


type Matcher interface {

    Search(feed *Feed, searchTerm string) ([]*Result, error)
}

