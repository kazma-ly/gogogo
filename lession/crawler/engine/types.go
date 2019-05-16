package engine

type ParserFunc func(contents []byte, url string) ParseResult

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Request struct {
	URL        string
	ParserFunc ParserFunc
}

type Item struct {
	Url     string
	Id      string
	Type    string // search type
	Payload interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
