package search

import (
    "encoding/json"
    "os"
)


const dataFile = "data/data.json"

// Feed包含需要处理的数据
type Feed struct {
    Name string `json:"site"`
    URI string `json:"link"`
    Type string `json:"type"`
}


// 读取反序列化源数据
func RetrieveFeeds() ([]*Feed, error) {
    // 打开文件
    file, err := os.Open(dataFile)
    if err != nil {
        return nil, err
    }

    defer file.close()

    // 将文件解码到一个切片里
    var feeds []*Feed
    err = json.NewDecoder(file).Decode(&feeds)

    return feeds, err
}
