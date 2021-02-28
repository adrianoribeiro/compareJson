package compare

import (
	"strings"
	"testing"
)

func TestSimpleTransformJsonToHashFromTenJsonShouldReturnTenHashs(t *testing.T){

	jsonData := transformJsonToHash("../dummy-json/json_simple_schema.json")

	if len(jsonData) != 10 {
		t.Error("A file with two json data must return two hash.")
	}
}

func TestComplexTransformJsonToHashFromSameJsonDataShouldReturnSameHashs(t *testing.T){

	jsonData1 := transformJsonToHash("../dummy-json/json_complex_one_json_schema.json")
	jsonData2 := transformJsonToHash("../dummy-json/json_complex_one_json_schema.json")

	if !strings.EqualFold(getFirstElement(&jsonData1), getFirstElement(&jsonData2)) {
		t.Error("A json data must return the same hash every time.")
	}
}

func TestComplexTransformJsonToHashCompareSameJsonSortedOrNotShouldReturnSameHashs(t *testing.T){

	jsonData1 := transformJsonToHash("../dummy-json/json_complex_one_json_schema.json")
	jsonData2 := transformJsonToHash("../dummy-json/json_complex_one_json_schema_clone.json")

	if !strings.EqualFold(getFirstElement(&jsonData1), getFirstElement(&jsonData2)) {
		t.Error("Compare same json data sorted or not must return the same hash.")
	}
}

func TestComplexTransformJsonToHashCompareDiffJsonSortedOrNotShouldReturnDiffHashs(t *testing.T){

	jsonData1 := transformJsonToHash("../dummy-json/json_complex_one_json_schema.json")
	jsonData2 := transformJsonToHash("../dummy-json/json_complex_one_json_schema_diff.json")

	if strings.EqualFold(getFirstElement(&jsonData1), getFirstElement(&jsonData2)) {
		t.Error("Compare diff json data sorted or not must return the diff hash.")
	}
}

func TestComplexTransformJsonToHashFromTwoJsonShouldReturnTwoHashs(t *testing.T){

	jsonData := transformJsonToHash("../dummy-json/json_complex_schema.json")

	if len(jsonData) != 2 {
		t.Error("A file with two json data must return two hash.")
	}
}

/*
func TestComplexTransformJsonToHashFromTwoJsonShouldReturnTwoHashs2(t *testing.T){

	jsonData := transformJsonToHash("../dummy-json/json_before_1G.json")

	if len(jsonData) != 4000000 {
		t.Error("A file with two json data must return two hash.")
	}
}
*/

func TestSimpleCompareSimilarJsonFilesMustReturnTrue(t *testing.T){

	isEqual, _ := Apply("../dummy-json/json_before_2_3M.json", "../dummy-json/json_before_2_3M_clone_revert_order.json")

	if !isEqual {
		t.Error("Compare same json files must return true")
	}
}

func TestSimpleCompareDiffJsonFilesMustReturnFalse(t *testing.T){

	isEqual, _ := Apply("../dummy-json/json_before_2_3M.json", "../dummy-json/json_before_2_3M_clone_revert_order_diff.json")

	if isEqual {
		t.Error("Compare diff json files must return false")
	}
}

func TestComplexCompareSameJsonFilesMustReturnTrue(t *testing.T){
	isEqual, _ := Apply("../dummy-json/json_complex_one_json_schema.json", "../dummy-json/json_complex_one_json_schema.json")

	if !isEqual {
		t.Error("Compare same json files must return true")
	}
}

func TestComplexCompareSimilarJsonFilesMustReturnTrue(t *testing.T){
	isEqual, _ := Apply("../dummy-json/json_complex_one_json_schema.json", "../dummy-json/json_complex_one_json_schema_clone.json")

  	if !isEqual {
		t.Error("Compare similar json files must return true")
	}
}

func TestMoreComplexOneCompareSimilarJsonFilesMustReturnTrue(t *testing.T){

	jsonData1 := transformJsonToHash("../dummy-json/json_more_complex_one_json_schema.json")
	jsonData2 := transformJsonToHash("../dummy-json/json_more_complex_one_json_schema_clone.json")

	if !strings.EqualFold(getFirstElement(&jsonData1), getFirstElement(&jsonData2)) {
		t.Error("Compare similar json files must return true")
	}
}

func TestMoreComplexCompareSameJsonFilesMustReturnTrue(t *testing.T){

	isEqual, _ := Apply("../dummy-json/json_other_more_complex_json_schema.json"	, "../dummy-json/json_other_more_complex_json_schema.json")

	if !isEqual {
		t.Error("Compare similar json files must return true")
	}
}

func TestMoreComplexCompareSimilarJsonFilesMustReturnTrue(t *testing.T){

	isEqual, _ := Apply("../dummy-json/json_other_more_complex_json_schema.json", "../dummy-json/json_other_more_complex_json_schema_clone.json")

	if !isEqual {
		t.Error("Compare similar json files must return true")
	}
}

func TestMoreComplexCompareDiffJsonFilesMustReturnFalse2(t *testing.T){

	isEqual, _ := Apply("../dummy-json/json_other_more_complex_json_schema.json", "../dummy-json/json_other_more_complex_json_schema_diff.json")

	if isEqual {
		t.Error("Compare similar json files must return trues")
	}
}

func TestMoreComplexCompareDiffJsonFilesMustReturnFalse(t *testing.T){

	jsonData1 := transformJsonToHash("../dummy-json/json_more_complex_one_json_schema.json")
	jsonData2 := transformJsonToHash("../dummy-json/json_more_complex_one_json_schema_diff.json")

	if strings.EqualFold(getFirstElement(&jsonData1), getFirstElement(&jsonData2)) {
		t.Error("Compare diff json files must return false")
	}
}

//Auxiliary functions
func getFirstElement(jsonData *map[string]int) string {
	var key string
	for key, _ = range *jsonData {}

	return key
}

