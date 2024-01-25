### Install > ```go get https://github.com/George2901/String_formater```
### Import > ``` import sf "github.com/George2901/String_formater" ```

### Full exemple
```
package main

import sf "github.com/George2901/String_formater"

func main() {
	a := "first"
	b := "succesfully"
	returned_string, err := sf.format(`{} test {}`, a, b)
	if err != nil {
		log.Fatal(err)
	}
	println(returned_string)
}
```
