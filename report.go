package main

import (
	"encoding/json"
	"sort"
	"strings"
)

func report(recipes recipeMap, postCodes postCodeMap, pc PostCodeCounter, findRecipes []string) (string, error) {
	recipeCounter := make([]ReportRecipeCounter, 0, len(recipes))
	searchingRecipes := make([]string, 0, len(recipes))
	for n, v := range recipes {
		recipe := ReportRecipeCounter{
			Recipe: fixedSizeArrToString(n[:]),
			Count:  v,
		}

		recipeCounter = append(recipeCounter, recipe)

		for _, name := range findRecipes {
			if strings.Contains(recipe.Recipe, name) {
				searchingRecipes = append(searchingRecipes, recipe.Recipe)
				break
			}
		}
	}

	sort.SliceStable(recipeCounter, func(i, j int) bool {
		return recipeCounter[i].Recipe < recipeCounter[j].Recipe
	})

	sort.SliceStable(searchingRecipes, func(i, j int) bool {
		return searchingRecipes[i] < searchingRecipes[j]
	})

	var postCode postCodeName
	var maxDeliveryCounts int64
	for n, v := range postCodes {
		if maxDeliveryCounts < v {
			maxDeliveryCounts = v
			postCode = n
		}
	}

	countForPostCodeFrom, err := convertTime24IntTo12Str(pc.Start)
	if err != nil {
		return "", err
	}

	countForPostCodeTo, err := convertTime24IntTo12Str(pc.End)
	if err != nil {
		return "", err
	}

	report := Report{
		UniqueRecipeCount: len(recipeCounter),
		CountPerPostcodeAndTime: CountForPostCode{
			Postcode:      pc.GetStrPostCode(),
			From:          countForPostCodeFrom,
			To:            countForPostCodeTo,
			DeliveryCount: pc.Counter,
		},
		BusiestPostcode: BusiestPostcode{
			Postcode:      fixedSizeArrToString(postCode[:]),
			DeliveryCount: maxDeliveryCounts,
		},
		CountPerRecipe: recipeCounter,
		MatchByName:    searchingRecipes,
	}

	jsonReport, _ := json.Marshal(report)
	return string(jsonReport), nil
}
