package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func getData() string {
	content, err := ioutil.ReadFile("try.json")
	if err != nil {
		fmt.Println("ReadFileErr")
		log.Fatal(err)
	}

	var tem []ResponseData
	err = json.Unmarshal(content, &tem)
	if err != nil {
		fmt.Println("UnmarshallErr")
		log.Fatal(err)
	}

	resStruct := ResponseStruct{
		"success", len(tem), tem}

	res, err := json.Marshal(resStruct)
	if err != nil {
		fmt.Println("MarshallErr")
		log.Fatal(err)
	}
	return string(res)
}

func getRestData(recid int) string {

	content, err := ioutil.ReadFile("restData.json")
	if err != nil {
		fmt.Println("ReadFileErr")
		log.Fatal(err)
	}

	var tem RestResponseStruct

	err = json.Unmarshal(content, &tem)
	if err != nil {
		fmt.Println("UnmarshallErr")
		log.Fatal(err)
	}

	if recid > 0 {
		fmt.Println(recid)
		var singleRec ResponseData
		for _, val := range tem.ResponseData {
			if recid == val.Recid {
				singleRec = val
				break
			}
		}
		res, err := json.Marshal(singleRec)
		if err != nil {
			fmt.Println("MarshallErr")
			log.Fatal(err)
		}
		return string(res)
	}

	//resStruct := RestResponseStruct{tem}

	res, err := json.Marshal(tem)
	if err != nil {
		fmt.Println("MarshallErr")
		log.Fatal(err)
	}
	return string(res)
}
