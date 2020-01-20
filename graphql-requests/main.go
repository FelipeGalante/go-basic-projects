package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("GraphQL")
	url := "https://graphql-pokemon.now.sh"
	var jsonStr = []byte(`{"query": "query($pokemonName: String!) {pokemon(name: $pokemonName) {number}}","variables":"{\"pokemonName\":\"Charmander\"}"}`)

	res, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)
}
