package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Username             string `json:"username"`
	Password             string `json:"password"` //md5加密以后的密码
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

// 获取秘钥
func getMySecret() (secret []byte) {
	defaultJwtSecret := "9n95sPw8PRUUVCpQ0q1M"
	secretnow, err := g.Cfg().Get(gctx.New(), "JwtSecret.Secret")

	if err != nil {
		glog.New().Error(gctx.New(), "获取秘钥失败")
		return []byte(defaultJwtSecret)
	}
	if secretnow.String() == "" || len(strings.TrimSpace(secretnow.String())) == 0 {
		return []byte(defaultJwtSecret)
	}

	glog.New().Info(gctx.New(), "正在使用秘钥", secretnow.String())
	return secretnow.Bytes()
}

// @生成token
// TODO: 从配置文件中获取过期时间
func MakeToken(Passport, Password string) (tokenString string, err error) {
	claim := MyClaims{
		Password: Password,
		Username: Passport,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(getMySecret())
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return getMySecret(), nil // 这是我的secret
	}
}

// 解析token
func ParseToken(tokens string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokens, &MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
