package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	dataFile, conf, err := initialization()
	if err != nil {
		log.Fatalf("config file initialization err: %s", err)
	}

	pcCounter := NewPostCodeCounter(
		conf.PostCodeCounter.PostCode,
		conf.PostCodeCounter.From,
		conf.PostCodeCounter.To,
	)

	file, err := os.Open(dataFile)
	if err != nil {
		log.Fatalf("data file open err: %s", err)
	}

	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	// fixed-size maps and arrays initialization to prevent a dynamic memory allocation
	recipes := make(recipeMap, recipeMapSize)
	postCodes := make(postCodeMap, postCodeMapSize)

	var buff buffType
	for i := 0; i < fieldsQuantity; i++ {
		buff[i] = make([]byte, 0, bufferSize)
	}

	var i int
	for scanner.Scan() {
		text := scanner.Bytes()
		switch text[0] {
		case ']', '[', '{': // skip json struct symbols
			continue
		case '}': // close bracket, buffer filled up
			i = 0
			obj, err := ParseBuffer(buff)
			if err != nil {
				log.Fatalf("parse error: %s (%s)", err.Error(), buff)
			}
			recipes[obj.Recipe]++
			postCodes[obj.PostCode]++
			pcCounter.Check(obj.PostCode, obj.Delivery.Start, obj.Delivery.End)
		default:
			buff[i] = append(buff[i][:0], text...)
			i++
		}
	}

	report, err := report(recipes, postCodes, pcCounter, conf.SearchByName)
	if err != nil {
		log.Fatalf("Report generation err: %s", err)
	}

	fmt.Println(report)
}
