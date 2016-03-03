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
Pour utiliser le mode authentifié il suffit de mettre le token d'identification dans la variable d'environnement `GITHUB_KEY`

## Lancement avec docker

Une image est disponible sur le docker hub. Pour lancer le service il suffit donc de lancer la commande suivante :

```shell
docker run --name="gitngo" --rm --publish="8080:8080" --env="GITHUB_KEY=<KEY>" johnsudaar/gitngo
```

## Utilisation
Une fois lancé, le serveur donne accès à deux ressources :

### Accueil

`Path : GET / `

Cette route correspond à la page d'accueil du site web. Cette page vous permet de choisir le langage sur lequel vous voulez filtrer.

De base cette recherche va se faire sur les 100 derniers projets mis à jour, mais en activant la recherche avancée, vous pouvez spécifier un champ de recherche et chercher, par exemple, les 100 derniers projets mise à jour par docker (en utilisant le filtre : `user:docker`).

Note :
* Vous pouvez utiliser tous les filtres github détaillés [sur cette page](https://help.github.com/articles/searching-repositories/).

### Recherche
`Path : GET /search`

Cette page va chercher et présenter les résultats en fonction de 4 paramètres.

* `language` : Doit obligatoirement être présent, permet de choisir le langage sur lequel on veut filtrer.
* `custom` : Doit être sur 'on' si l'on veut faire une recherche avancée.
* `query` : la requête à exécuter (le champ `custom` doit être présent pour que ce champ soit pris en compte)
* `max_routines` : le serveur utilisera au maximum ce nombre de routines pour faire le calcul (le champ `custom` doit être présent pour que ce champ soit pris en compte)
* `no_lines` : Doit être sur 'on' pour que le serveur ne calcule pas le nombre de lignes écrite. (Accélère grandement la recherche) (le champ `custom` doit être présent pour que ce champ soit pris en compte).

## Fonctionnement

Ce projet est découpé en plusieurs modules :
* Le webserver
* Le gitprocessor
* Les filtres

### Module webserver

Ce module sert à récupérer les requêtes HTTP entrantes et à générer les reponses adéquates.

Il est décomposé en plusieurs parties.
Les handlers sont les fonctions qui sont appelés lors du chargement d'une page. Leur but est de récupérer les données nécessaires au chargement des pages et de charger le template adéquat.

Le moteur de template utilisé est celui fourni dans le package `html/template`. Une fonction render à été ajouté permettant de simplifier son utilisation.

Pour le routeur (fichier webserver.go), nous utilisons httprouter. Il y a deux types de routes. Les routes "connues" qui sont directement relièes aux Handler et le 404 qui est relié à un serveur de fichier pour servir les assets. De plus un adaptateur pour ajouter des middleware à été ajouté permettant d'utiliser alice (bien que pas forcement utile dans ce cas précis). Et un middleware permettant de logger les actions à été ajouté.

### Module GitProcessor

Ce module est en charge de la communication avec l'API github. Il fournit trois fonctions utiles :

* `GetRepositoryLanguages` : qui permet de récupérer les langages utilisés dans un repository
* `GetGithubRepositories` : qui permet de récupérer les 100 derniers repositories correspondant à une certaine requête.
* `GetRepositoryLines` : qui permet de récupérer le nombre de lignes écrites dans un repository.

Note :
GetRepositoryLines n'est pas fiable. Elle est basée sur la statistique `code_frequency` qui doit nous donner les lignes ajoutées et supprimées pour chaque semaine sur la branche par défaut du repository. Cependant sur certains repository le nombre de lignes ajoutées est supérieur au nombre de lignes supprimées. (Ex : zorluhan/spogram)

Pour simplifier les appels HTTP, ce module se base sur le package sling.

### Module Filter

Son but est de filtrer les résultats. Il recoit en entrée une liste de repository et un langage, il va filtrer les repository de manière a ce qu'il ne reste plus que ceux qui utilisent ce langage et va isoler les données nécessaires à l'affichage.

### Front end
Tous les fichiers du front-end sont disponibles dans le dossier ressources/.
Le template CSS utilisé est Bootstrap.
Pour l'affichage des graph, c'est la librairie HighCharts qui est utilisée.

### Calcul du nombre de lignes (Expérimental)

Vu que github ne nous donne pas le nombre de lignes écrites pour un projet (et encore moins pour un projet par langage),
le serveur essaie d'estimer le nombre de lignes d'un repository en utilisant les statistiques (additions/deletions) par semaine sur la branche par défaut.
Pour estimer le nombre de lignes par langage, nous calculons le ratio : BytesDansCeLanguage / BytesTotal que nous appliquons au nombre total de lignes.

A cause de l'approximation faite par la règle de 3 et de l'inexactitude des données fournis par l'API github. Le champ lignes peut parfois être assez loin du compte.

Vu que les statistiques fournis par github ne sont pas toujours disponible en cache, il faut parfois lui laisser le temps de calculer les résultats. Dans notre cas nous lui laissons 1s pour les calculer. (20 essais espacés de 50ms).

Cette étape étant extrèmement lente, le calcul du nombre de lignes peut être désactive en utilisant le paramêtre `no_lines`. (Aussi disponible depuis la recherche avancée).
