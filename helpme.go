package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
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
	if n == 0 {
		n = len(results)
	}
	for key := range results[:n] {
		conv = results[key].(map[string]interface{})
		if conv["is_answered"].(bool) {
			conv["title"] = fmt.Sprintf("[ANSWERED] %s", conv["title"])
		}
		fmt.Printf("\nTitre: %s\nLien: %s\nTags: ", conv["title"], conv["link"])
		for _, tag := range conv["tags"].([]interface{}) {
			fmt.Printf("%s, ", tag)
		}
		fmt.Printf("\nScore: %2.f\n", conv["score"])
	}
}

func displayResultsSorted(results []interface{}, n int) {
	var resList []map[string]interface{}

	if n == 0 {
		n = len(results)
	}

	for key := range results[:n] {
		resList = append(resList, results[key].(map[string]interface{}))
	}
	sort.Slice(resList, func(i, j int) bool {
		return resList[j]["score"].(float64) > resList[i]["score"].(float64)
	})

	for _, res := range resList {
		if !res["is_answered"].(bool) {
			res["title"] = fmt.Sprintf("[UN-ANSWERED] %s", res["title"])
		}
		fmt.Printf("\nTitre: %s\nLien: %s\nTags: ", res["title"], res["link"])
		for _, tag := range res["tags"].([]interface{}) {
			fmt.Printf("%s, ", tag)
		}
		fmt.Printf("\nScore: %2.f\n", res["score"])
	}
}

/*
	Fonction qui :
	-	Effectue la requête de méthode GET vers l'API de stackexchange
	-	Traîte la réponse de l'API
	-	Affiche les résultats dans le terminal
*/
func fetch(url string, n int, sortFlag bool) {
	var data map[string]interface{}
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(result, &data); err != nil {
		panic(err)
	}
	items := data["items"].([]interface{})
	if sortFlag {
		displayResultsSorted(items, n)
	} else {
		displayResults(items, n)
	}
}

func main() {
	args := os.Args
	sortFlag := false

	if len(args) > 1 {
		url := buildURL(args[1])
		if len(args) == 4 && args[3] == "-s" {
			sortFlag = true
		}
		if len(args) == 3 {
			n, err := strconv.Atoi(args[2])
			if err != nil {
				panic("The second argument is not a number")
			}
			fetch(url, n, sortFlag)
		} else {
			fetch(url, 0, sortFlag)
		}
	} else {
		fmt.Println("Error: no arguments were provided")
	}
}
