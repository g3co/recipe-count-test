package main

// ParseBuffer parses raw data to the RecipeObj
func ParseBuffer(b buffType) (res RecipeObj, err error) {
	code := postCodeName{}
	recipeName := recipeName{}
	startTime := workTime{}
	endTime := workTime{}
	day := weekDay{}

	for _, item := range b {
		// cut quotes and comma if exist
		offset := 1
		if item[len(item)-1] == ',' {
			offset = 2
		}

		// for performance optimization check first 6 chars,
		// without reallocate memory for strings
		// it's will be enough for separate fields
		itemFieldName := [6]byte{}
		copy(itemFieldName[:], item[3:9])

		// [postco]de": " - field recipeName length with quotes and spaces 15 chars
		if [6]byte{'p', 'o', 's', 't', 'c', 'o'} == itemFieldName {
			copy(code[:], item[15:len(item)-offset])
		}

		// [recipe]": " - field recipeName length with quotes and spaces 13 chars
		if [6]byte{'r', 'e', 'c', 'i', 'p', 'e'} == itemFieldName {
			copy(recipeName[:], item[13:len(item)-offset])
		}

		// [delive]ry": " - field recipeName dlength with quotes and spaces 15 chars
		if [6]byte{'d', 'e', 'l', 'i', 'v', 'e'} == itemFieldName {
			deliveryContent := item[15 : len(item)-offset]

			step := 0
			counter := 0
			for _, item := range deliveryContent {
				// space symbol is separator
				if item == ' ' {
					step++
					counter = 0
					continue
				}

				// dash "-" symbol is 2, need to skip
				switch step {
				case 0:
					day[counter] = item
				case 1:
					startTime[counter] = item
				case 3:
					endTime[counter] = item
				}

				counter++
			}
		}
	}

	startTimeInt, err := convertTimeByteToInt(startTime)
	if err != nil {
		return
	}

	endTimeInt, err := convertTimeByteToInt(endTime)
	if err != nil {
		return
	}

	res = RecipeObj{
		Recipe:   recipeName,
		PostCode: code,
		Delivery: DeliveryObj{
			WeekDay: day,
			Start:   startTimeInt,
			End:     endTimeInt,
		},
	}

	return
}
