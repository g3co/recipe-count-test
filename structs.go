package main

const (
	fieldsQuantity  = 3       // quantity of given fields postcode, delivery, recipe
	recipeMapSize   = 2010    // allocate memory for 2K+ recipe names
	postCodeMapSize = 1000100 // allocate memory for 1M+ post codes
	bufferSize      = 1024
)

type (
	buffType     [fieldsQuantity][]byte
	recipeName   [110]byte // Max recipe name size 100 chars
	postCodeName [11]byte  // Max postcode name 10 chars
	recipeMap    map[recipeName]int64
	postCodeMap  map[postCodeName]int64
	weekDay      [10]byte
	workTime     [4]byte

	Config struct {
		PostCodeCounter struct {
			PostCode string `json:"postcode"`
			From     int64  `json:"from"`
			To       int64  `json:"to"`
		} `json:"postcode_counter"`
		SearchByName []string `json:"search_by_name"`
	}

	DeliveryObj struct {
		WeekDay weekDay
		Start   int64
		End     int64
	}

	RecipeObj struct {
		PostCode postCodeName
		Recipe   recipeName
		Delivery DeliveryObj
	}

	// Report structs
	ReportRecipeCounter struct {
		Recipe string `json:"recipe"`
		Count  int64  `json:"count"`
	}

	BusiestPostcode struct {
		Postcode      string `json:"postcode"`
		DeliveryCount int64  `json:"delivery_count"`
	}

	CountForPostCode struct {
		Postcode      string `json:"postcode"`
		From          string `json:"from"`
		To            string `json:"to"`
		DeliveryCount int64  `json:"delivery_count"`
	}

	Report struct {
		UniqueRecipeCount       int                   `json:"unique_recipe_count"`
		CountPerRecipe          []ReportRecipeCounter `json:"count_per_recipe"`
		BusiestPostcode         BusiestPostcode       `json:"busiest_postcode"`
		CountPerPostcodeAndTime CountForPostCode      `json:"count_per_postcode_and_time"`
		MatchByName             []string              `json:"match_by_name"`
	}
)
