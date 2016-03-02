# GitNGo

## Installation

Ce package est disponible à l'adresse : `github.com/johnsudaar/gitngo/`.

Nous pouvons donc utiliser go get :
```shell
go get github.com/johnsudaar/gitngo
go install github.com/johnsudaar/gitngo
```

## Lancement

Par défaut le serveur se lance sur le port 8080. Cependant ceci peut être changé en utilisant le flag `-port=[PORT]`.

Par défaut le serveur cherche les ressources (templates, CSS, JS) dans le dossier `./ressources`. Ceci peut être changé en utilisant le flag `-ressources=[PATH]`

Par défaut le serveur envoie des requêtes anonymes à l'api GITHUB. Cependant le mode anonyme est limité et peut conduire à des plantages.
Pour utiliser le mode authentifié il suffit de mettre le token d'indentication dans la variable d'environnement `GITHUB_KEY`

## Lancement avec docker

Une image est disponible sur le docker hub. Pour lancer le service il suffit donc de lancer la commande suivante :

```shell
docker run --name="gitngo" --publish="8080:8080" --env="GITHUB_KEY=<KEY>" johnsudaar/gitngo
```

## Utilisation
Une fois lancé, le serveur donne accès à deux ressources :

### Acceuil

`Path : GET / `

Cette route correspond à la page d'accueil du site web. Cette page vous permet de choisir le langage sur lequel vous voulez filtrer.

De base cette recherche va se faire sur les 100 derniers projets mis à jour, mais en activant la recherche avancée, vous pouvez spécifier un champ de recherche et chercher, par exemple, les 100 derniers projets mise à jour par docker (en utilisant le filtre : `user:docker`).

Notes :
* Le filtre langage est sensible à la casse.
* Vous pouvez utiliser tous les filtres github détailés [sur cette page](https://help.github.com/articles/searching-repositories/).

### Recherche
`Path : GET /search`

Cette page va chercher et présenter les résultats en fonction de 4 paramètres.

* language : Doit obligatoirement être présent, permet de chosir le language sur lequel on veut filtrer.
* custom : Doit être sur 'on' si l'on veut faire une recherche avancée.
* query : la requète a executer (le champ custom doit être présent pour que ce champ soit pris en compte)
* max_routines : le serveur utilisera au maximum ce nombre de routines pour faire le calcul (le champ custom doit être présent pour que ce champ soit pris en compte)

## Fonctionnement

Ce projet est découpé en plusieurs modules :
* Le webserver
* Le gitprocessor
* Les filtres

### Module webserver

Ce module sert à récuperer les requêtes HTTP entrantes et à générer les reponses adéquates.

Il est décomposé en plusieurs parties.
Les handlers sont les fonctions qui sont appelés lors du chargement d'une page. Leur but est de récupérer les données nécessaire au chargement des pages et de charger le template adéquat.

Le moteur de template utilisé est celui fourni dans le package html/template. Une fonction render à été ajouté permettant de simplifier son utilisation.

Pour le routeur (fichier webserver.go), nous utilisons httprouter. Il y a deux types de routes. Les routes "connues" qui sont directement relièes aux Handler et le 404 qui est relié à un serveur de fichier pour servir les assets. De plus un adaptateur pour ajouter des middleware à été ajouté permettant d'utiliser alice (bien que pas forcement utile dans ce cas précis). Et un middleware permettant de logger les actions à été ajouté.

### Module GitProcessor

Ce module est en charge de la communication avec l'API github. Il fournis deux fonctions utiles :

GetRepositoryLanguages : qui permet de récupérer les langages utilisés dans un repository*
GetGithubRepositories : qui permet de récupérer les 100 derniers repositories correspondant à une certaine requête.

Pour simplifier les appels HTTP, ce module se base sur le package sling.

### Module Filter

Son but est de filtrer les résultats. Il recoit en entrée une liste de repository et un langage, il va filtrer les repository de manière a ce qu'il ne reste plus que ceux qui utilisent ce language et va isoler les données nécessaires à l'affichage.

### Front end
Tous les fichiers du front-end sont disponibles dans le dossier ressources/.
Le template CSS utilisé est Bootstrap.
Pour l'affichage des graph, c'est la librairie HighCharts qui est utilisées.
