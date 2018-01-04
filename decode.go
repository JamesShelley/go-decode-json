// How to decode a JSON string in Go
package go_decode_json

import (
	"fmt"
	"encoding/json" // Encoding and Decoding Package
)

var JSON = `{
    "name":"James Shelley",
    "university":"Bournemouth",
	"degree":"Software Engineering",
    "websites":{
        "personal":"jshelley.uk",
        "work":"jsweb.design"
    },
    "email":"jshelley@tuta.io"
}`

func main() {
	// Unmarshal the JSON string into info map variable.
	var info map[string]interface{}
	json.Unmarshal([]byte(JSON),&info)

	// Print the output from info map.
	fmt.Println("Name:",info["name"])
	fmt.Println("University:",info["university"])
	fmt.Println("Degree:",info["degree"])
	fmt.Println("Email:",info["email"])
	fmt.Println("Personal Website:",info["websites"].(map[string]interface{})["personal"])
	fmt.Println("Work Website:",info["websites"].(map[string]interface{})["work"])
}