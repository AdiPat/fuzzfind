package core

type Searcher struct {
	Id string
}

type SearcherOptions struct {
	Id string
}

func NewSearcher(options SearcherOptions) *Searcher {
	if options.Id == "" {
		options.Id = "searcher"
	}

	return &Searcher{
		Id: options.Id,
	}
}
