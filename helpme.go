package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/* Fonction qui construit une url de l'API stackexchange à partir d'un contenu de recherche*/
func buildURL(content string) string {
	first := "https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=relevance&q="
	second := "&accepted=True&site=stackoverflow"
	return first + strings.Replace(content, " ", "+", -1) + second
}

/* Fonction qui affiche les resultats formatés de l'API stackexchange */
/* 	-	Titre
-	Lien de la question
-	Score de la question
*/
func displayResults(results []interface{}, n int) {
	var conv map[string]interface{}
	if (n == 0) || (n > len(results)) {
		n = len(results)
	}
	for key := range results[:n] {
		conv = results[key].(map[string]interface{})
		fmt.Printf("\nTitre: %s\nLien: %s\nScore: %f\n", conv["title"], conv["link"], conv["score"])
	}
}

/*
	Fonction qui :
	-	Effectue la requête de méthode GET vers l'API de stackexchange
	-	Traîte la réponse de l'API
	-	Affiche les résultats dans le terminal
*/
func fetch(url string, n int) {
	var data map[string]interface{}
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(result, &data); err != nil {
		panic(err)
	}
	items := data["items"].([]interface{})
	displayResults(items, n)
}

func main() {
	args := os.Args
	if len(args) > 1 {
		url := buildURL(args[1])
		if len(args) == 3 {
			n, err := strconv.Atoi(args[2])
			if err != nil {
				panic("The second argument is not a number")
			}
			fetch(url, n)
		} else {
			fetch(url, 0)
		}
	} else {
		fmt.Println("Error: no arguments were provided")
	}
}
