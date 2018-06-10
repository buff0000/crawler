package engine

import (
	"log"
	"com.buff/Crawler/crawler/crawler/fetcher"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request){
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got items -------[%v%s", item, "]")
		}
	}
}

func worker(r Request) (ParseResult, error){
	log.Printf("Fetch Url: %v",r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}