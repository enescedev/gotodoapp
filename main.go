package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/enescedev/gotodo/config"
	"github.com/enescedev/gotodo/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

var cache *redis.Client
var router = gin.Default()

// Bu middleware, bir HTTP isteği aldığında JWT'yi doğrular ve
// isteği yalnızca geçerli bir JWT ile birlikte gelen bir istek olarak yönlendirir.
func jwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// JWT'yi HTTP isteğinden ayıkla
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "JWT bulunamadı", http.StatusUnauthorized)
			return
		}
		// JWT'yi doğrula
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Bu işlev, JWT'nin imzasını doğrular ve
			// JWT'nin içeriğine erişim için gereken anahtarı döndürür.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("geçersiz imza yöntemi")
			}
			return []byte("gizli-anahtar"), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Eğer JWT geçerli ise, isteği yönlendirin
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "geçersiz JWT", http.StatusUnauthorized)
		}
	})
}

// Bu işlev, bir kullanıcının giriş bilgilerini kullanarak bir JWT oluşturur.
func createJWT(w http.ResponseWriter, r *http.Request) {
	// Kullanıcının giriş bilgilerini kullanarak bir JWT oluşturun
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  123,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, err := token.SignedString([]byte("gizli-anahtar"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Oluşturulan JWT'yi kullanıcıya gönderin
	w.Write([]byte(tokenStr))
}

func main() {
	var jwt_secret = os.Getenv("jwt_secret")

	if jwt_secret == "" {
		log.Fatal("jwt anahtarı yok")
		return
	}

	// init cache connection
	cache = config.CacheConnection("localhost:6379", "")
	log.Println("Redis e bağlandı")
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	config.Connect("postgres://bbpguser:bbpguserpassword@localhost:5432/postgres")
	routes.ToDoRoute(router)
	log.Println("Sunucu ayağa kalktı")
	log.Fatal(router.Run(":8080"))

	//port
	router.Run(":8080")

	http.HandleFunc("/create-jwt", createJWT)

	// JWT middleware'ini kullanarak güvenli bir API işlevi oluşturun
	http.HandleFunc("/secure-api", jwtMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("güvenli API işlevi"))
	}))

	http.ListenAndServe(":8000", nil)

}
