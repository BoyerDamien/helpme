# helpme

Un petit outil en ligne de commande écrit en Go facilitant la recherche de solutions liées à une erreur.
Ce programme utilise l'API de [Stackexchange](https://api.stackexchange.com/docs).

## Examples
```bash
$ helpme segfault 2

Titre: Why does this code segfault on 64-bit architecture but work fine on 32-bit?
Lien: https://stackoverflow.com/questions/7545365/why-does-this-code-segfault-on-64-bit-architecture-but-work-fine-on-32-bit
Tags: c, pointers, segmentation-fault, 
Score: 110

Titre: [UN-ANSWERED] What is a bus error?
Lien: https://stackoverflow.com/questions/212466/what-is-a-bus-error
Tags: c, unix, segmentation-fault, bus-error, 
Score: 0
```
Cette commande affichera sur la sortie standard les 2 pages de stackexchange les plus pertinentes par rapport à l'erreur entrée.

## Syntaxe:
```bash
$ helpme <erreur> <nombre de pages> [-s]
```
L'option -s permet de trier les resultats par score

## Installation

L'installation nécessite la présence de Go sur votre système -> [Installation de go](https://golang.org/dl/) 
```bash
$ go build github.com/BoyerDamien/helpme && sudo mv helpme /usr/bin
```
