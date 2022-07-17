# Events-to-Go

## Business Background
Ein Service zum Erstellen von Events 
(z.B Fußball spielen am Samstag, Geburtstagsfeier von Tim, Party am Campus, usw.)

Nutzer können entweder als Host eine Veranstaltung anlegen und pflegen oder als Participants nach Veranstaltungen suchen und an ihnen teilnehmen.
Jedes Event beinhaltet Informationen über die Location, die Teilnehmer und den Begin/das Ende der Veranstaltung. 
Participants können zusätzlich zu den Veranstaltungen eingeladen werden und ihnen zusagen.
Bei Veranstaltungen welche als OPEN gekennzeichnet sind können Participents zudem auch ohne die Zustimmung des Hosts teilnehmen.


### Google Calender API
```sh
go get -u google.golang.org/api/calendar/v3
```

```sh
go get -u golang.org/x/oauth2/google
```
Anleitung:
https://developers.google.com/calendar/api/v3/reference/events

Anleitung für einen quickstart
https://developers.google.com/calendar/api/quickstart/go

### JSON WEB TOKEN (JWT)
"github.com/dgrijalva/jwt-go"

## Deployment
1.  Clone the repo
```sh
git clone https://github.com/RomanoLu/events-to-go.git
```
2. Build and run the Docker container
```sh
docker compose up --build
```

### Request
You find the exported Postman workspace in die Repo above 
