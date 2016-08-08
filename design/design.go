package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() {
	Title("Rio Medal Count")
	Description("API for olympic medals")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("medals", func() {
	BasePath("/medals")
	Action("index", func() {
		Description("Get list of medals")
		Routing(GET("/"))
		Params(func() {
			Param("sort", String, "Sort by", func() {
				Enum("country", "total", "gold", "silver", "bronze")
			})
		})
		Response(OK, CollectionOf(CountryMedals))
		Response(NotFound)
	})
})

var CountryMedals = MediaType("application/vnd.medals+json", func() {
	Description("A list of country medals")
	Attributes(func() {
		Attribute("countryId", String, "Short identifier for country")
		Attribute("countryName", String, "Descriptive name for country")
		Attribute("gold", Integer, "Number of gold medals")
		Attribute("silver", Integer, "Number of silver medals")
		Attribute("bronze", Integer, "Number of bronze medals")
		Required("countryId", "gold", "silver", "bronze")
	})

	View("default", func() {
		Attribute("countryId")
		Attribute("countryName")
		Attribute("gold")
		Attribute("silver")
		Attribute("bronze")
	})
})
