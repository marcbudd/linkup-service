# linkup-service

Unser eigenes soziales Netzwerk!

## Anforderungen

Um das Projekt auszuführen, sind folgende Voraussetzungen erforderlich:

- Go (Version 1.18): [Installationsanleitung](https://golang.org/doc/install)
- PostgreSQL-Datenbank: [Installationsanleitung](https://www.postgresql.org/download)

## Konfiguration

1. Erstelle eine `.env`-Datei im Projektverzeichnis mit den folgenden Inhalten:

```shell
PORT=
JWT_SECRET=

DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
DB_SSLMODE=

EMAIL_HOST=
EMAIL_PORT=
EMAIL_ADDRESS=
EMAIL_PASSWORD=

PROXY_HOST=

OPENAI_API_KEY=
OPENAI_API_URL=
```

Stelle sicher, dass du eine PostgreSQL-Datenbank und einen Email-Server mit den entsprechenden Credentials eingerichtet hast und die Werte in der `.env`-Datei aktualisierst, um eine erfolgreiche Verbindung herzustellen.

## Installation

1. Klone das Git-Repository:

```shell
git clone https://github.com/marcbudd/linkup-service.git
cd linkup-service
```

2. Installiere die Abhängigkeiten:

```shell
go mod download
```

## Ausführung

```shell
go run main.go
```