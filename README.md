# helpme

Un petit outil en ligne de commande écrit en Go facilitant la recherche de solutions liées à une erreur.
Ce programme utilise l'API de [Stackexchange](https://api.stackexchange.com/docs).

## Examples
```bash
$ helpme segfault 2

Titre: Interpreting segfault messages
Lien: https://stackoverflow.com/questions/2549214/interpreting-segfault-messages
Score: 56.000000

Titre: What is a bus error?
Lien: https://stackoverflow.com/questions/212466/what-is-a-bus-error
Score: 252.000000
```
Cette commande affichera sur la sortie standard les 2 pages de stackexchange les plus pertinentes par rapport à l'erreur entrée.

## Syntaxe:
```bash
$ helpme <erreur> <nombre de pages>
```

## Installation

L'installation nécessite la présence de Go sur votre système -> [Installation de go](https://golang.org/dl/) 
```bash
$ go build github.com/BoyerDamien/helpme && sudo mv helpme /usr/bin
```
