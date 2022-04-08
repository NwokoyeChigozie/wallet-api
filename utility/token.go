package utility

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NwokoyeChigozie/quik_task/internal/config"
	"github.com/golang-jwt/jwt"
)

func CreateToken(userID int) (string, time.Time, error) {

	var (
		config  = config.GetConfig()
		UnixExp = time.Now().AddDate(0, 0, 7).Unix()
		exp     = time.Now().AddDate(0, 0, 7)
	)

	//create token
	userid := strconv.Itoa(userID)
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["authorised"] = true
	atClaims["exp"] = UnixExp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	accessToken, err := token.SignedString([]byte(config.Server.Secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return accessToken, exp, nil
}

func TokenValid(bearerToken string) (*jwt.Token, error) {
	token, err := verifyToken(bearerToken)
	if err != nil {
		if token != nil {
			return token, err
		}
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}
	return token, nil
}

//verifyToken verify token
func verifyToken(tokenString string) (*jwt.Token, error) {
	config := config.GetConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Server.Secret), nil
	})
	if err != nil {
		return token, fmt.Errorf("Unauthorized")
	}
	return token, nil
}
