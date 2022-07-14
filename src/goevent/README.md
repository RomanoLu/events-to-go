# Events-to-Go
Ein Service zum Erstellen von Events 
(z.B Fußball spielen am Samstag, Geburtstagsfeier von Tim, Party am Campus, usw.). 

Jeder Nutzer kann sowohl Veranstalter als auch Teilnehmer sein. Ein Veranstalter erstellt ein Event. Ein Event findet an einer Location statt. Ein Veranstalter kann zudem seine eigenen Events updaten und löschen. Teilnehmer können einem Event zusagen. Teilnehmer können jedem Event zusagen welches als OPEN gekennzeichnet ist. Der Veranstalter eines Events kann, wenn es INVETATION_ONLY ist, Inventation erstellen und an User schicken. Ein User kann seine Einladungen danach ansehen und akzeptieren. 


Im Outlook speichern

# Google Calender API
go get -u google.golang.org/api/calendar/v3
go get -u golang.org/x/oauth2/google
