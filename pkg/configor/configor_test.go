package configor

import (
	"encoding/json"
	"fmt"
)

func ExampleGet() {
	configuration := Get()

	data, _ := json.MarshalIndent(configuration, "", "  ")
	fmt.Println(string(data))

	// Output:
	// {
	// 	"Mode": "production",
	// 	"Cors": {
	// 		"AllowOrigins": [
	// 			"GET",
	// 			"POST"
	// 		],
	// 		"AllowMethods": null,
	// 		"AllowHeaders": null
	// 	},
	// 	"StringArray": [
	// 		"foo",
	// 		"bar"
	// 	],
	// 	"StructArray": [
	// 		{
	// 			"Name": "foo"
	// 		},
	// 		{
	// 			"Name": "poop"
	// 		}
	// 	]
	// }
}
