package main

import (
	"github.com/goadesign/goa"
	"github.com/ivan3bx/medal-service/app"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"strconv"
)

// MedalsController implements the medals resource.
type MedalsController struct {
	*goa.Controller
}

// NewMedalsController creates a medals controller.
func NewMedalsController(service *goa.Service) *MedalsController {
	return &MedalsController{Controller: service.NewController("MedalsController")}
}

// Index runs the index action.
func (c *MedalsController) Index(ctx *app.IndexMedalsContext) error {
	// MedalsController_Index: start_implement

	resp, err := http.Get("https://www.rio2016.com/en/medal-count-country?rank=total")

	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}

	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.Tr {
			return scrape.Attr(n, "class") == "table-medal-countries__link-table"
		}
		return false
	}

	res := app.MedalsCollection{}

	countries := scrape.FindAll(root, matcher)
	for _, entry := range countries {
		code, _ := scrape.Find(entry, scrape.ByClass("col-2"))
		name, _ := scrape.Find(entry, scrape.ByClass("col-3"))
		gold, _ := scrape.Find(entry, scrape.ByClass("col-4"))
		silver, _ := scrape.Find(entry, scrape.ByClass("col-5"))
		bronze, _ := scrape.Find(entry, scrape.ByClass("col-6"))

		res = append(res, &app.Medals{
			CountryID:   scrape.Text(code),
			CountryName: toString(name),
			Gold:        toInt(gold),
			Silver:      toInt(silver),
			Bronze:      toInt(bronze),
		})
	}

	return ctx.OK(res)
}

func toString(node *html.Node) *string {
	val := scrape.Text(node)
	return &val
}

func toInt(node *html.Node) int {
	txt := scrape.Text(node)
	i, _ := strconv.Atoi(txt)
	return i
}
