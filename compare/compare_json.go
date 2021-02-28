package compare

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

const (
	noImportantValue = 1
	problemToReadJsonFile = "problem to read json file %v"
	problemToRetrieveJsonData = "problem to retrieve json data"
)

func Apply(jsonFile string, jsonFileBkp string) (bool, error) {

	jsonHash := transformJsonToHash(jsonFile)

	file, _ := OpenFile(jsonFileBkp)
	defer file.Close()

	jsonResult := json.NewDecoder(bufio.NewReader(file))

	// read open bracket
	_, err := jsonResult.Token()
	if err != nil {
		return false, errors.New(problemToRetrieveJsonData)
	}

	for jsonResult.More() {

		var jsonObject map[string]interface{}
		jsonResult.Decode(&jsonObject)
		jsonNodeHash := transformJsonNodeToHash(jsonObject)

		//If this hash doesn't exist into the first Json file they are different
		_, found := jsonHash[jsonNodeHash]
		if !found {
			return false, nil
		}
	}

	return true, nil
}

func OpenFile(jsonFile string) (*os.File, error) {
	jsonData, err := os.Open(jsonFile)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(problemToReadJsonFile, jsonFile))
	}
	return jsonData, nil
}

func transformJsonToHash(jsonFile string) map[string]int {

	file, _ := OpenFile(jsonFile)
	defer file.Close()

	jsonResult := json.NewDecoder(bufio.NewReader(file))

	// read open bracket
	_, err := jsonResult.Token()
	if err != nil {
		log.Fatal(problemToRetrieveJsonData, err)
	}

	jsonHash := map[string]int{}
	for jsonResult.More(){

		var jsonObject map[string]interface{}
		jsonResult.Decode(&jsonObject)
		jsonNodeHash := transformJsonNodeToHash(jsonObject)
		jsonHash[jsonNodeHash] = noImportantValue
	}

	return jsonHash
}

func transformJsonNodeToHash(data map[string]interface{}) string {

	resultMap := map[string]string{}
	for key, value := range data {

		switch reflect.ValueOf(value).Kind().String() {
		case "map":{
			mapHandler(value, resultMap, key)
		}
		case "slice": {
			//To handle arrays
			sliceHandler(value, resultMap, key)
		}
		default:
			resultMap[key] = fmt.Sprintf("%v", value)
		}
	}

	resultMapSorted := resultMapSorted(resultMap)

	//join the strings
	groupResultMapContent := groupResultMapContent(resultMapSorted, resultMap)
	hashContent := fmt.Sprintf("%x", md5.Sum([]byte(groupResultMapContent)))

	return hashContent
}

func mapHandler(value interface{}, resultMap map[string]string, key string) {
	resultMap[key] = transformJsonNodeToHash(value.(map[string]interface{}))
}

func sliceHandler(value interface{}, resultMap map[string]string, key string) {
	arrayData := value.([]interface{})
	switch reflect.ValueOf(arrayData[0]).Kind().String() {
	case "map":
		{

			groupJsonHash := sliceMapHandler(&arrayData)
			//apply hash
			strHash := fmt.Sprintf("%x", md5.Sum([]byte(groupJsonHash)))
			resultMap[key] = strHash
		}
	default:
		{
			groupJsonData := sliceStringArrayHandler(&arrayData)
			//apply hash
			strHash := fmt.Sprintf("%x", md5.Sum([]byte(groupJsonData)))
			resultMap[key] = strHash
		}
	}
}

func sliceStringArrayHandler(arrayData *[]interface{}) string {
	arrayString := make([]string, 0)
	for _, elementValue := range *arrayData {
		arrayString = append(arrayString, fmt.Sprintf("%v", elementValue))
	}

	sort.Strings(arrayString)
	groupJsonData := strings.Join(arrayString, "|")
	return groupJsonData
}

func sliceMapHandler(arrayData *[]interface{}) string {
	arrayHash := make([]string, 0)
	for _, elementValue := range *arrayData {
		arrayHash = append(arrayHash, transformJsonNodeToHash(elementValue.(map[string]interface{})))
	}
	sort.Strings(arrayHash)
	groupJsonHash := strings.Join(arrayHash, "|")
	return groupJsonHash
}

func groupResultMapContent(resultMapSorted []string, resultMap map[string]string) string {
	values := make([]string, 0)
	for _, value := range resultMapSorted {
		values = append(values, resultMap[value])
	}
	return strings.Join(values, "|")
}

func resultMapSorted(resultMap map[string]string) []string {
	attributes := make([]string, 0)
	for key, _ := range resultMap {
		attributes = append(attributes, key)
	}
	sort.Strings(attributes)
	return attributes
}
