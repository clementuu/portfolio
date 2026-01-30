# Portfolio

Ce projet est une application web de portfolio full-stack, composée d'un backend en Go et d'un frontend en Svelte.

## Stack-technique

| Catégorie | Technologie |
| --- | --- |
| **Backend** | Go (Golang) |
| **Frontend** | Svelte |
| **Tests** | Go Testing (backend), Vitest (frontend) |
| **Conteneurisation** | Docker, Docker Compose |
| **CI/CD** | GitHub Actions |

## Structure du projet

Le projet est divisé en deux répertoires principaux :

-   `back/` : Contient l'API backend en Go.
-   `ihm/` : Contient l'application frontend Svelte (IHM - Interface Homme-Machine).

## Pages Principales (Frontend)

-   **Accueil** : Affiche le parcours académique (`Formations`) et `Expériences` professionnelles.
-   **Compétences** : Liste des compétences techniques et non techniques.
-   **Projets** : Présentation des projets réalisés.
-   **Contact** : Informations de contact.

## Endpoints de l'API (Backend)

L'API Go expose les endpoints suivants pour fournir les données au frontend :

| Méthode | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/competences` | Récupère la liste de toutes les compétences. |
| `GET` | `/competence/{id}` | Récupère une compétence spécifique par son ID. |
| `GET` | `/projets` | Récupère la liste de tous les projets. |
| `GET` | `/projets/names` | Récupère uniquement les noms de tous les projets. |
| `GET` | `/projet/{id}` | Récupère un projet spécifique par son ID. |
| `GET` | `/formations` | Récupère la liste des formations académiques. |
| `GET` | `/experiences` | Récupère la liste des expériences professionnelles. |

## Prérequis

Avant de commencer, assurez-vous d'avoir installé les outils suivants :

-   [Go](https://golang.org/dl/) (version 1.22+ recommandée)
-   [Node.js](https://nodejs.org/en/) (version 20+ recommandée)
-   [Docker](https://www.docker.com/get-started)
-   [Docker Compose](https://docs.docker.com/compose/install/)
-   `make`

## Démarrage rapide avec Docker

1.  **Buildez le conteneur :**
    ```sh
    make build
    ```

2.  **Démarrez le conteneur :**
    ```sh
    make start
    ```
    L'application sera accessible sur [`http://localhost:8080`](http://localhost:8080).

3.  **Arrêtez le conteneur :**
    ```sh
    make stop
    ```

## Développement local

Pour un développement complet, vous devez lancer le backend et le frontend séparément.

### Lancer le serveur Backend

```sh
make srv
```
Le serveur Go démarrera sur le port 8080.

### Lancer le client Frontend

```sh
make web
```
Ceci installera les dépendances `npm` et lancera le serveur de développement.

## Tests

Le `Makefile` fournit des cibles pour exécuter les tests du backend et du frontend.

-   **Exécuter les tests du backend :**
    ```sh
    make test-back
    ```

-   **Exécuter les tests du frontend avec couverture de code :**
    ```sh
    make test-ihm
    ```

## Intégration et Déploiement Continus (CI/CD)

Ce projet utilise GitHub Actions pour l'intégration continue. Le workflow, défini dans `.github/workflows/ci.yml`, se déclenche à chaque `push` ou `pull request` sur la branche `main`.

Le pipeline exécute les tâches suivantes :
1.  **Build** : Construit l'image Docker de l'application.
2.  **Test** : Exécute les tests unitaires et de couverture pour le backend Go et le frontend Svelte.
