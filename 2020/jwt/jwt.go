package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func NewToken(hmacSampleSecret string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  "plainchant",             //用户
		ExpiresAt: time.Now().Unix() + 1000, //到期时间
		Id:        "123",                    //jwt标识
		IssuedAt:  time.Now().Unix(),        //发布时间
		Issuer:    "pct",                    //发行人
		NotBefore: time.Now().Unix(),        //在此之前不可用
		Subject:   "test",                   //主题
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(tokenString, err)
	return tokenString
}

func Verify(hmacSampleSecret, tokenString string) error {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})
	if err != nil {
		return fmt.Errorf("jwt parse:%v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
	return nil
}
