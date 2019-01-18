package jwt

import (
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	// 盐
	jwtSecret = []byte("baiyang")
	// 过期时间
	expireSecond = 1 * 60 * 60
	// 发行人
	issur = "baiyang"

	// ErrorExpired 过期
	ErrorExpired = errors.New("Token is expired")
	// ErrorNotValidYet 未生效
	ErrorNotValidYet = errors.New("Token not active yet")
	// ErrorMalformed 畸形Token
	ErrorMalformed = errors.New("That's not even a token")
	// ErrorInvalid 其他错误
	ErrorInvalid = errors.New("Couldn't handle this token")
)

// CustomClaims 定义一个Claims结构体,可以添加自定义属性
type CustomClaims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(userID int) (string, error) {
	// 新建一个Claims
	claims := CustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expireSecond) * time.Second).Unix(), // 设置超时时间
			Issuer:    issur,                                                            // 设置发行人
		},
	}
	// 生成Token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签署
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrorMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrorExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrorNotValidYet
			} else {
				return nil, ErrorInvalid
			}
		}
	}
	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, ErrorInvalid
}

// Auth 中间件，检查token
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// parseToken 解析token包含的信息
		claims, err := ParseToken(token)
		if err != nil {
			if err == ErrorExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": -1,
					"msg":  "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("UserID", claims.UserID)
		// c.Next()
	}
}
