# Compare json files
Given two json files this program will check if they are equals.

## A simple example of equal JSON files

![Alt text](json-example.png?raw=true)

## What was my approach
Lets compare these two json data :  
jsonOne.json
```
[
  {
    "id": "xpto001",
    "value": "Adriano",
    "address":
    {
      "street": "11 Rue du Grenier Saint-Lazare",
      "postalCode": "75003-111",
      "city": "Paris",
      "countryCode": "FRA",
      "country": "France",
      "text": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  },
  {
    "id": "abcd001",
    "value": "Juca",
    "address":
    {
      "street": "22 Rue du Grenier Saint-Lazare",
      "postalCode": "75003-222",
      "city": "Paris",
      "countryCode": "FRA",
      "country": "France",
      "text": "22 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  }
]
```
jsonTwo.json
```
[
  {
    "address":
    {
      "city": "Paris",
      "postalCode": "75003-222",
      "countryCode": "FRA",
      "street": "22 Rue du Grenier Saint-Lazare",
      "country": "France",
      "text": "22 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    },
    "value": "Juca",
    "id": "abcd001"
  },
  {
    "value": "Adriano",
    "id": "xpto001",
    "address":
    {
      "street": "11 Rue du Grenier Saint-Lazare",
      "postalCode": "75003-111",
      "city": "Paris",
      "countryCode": "FRA",
      "country": "France",
      "text": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  }
]
```
The first challenge I had to solve was the randomly order.  I'll take the jsonOne.json as example
```
[
  {
    "id": "xpto001",
    "value": "Adriano",
    "address":
    {
      "street": "11 Rue du Grenier Saint-Lazare",
      "postalCode": "75003-111",
      "city": "Paris",
      "countryCode": "FRA",
      "country": "France",
      "text": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  },
  {
    "id": "abcd001",
    "value": "Juca",
    "address":
    {
      "street": "22 Rue du Grenier Saint-Lazare",
      "postalCode": "75003-222",
      "city": "Paris",
      "countryCode": "FRA",
      "country": "France",
      "text": "22 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  }
]
```
Note that the address fields are ordered
```
[
  {
    "id": "xpto001",
    "value": "Adriano",
    "address":
    {
      "city": "Paris",
      "country": "France",
      "countryCode": "FRA",
      "postalCode": "75003-111",
      "street": "11 Rue du Grenier Saint-Lazare",
      "text": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  },
  {
    "id": "abcd001",
    "value": "Juca",
    "address":
    {
      "city": "Paris",
      "country": "France",
      "countryCode": "FRA",
      "postalCode": "75003-222",
      "street": "22 Rue du Grenier Saint-Lazare",
      "text": "22 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  }
]
```

The second challenge I had to solve was the approach to compare the data I think to create a hash
is more simple instead of comparing field by field, take a look at the address transformation:  

jsonOne.json with the address sorted
```
[
  {
    "id": "xpto001",
    "value": "Adriano",
    "address":
    {
      "city": "Paris",
      "country": "France",
      "countryCode": "FRA",
      "postalCode": "75003-111",
      "street": "11 Rue du Grenier Saint-Lazare",
      "text": "11 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  },
  {
    "id": "abcd001",
    "value": "Juca",
    "address":
    {
      "city": "Paris",
      "country": "France",
      "countryCode": "FRA",
      "postalCode": "75003-222",
      "street": "22 Rue du Grenier Saint-Lazare",
      "text": "22 Rue du Grenier Saint-Lazare, 75003 Paris, France"
    }
  }
]
```
Notice that the address was changed for a hash
```
[
  {
    "id": "xpto001",
    "value": "Adriano",
    "address": "42577a52d4c67e25ed7b7b77274222ab"
  },
  {
    "id": "abcd001",
    "value": "Juca",
    "address": "b4f76897da37c065912902487a95a64b"
  }
]
```
And recursively I also sorted the fields:
```
[
  {
    "address": "42577a52d4c67e25ed7b7b77274222ab"
    "id": "xpto001",
    "value": "Adriano",
  },
  {
    "address": "b4f76897da37c065912902487a95a64b"
    "id": "abcd001",
    "value": "Juca",
  }
]
```
And... created a hash for each node and store all of these hash into a MAP
```
[
  4cad05aff79a305327553be09ef348ac,
  6c93155befa80e9a5eba58ec7289c3d7 
]
```
At the end to compare the files, for each node from the jsonTwo.json I create a hash and check if it exists in the first map hash.
## How to run the test
```
go test ./...
```

## How to run
```
go run main.go <jsonFile1.json> <jsonFile2.json>
```

### Suggestion to test
```
go run main.go dummy-json/json_pdf_sample1.json dummy-json/json_pdf_sample2.json
```