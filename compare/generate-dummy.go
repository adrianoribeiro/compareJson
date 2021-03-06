package compare

import "fmt"

const (
	start = "[{"
	end = "}]"
	address = "    \"address\": {\n      \"street\": \"Street %d\",\n      \"postalCode\": \"00000-1%d\",\n      \"city\": \"Paris\",\n      \"countryCode\": \"FRA\",\n      \"country\": \"France\",\n      \"text\": \"Text %d\"\n    }"
)

func GenerateDummy(numberItems int){

	//var str string
	for i:= 0;i<numberItems;i++ {
		println(fmt.Sprintf(address, i, i, i))
	}
}
