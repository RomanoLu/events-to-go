# Events-to-Go
Ein Service zum Erstellen von Events (z.B Fußball spielen am Samstag, Geburtstagsfeier von Tim, Party am Campus, usw.). Jeder Nutzer kann sowohl Organisator als auch Teilnehmer sein. Ein Organisator erstellt ein Event. Ein Event findet an einer Location statt. Ein Organisator kann zudem seine eigenen Events updaten und löschen. Teilnehmer können einem Event zusagen. Teilnehmer können jedem Event zusagen.    

# Curls
1. Defualt addEvent
curl -H "Content-Type: application/json" -d '{"titel":"Fußball","description":"Heute Mittag Kicken beim Gymi","Beginn":"Heute Mittag","End":"Heute Abend", "maxnumberofparticipants": 12, "participants": [{}], "location" : {"name": "Gymi Pfullingen", "Postcode": "72793", "city":"Pfullingen", "adress": "Hier", "maxcapacity": 20, "latitude": 2.22, "longitude": 2.22}, "type": "OPEN", "host": {"name": "Luca", "email": "Luca@mail.de", "passwort": "1234"}}' localhost:8000/addevent
