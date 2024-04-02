package login

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2/log"
	cmn "myIndia/common"
	"net/http"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w *http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	up := fmt.Sprintf("Username : %s, Password: %s", username, password)

	log.Infof("Login value is ", up)

	if username == "Niranjan" && password == "Niran@123" {
		expireTime := time.Now().Add(5 * time.Minute)
		log.Infof("Expire Time: ", expireTime)

		claims := &Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expireTime.Unix(),
			},
		}
		log.Infof("Claims value is ", claims)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		log.Infof("Token with claims is ", token)

		jwtKey := cmn.GenerateSecretKey(12)

		log.Infof("JWT Key is ", jwtKey)

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			log.Infof("Error while getting signed string : ", err)
		}

		log.Infof("Token string is ", tokenString)
	}
}
