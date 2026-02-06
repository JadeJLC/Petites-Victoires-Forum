# Petites Victoires - Forum

<div align="center">

**Plateforme communautaire bienveillante pour partager et célébrer les petites victoires du quotidien.**

</div>

## Aperçu

Petites Victoires est un forum web programmé en Go. L'objectif est de créer un espace zen et positif où les utilisateurs peuvent partager leurs succès quotidiens, aussi petits soient-ils, et encourager les autres membres de la communauté.
C'est un projet d'études en groupe développé dans le cadre de la formation à Zone01, visant à créer une plateforme complète avec authentification, gestion de base de données et interactions sociales.

## Fonctionnalités

- **Authentification sécurisée** : Système d'inscription et de connexion avec plusieurs options (email, Google, GitHub, Discord).
- **Publications de victoires** : Partage de sujets classés par catégories pour célébrer les succès quotidiens.
- **Interactions communautaires** : Système de réponses par message pour encourager et soutenir les autres membres.
- **Profil utilisateur** : Gestion des informations personnelles et historique des publications et réactions.
- **Système de modération** : Outils d'administration et de modération selon le statut de l'utilisateur.
- **Gestion de sessions** : Maintien des connexions utilisateurs de manière sécurisée.
- **Notifications** : Notifications en cas de réponse ou de réaction à un sujet suivi par l'utilisateur.

## Technologies utilisées

**Language & Framework:**

[![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)](https://www.sqlite.org/)
[![HTML](https://img.shields.io/badge/HTML-E34F26?style=for-the-badge&logo=html5&logoColor=white)](https://developer.mozilla.org/fr/docs/Web/HTML)
[![CSS](https://img.shields.io/badge/CSS-1572B6?style=for-the-badge&logo=css3&logoColor=white)](https://developer.mozilla.org/fr/docs/Web/CSS)
[![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)](https://developer.mozilla.org/fr/docs/Web/JavaScript)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

## Utilisation

### Prérequis

Docker doit être installé sur votre ordinateur pour l'installation simplifiée, ou Go et SQLite pour une installation manuelle.

- **Docker**: Téléchargement : [docker.com](https://www.docker.com/).
- **Go** (installation manuelle) : Version 1.21 ou plus. Téléchargement : [golang.org](https://golang.org/dl/).

### Installation avec Docker (recommandée)

1.  **Cloner le repo**
```bash
    git clone https://github.com/JadeJLC/Petites-Victoires-forum-.git
    cd Petites-Victoires-forum-
```

2.  **Construire et lancer le container**
```bash
    docker build -t forum .
    docker run -p 5090:5080 -v $(pwd)/data:/data forum
```

3.  **Accéder à l'interface**
    
    Ouvrez votre navigateur et accédez à :
```
    http://localhost:5090
```

### Installation manuelle

1.  **Cloner le repo**
```bash
    git clone https://github.com/JadeJLC/Petites-Victoires-forum-.git
    cd Petites-Victoires-forum-
```

2.  **Lancer l'application**
```bash
    go run .
```
    
    ou
```bash
    go run main.go
```

3.  **Accéder à l'interface**
    
    Ouvrez votre navigateur et accédez à :
```
    http://localhost:5080
```

### **Configuration de l'authentification OAuth**

**Note** : Les clés d'authentification pour Google, GitHub et Discord ne sont pas fournies publiquement. Pour activer ces fonctionnalités :

1. Créez vos propres applications OAuth sur les plateformes respectives
2. Récupérez vos clés et secrets d'API
3. Configurez les variables d'environnement appropriées

### **Utilisation du forum**

1. **Inscription/Connexion** : Créez un compte ou connectez-vous via email ou OAuth
2. **Partage de victoires** : Créez un sujet en choisissant une catégorie
3. **Interaction** : Répondez aux messages des autres pour les encourager
4. **Profil** : Gérez vos informations et consultez votre historique

## Structure du projet
```
project-root/
├── builddb/        # Scripts de construction de la base de données
├── data/           # Fichiers de données (base SQLite, uploads)
├── doc/            # Documentation du projet
├── handlers/       # Gestionnaires HTTP pour les différentes routes
├── middleware/     # Middlewares (authentification, logging, CORS)
├── models/         # Structures de données et modèles
├── routes/         # Configuration des routes de l'application
├── sessions/       # Gestion des sessions utilisateurs
├── static/         # Fichiers statiques (CSS, JavaScript, images)
├── templates/      # Templates HTML pour l'interface web
├── test/           # Tests unitaires et d'intégration
├── utils/          # Fonctions utilitaires et helpers
├── Dockerfile      # Configuration Docker
├── compose.yaml    # Docker Compose configuration
├── main.go         # Fichier principal de l'application
├── go.mod          # Go module
└── readme.md       # Ce fichier
```

## Development

### Architecture

- **Backend** : Go avec architecture MVC (Models-Views-Controllers)
- **Base de données** : SQLite3 pour le stockage persistant
- **Frontend** : HTML5, CSS3 et JavaScript vanilla
- **Authentification** : Sessions sécurisées + OAuth2 (Google, GitHub, Discord)
- **Serveur** : Serveur HTTP Go natif sur le port 5080

### Tests

Le projet inclut des tests unitaires et d'intégration dans le dossier `test/`.
```bash
go test ./test/...
```

## Sécurité

- Hashage sécurisé des mots de passe
- Protection CSRF
- Validation des entrées utilisateurs
- Gestion sécurisée des sessions
- Authentification OAuth2

## Apprentissages clés

Ce projet permet de développer des compétences dans :

- Création d'une application web complète
- Gestion de base de données SQLite
- Authentification et autorisation (session-based + OAuth)
- Architecture MVC en Go
- Déploiement avec Docker
- Tests unitaires et d'intégration
- Gestion de communauté et modération

## Autres informations

- Ce projet est un projet de groupe développé dans le cadre de ma formation à Zone01.
- L'objectif était de créer un forum complet avec toutes les fonctionnalités essentielles : authentification, CRUD, interactions sociales et modération.
- Projet collaboratif avec 4 contributeurs.


<div align="center">

Par [JadeJLC](https://github.com/JadeJLC), [Mathis Pain](https://github.com/Mathis-Pain), [Valentine Ladjyn](https://github.com/vladjyn), [Clara Hiesse](https://github.com/clarahiesse)

</div>
