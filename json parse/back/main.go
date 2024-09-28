package main

import (
	// "encoding/json"
	// "encoding/base64"
	"fmt"
	"strings"
	"back/golang"
	"back/handlerjson"
	"back/java"
	"back/javascript"
	"back/kotlin"
	"back/cpp"
	"back/dart"
	"back/dotnet"
	"back/python"
	"back/ruby"
	"back/swift"
	"back/rust"
	"back/typescript"
)

func main() {
	// Sample filled JSON string

	base64JSON:=`ewoJCSJpbnB1dCI6IHsKCQkJImdhYW5kIjogW1siMSIsICIyIiwgIjMiXSwgWyI0IiwgIjUiLCAiNiJdXSwKCQkJInRpdHMiOiBbIjciLCAiOCIsICI5Il0sCgkJCSJjcmVhbXBpZSI6IFtbWzEwLCAxMV0sIFsxMiwgMTNdXSwgW1sxNCwgMTVdLCBbMTYsIDE3XV1dLAoJCQkiY2h1ZCI6IFtbMTgsIDE5LCAyMF0sIFsyMSwgMjIsIDIzXV0sCgkJCSJ2YWdpbmEiOiBbMjQsIDI1LCAyNl0sCgkJCSJsdW5kIjogWzI3LCAyOCwgMjldLAoJCQkicGVuaXMiOiBbNywgOF0sCgkJCSJjbGl0cyI6IFtbWzEwLCAxMV0sIFsxMiwgMTNdXV0KCQl9LAoJCSJvdXRwdXQiOiB7CgkJCSJib29icyI6IFtbWzMwLCAzMV0sIFszMiwgMzNdXSwgW1szNCwgMzVdLCBbMzYsIDM3XV1dLAoJCQkicHVkaSI6IFtbWzEwLCAxMV0sIFsxMiwgMTNdXV0KCgkJfQoJfQ==`
	

	// jsonString := `{
	// 	"input": {
	// 		"gaand": [["1", "2", "3"], ["4", "5", "6"]],
	// 		"tits": ["7", "8", "9"],
	// 		"creampie": [[[10, 11], [12, 13]], [[14, 15], [16, 17]]],
	// 		"chud": [[18, 19, 20], [21, 22, 23]],
	// 		"vagina": [24, 25, 26],
	// 		"lund": [27, 28, 29],
	// 		"penis": [7, 8],
	// 		"clits": [[[10, 11], [12, 13]]]
	// 	},
	// 	"output": {
	// 		"boobs": [[[30, 31], [32, 33]], [[34, 35], [36, 37]]],
	// 		"pudi": [[[10, 11], [12, 13]]]

	// 	}
	// }`


	// array:=declarationGo(base64JSON)
	array := declarationTypescript(base64JSON)
	fmt.Println(array)

}

func declarationGo(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := golang.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := golang.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func declarationJava(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := java.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := java.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}

func declarationJavascript(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := javascript.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := javascript.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}

func declarationKotlin(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := kotlin.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := kotlin.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}



func declarationCplusplus(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := cpp.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := cpp.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func declarationDart(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := dart.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := dart.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func declarationCSharp(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := dotnet.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := dotnet.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func declarationPython(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := python.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := python.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func declarationRuby(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := ruby.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := ruby.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func declarationSwift(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := swift.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := swift.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func declarationRust(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := rust.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := rust.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}



func declarationTypescript(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := typescript.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := typescript.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}