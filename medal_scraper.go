package main

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)


func main() {
	resp, err := http.Get("https://www.rio2016.com/en/medal-count-country?rank=total")

	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}
	
	// define a matcher
	/*
	<td class="col-3"><span class="country">Australia</span></td>
	<td class="col-4">2</td> 					GOLD
	<td class="col-5">0</td> 					SILVER
	<td class="col-6">1</td> 					BRONZE
	<td class="col-7"><strong>3</strong></td>  TOTAL
	*/
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.Tr {
			return scrape.Attr(n, "class") == "table-medal-countries__link-table"
		}
		return false
	}
	
	countries := scrape.FindAll(root, matcher)
	for i, entry := range countries {
		code, _   := scrape.Find(entry, scrape.ByClass("col-2"))
		name, _   := scrape.Find(entry, scrape.ByClass("col-3"))
		gold, _   := scrape.Find(entry, scrape.ByClass("col-4"))
		silver, _ := scrape.Find(entry, scrape.ByClass("col-5"))
		bronze, _ := scrape.Find(entry, scrape.ByClass("col-6"))

		fmt.Printf("%2d %s - %s %s %s (%s)\n", i, scrape.Text(code), scrape.Text(gold), scrape.Text(silver), scrape.Text(bronze), scrape.Text(name))
	}
}