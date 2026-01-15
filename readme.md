# Petites Victoires - Forum

Petites Victoires est une plateforme communautaire (forum) conçue pour permettre aux utilisateurs de partager leurs succès quotidiens, aussi petits soient-ils. L'objectif est de créer un espace bienveillant pour célébrer le progrès personnel et encourager la gratitude.

## Fonctionnalités

- Authentification sécurisée : Inscription et connexion des utilisateurs.

- Espace de discussion : Publication de messages ("Petites Victoires") classés par catégories.

- Interactions : Possibilité de commenter les publications pour encourager les autres membres.

- Profil Utilisateur : Gestion des informations personnelles et historique des publications.

- Administration : Outils de modération pour garantir la bienveillance sur la plateforme.

## Technique

Ce projet a été développé avec les technologies suivantes :

- Frontend : HTML5, CSS3, JavaScript

- Backend : Go

- Base de données : Sqlite3

## Installation

Pour lancer le projet localement, suivez ces étapes :

Cloner le repo :

`git clone https://github.com/JadeJLC/Petites-Victoires-forum-.git
cd Petites-Victoires-forum-
Installation des dépendances : (Exemple pour un projet PHP/Composer)`

`docker build -t forum .
docker run -p 5090:5080 -v $(pwd)/data:/data forum`

Les données d'environnement (Authentification Google, Github et Discord) ne sont pas transmises en public. Si vous voulez activer ces fonctionnalités, vous devez posséder vos propres clés d'authentification.

## Réalisation

Projet réalisé dans le cadre de la formation Zone01.
