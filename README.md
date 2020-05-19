# helpme

Un petit outil en ligne de commande écrit en Go facilitant la recherche de solutions liées à une erreur.
Ce programme utilise l'API de [Stackexchange](https://api.stackexchange.com/docs).

## Examples
```bash
$ helpme "indentation error" 2                                                                                    (base) 

Titre: Indentation error
Lien: https://stackoverflow.com/questions/26615080/indentation-error
Score: 0.000000

Titre: How to change indentation in Visual Studio Code?
Lien: https://stackoverflow.com/questions/34174207/how-to-change-indentation-in-visual-studio-code
Score: 332.000000
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
