package validate

import (
	"testing"
)

func TestValidateDiffSizeMoreThan50PercMustReturnError(t *testing.T){
	//The first file has 100k
	//The second file has 10k
	//The diff size between the first and second file is more than 10%, so no make sense to compare them
	err := validateSize(100, 10)
	if err == nil {
		t.Error("Diff size more than 50% must return error")
	}
}

func TestValidateDiffSizeLessThan50PercMustReturnNoError(t *testing.T){
	//The first file has 100k
	//The second file has 91k
	//The diff size between the first and second file is less than 50%, so these files must be compared
	err := validateSize(100, 51)
	if err != nil {
		t.Error("Compare diff json files must return NO error")
	}
}
