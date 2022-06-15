package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

//jwt service
type JWTService interface {
	GenerateToken(user entity.Users) string
	ValidateToken(token string) (*jwt.Token, error)
	ParseToken(encodedToken string) (*jwt.Token, error)
	DeleteToken(token string) error
	CheckToken(token string) error
}
type authCustomClaims struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

var whiteListTokens = make([]string, 0, 10000)

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issuer:    "teenager",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_ACCESS_SECRET_KEY")
	if secret == "" {
		secret = "your secret api key"
	}
	return secret
}

func (service *jwtServices) GenerateToken(user entity.Users) string {
	claims := &authCustomClaims{
		Id:   user.Id,
		Name: user.Name,
		Role: strconv.Itoa(user.Role),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    service.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	tokenstring, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}

	whiteListTokens = append(whiteListTokens, tokenstring)

	return tokenstring
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

}

func (service *jwtServices) ParseToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func (service *jwtServices) CheckToken(token string) error {
	for _, v := range whiteListTokens {
		if v == token {
			return nil
		}
	}
	return fmt.Errorf("token not found")
}

func (service *jwtServices) DeleteToken(token string) error {
	whiteListTokens = remove(whiteListTokens, token)
	return nil
}

func remove(slice []string, token string) []string {
	for i, t := range slice {
		if t == token {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}
