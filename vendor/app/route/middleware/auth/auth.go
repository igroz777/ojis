package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var (
	webTemplates = template.Must(template.ParseGlob("templates/*.html"))
	hmacSecret   = []byte("serverkey")
)

type Credentials struct {
	Username    string
	AccessLevel string
	TokenId     string
}

func DissalowAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check if authenticated
		cookie, err := r.Cookie("USRCRED")
		if err != nil {
			h.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		//If authenticated, dont't allow access to page

	})
}

func DissalowAnon(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check if authenticated
		cookie, err := r.Cookie("USRCRED")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		} else {
			h.ServeHTTP(w, r)
		}

		//If authenticated, dont't allow access to page

	})
}

func BlockCSRF(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, err := r.Cookie("USRCRED")
			if err != nil {
				//
				//REVOKE COOKIES HERE
				//
				http.Redirect(w, r, "/401", http.StatusUnauthorized)
				return
			}
			verifyTKN := r.PostFormValue("VRFTKN")
			if verifyTKN == "" {
				//
				//REVOKE COOKIES HERE
				//
				http.Redirect(w, r, "/401", http.StatusUnauthorized)
				return
			}
			userCred, err := getCredentials(r)
			if err != nil {
				//
				//REVOKE COOKIES HERE
				//
				http.Redirect(w, r, "/401", http.StatusUnauthorized)
				return
			}
			if userCred.Username != verifyTKN {
				//
				//REVOKE COOKIES HERE
				//
				http.Redirect(w, r, "/401", http.StatusUnauthorized)
				return
			}
			h.ServeHTTP(w, r)

		})
}

func getCredentials(r *http.Request) (Credentials, error) {
	var err error
	tokenstring := (r.Cookie("USRCRED")).Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method:")
		}
		return hmacSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return Credentials{
			Username:    claims["username"],
			AccessLevel: claims["accesslevel"],
			TokenId:     claims["tokenid"],
		}, nil
	} else {
		return Credentials{}, err
	}

}
