# helpme

Un petit outil en ligne de commande écrit en Go facilitant la recherche de solutions liées à une erreur.
Ce programme utilise l'API de [Stackexchange](https://api.stackexchange.com/docs).

## Exemples
```bash
$ helpme -s -n=5 "unindent does not match any outer indentation level" 

Titre: IndentationError: unindent does not match any outer indentation level
Lien: https://stackoverflow.com/questions/492387/indentationerror-unindent-does-not-match-any-outer-indentation-level
Score: 575.00

Titre: How does IPython's magic paste work?
Lien: https://stackoverflow.com/questions/10886946/how-does-ipythons-magic-paste-work
Score: 92.00

Titre: I'm getting an IndentationError. How do I fix it?
Lien: https://stackoverflow.com/questions/45621722/im-getting-an-indentationerror-how-do-i-fix-it
Score: 39.00

Titre: PyCharm shows pep8 expected 2 blank lines, found 1;
Lien: https://stackoverflow.com/questions/40275866/pycharm-shows-pep8-expected-2-blank-lines-found-1
Score: 31.00

Titre: (python) docstring is causing indentation error
Lien: https://stackoverflow.com/questions/2243009/python-docstring-is-causing-indentation-error
Score: 12.00
```
Cette commande affichera sur la sortie standard les 5 pages de stackexchange correspondant le plus à la recherche.

La commande en détail:

-s : Ordonne les éléments par score

-n=5 : Selectionne les 5 éléments les plus pertinents

## Syntaxe
```bash
$ helpme <flags> <contenu de la recherche>
```

Une option "help" est disponible:
```bash
$ helpme -h
Usage of helpme -> helpme <flags> <research content>

Flags   :
        -n : Number of elements to display  (-n=20)
        -h : Help
        -s : Sort elements by score
        -t : Filter post not tagged <tag>   (-t=c++)
```
PS: Les flags sont facultatifs

## Installation

L'installation nécessite la présence de Go sur votre système -> [Installation de go](https://golang.org/dl/) 
```bash
$ go install github.com/BoyerDamien/helpme
```
