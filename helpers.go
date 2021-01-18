package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func fixedSizeArrToString(s []byte) string {
	return string(s[:bytes.IndexByte(s, 0)])
}

// convertTime24IntTo12Str converts workTime type 24h int representation to the 12h string
func convertTime24IntTo12Str(time int64) (string, error) {
	if time > 24 || time < 0 {
		return "", fmt.Errorf("wrong time format %d", time)
	}

	if time > 12 {
		return fmt.Sprintf("%dPM", time-12), nil
	}

	if time < 12 && time > 0 {
		return fmt.Sprintf("%dAM", time), nil
	}

	if time == 0 {
		return "12AM", nil
	}

	if time == 12 {
		return "12PM", nil
	}

	return "", fmt.Errorf("wrong time format %d", time)
}

// convertTimeByteToInt converts workTime type time representation to the int
func convertTimeByteToInt(time workTime) (res int64, err error) {
	switch time {
	case workTime{'1', 'A', 'M'}:
		res = 1
	case workTime{'2', 'A', 'M'}:
		res = 2
	case workTime{'3', 'A', 'M'}:
		res = 3
	case workTime{'4', 'A', 'M'}:
		res = 4
	case workTime{'5', 'A', 'M'}:
		res = 5
	case workTime{'6', 'A', 'M'}:
		res = 6
	case workTime{'7', 'A', 'M'}:
		res = 7
	case workTime{'8', 'A', 'M'}:
		res = 8
	case workTime{'9', 'A', 'M'}:
		res = 9
	case workTime{'1', '0', 'A', 'M'}:
		res = 10
	case workTime{'1', '1', 'A', 'M'}:
		res = 11
	case workTime{'1', '2', 'A', 'M'}:
		res = 0
	case workTime{'1', 'P', 'M'}:
		res = 13
	case workTime{'2', 'P', 'M'}:
		res = 14
	case workTime{'3', 'P', 'M'}:
		res = 15
	case workTime{'4', 'P', 'M'}:
		res = 16
	case workTime{'5', 'P', 'M'}:
		res = 17
	case workTime{'6', 'P', 'M'}:
		res = 18
	case workTime{'7', 'P', 'M'}:
		res = 19
	case workTime{'8', 'P', 'M'}:
		res = 20
	case workTime{'9', 'P', 'M'}:
		res = 21
	case workTime{'1', '0', 'P', 'M'}:
		res = 22
	case workTime{'1', '1', 'P', 'M'}:
		res = 23
	case workTime{'1', '2', 'P', 'M'}:
		res = 12
	default:
		err = fmt.Errorf("wrong time format: %v", time)
	}

	return
}

func initialization() (dataFileName string, conf Config, err error) {
	if len(os.Args) < 3 {
		err = errors.New("need to provide the config and data file [-conf, -data]")
		return
	}

	f := flag.NewFlagSet("config", flag.ExitOnError)
	configFile := f.String("conf", "", "config file name")
	dataFile := f.String("data", "", "data file name")
	err = f.Parse(os.Args[1:])
	if err != nil {
		return
	}

	if *configFile == "" || *dataFile == "" {
		err = errors.New("need to provide the config and data file [-conf, -data]")
		return
	}

	dataFileName = *dataFile

	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		return
	}

	return
}
