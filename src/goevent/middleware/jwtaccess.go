package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RomanoLu/events-to-go/src/goevent/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var SECRET []byte
var api_key = "1234"

func CreateJWT(secret string) (string, error) {
	//Hier muss die UserID übergeben werden, welche künftig das Secret ist

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	SECRET = []byte(secret)
	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]string)
		users, err := service.GetAllUsers()
		if err != nil {
			return
		}
		for _, u := range users {
			m[u.Email] = u.Passwort
		}
		var secret string
		for key, value := range m {
			if r.Header["Access"][0] == key {
				secret = value
			}
		}

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return []byte(secret), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}
			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func ValidateHostRole(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var secret string
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 0)
		if err != nil {
			log.Errorf("Can't get ID from request: %v", err)
			return
		}
		event, err := service.GetEventById(uint(id))
		if r.Header["Access"][0] == event.Host.Email {
			secret = event.Host.Passwort
			if r.Header["Token"] != nil {
				token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
					_, ok := t.Method.(*jwt.SigningMethodHMAC)
					if !ok {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("not authorized"))
					}
					return []byte(secret), nil
				})

				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized: " + err.Error()))
				}
				if token.Valid {
					next(w, r)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized"))
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func GetJwt(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	users, err := service.GetAllUsers()
	if err != nil {
		return
	}
	for _, u := range users {
		m[u.Email] = u.Passwort
	}

	//prüfen ob accesss == wie irgeneine email ist,
	//und wenn ja, dann createJTW mit der userID als Parameter aufrunfen
	if r.Header["Access"] != nil {
		for key, value := range m {
			if r.Header["Access"][0] == key {
				token, err := CreateJWT(value)
				if err != nil {
					return
				}
				fmt.Fprint(w, token)
			}
		}

	}
}
