package search

// defaultMatcher default
type defaultMatcher struct {}


func init() {
    var matcher defaultMatcher
    Register("default", matcher)
}


func (m defaultMather) Search(feed *Feed, searchTerm string) ([]*Result, error) {
    return nil, nil
}
