package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	first  = "https://api.stackexchange.com/2.2/search/advanced?order=desc"
	second = "&accepted=True&site=stackoverflow"
)

type Request struct {
	content string
	url     string
	sort    bool
	n       int
}

/* Méthode qui construit une url de l'API stackexchange à partir d'un contenu de recherche*/
func (v *Request) buildURL() {
	v.content = strings.ReplaceAll(v.content, " ", "+")
	if v.sort {
		v.url = first + "&sort=votes" + "&q=" + v.content + second
	} else {
		v.url = first + "&sort=relevance" + "&q=" + v.content + second
	}
}

/* Méthode qui affiche les resultats formatés de l'API stackexchange */
/* 	-	Titre
-	Lien de la question
-	Score de la question
*/
func (v *Request) displayResults(results []interface{}) {
	var conv map[string]interface{}
	if (v.n == 0) || (v.n > len(results)) {
		v.n = len(results)
	}
	for key := range results[:v.n] {
		conv = results[key].(map[string]interface{})
		fmt.Printf("\nTitre: %s\nLien: %s\nScore: %.2f\n", conv["title"], conv["link"], conv["score"])
	}
}

/*
	Méthode qui :
	-	Effectue la requête de méthode GET vers l'API de stackexchange
	-	Traîte la réponse de l'API
	-	Affiche les résultats dans le terminal
*/
func (v *Request) fetch() {
	var data map[string]interface{}
	response, err := http.Get(v.url)
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
	v.displayResults(items)
}
