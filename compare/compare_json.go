package compare

import (
	"bufio"
	"encoding/json"
	"errors"
)

func Exec(jsonFile string, jsonFileBkp string) (bool, error) {

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

