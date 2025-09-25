Mini-CRM CLI (TP3 Go)

Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go.
Ce projet illustre l’utilisation de Cobra, Viper, GORM et SQLite pour créer une application CLI modulaire et configurable.

🚀 Fonctionnalités

Gestion de contacts (CRUD) : ajouter, lister, mettre à jour, supprimer.

CLI professionnelle avec Cobra
.

Configuration externe avec Viper
.

Persistance multi-backend :

SQLite via GORM (type: gorm)

Fichier JSON (type: json)

Mémoire (type: memory)

📦 Installation
Prérequis

Go 1.22+

GCC (ex: via MSYS2 sur Windows) pour compiler go-sqlite3

Git

Cloner le projet
git clone https://github.com/lulu73211/go_crm_tp3.git
cd go_crm_tp3

Installer les dépendances
go mod tidy

Compiler
go build -o crm.exe .

⚙️ Configuration (config.yaml)

Le fichier config.yaml définit le backend de persistance :

Exemple avec SQLite (GORM)
type: gorm
db_path: data/crm.db
json_path: data/contacts.json

Exemple avec JSON
type: json
db_path: data/crm.db
json_path: data/contacts.json

Exemple avec Mémoire
type: memory
db_path: data/crm.db
json_path: data/contacts.json

📖 Utilisation
Ajouter un contact
./crm.exe add --name "Jeanne" --email jeanne@example.com --phone 0612345678

Lister les contacts
./crm.exe list

Mettre à jour un contact
./crm.exe update --id 1 --name "Jeanne Doe" --email jeanne.doe@example.com

Supprimer un contact
./crm.exe delete --id 1

🗂️ Persistance

Avec type: gorm → les contacts sont stockés dans data/crm.db (SQLite).

Avec type: json → les contacts sont stockés dans data/contacts.json.

Avec type: memory → rien n’est sauvegardé (idéal pour les tests rapides).

✅ Critères de réussite (TP3)

 CRUD fonctionnel en CLI

 Configurable via config.yaml (Viper)

 Backend SQLite avec GORM

 Backend JSON

 Backend mémoire

 Code organisé en packages (cmd, config, internal/domain, internal/store, internal/service)

📌 Exemple d’exécution
# Ajouter un contact
./crm.exe add --name "Paul" --email paul@example.com --phone 0611222333

# Lister
./crm.exe list
ID   NAME    EMAIL              PHONE
1    Jeanne  jeanne@example.com 0612345678
2    Paul    paul@example.com   0611222333