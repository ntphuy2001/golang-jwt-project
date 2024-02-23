package dotenvsingleton

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Singleton interface {
	GetPort() string
	GetMongodbURL() string
	GetSecretKey() string
}

type dotenvSingleton struct {
	PORT        string
	MONGODB_URL string
	SECRET_KEY  string
	// Add other configuration fields as needed
}

func (s *dotenvSingleton) GetPort() string {
	return s.PORT
}

func (s *dotenvSingleton) GetMongodbURL() string {
	return s.MONGODB_URL
}

func (s *dotenvSingleton) GetSecretKey() string {
	return s.SECRET_KEY
}

var (
	once     sync.Once
	instance *dotenvSingleton
)

func GetDotenvInstance() Singleton {
	once.Do(func() {
		instance = loadEnv()
	})
	return instance
}

func loadEnv() *dotenvSingleton {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &dotenvSingleton{
		PORT:        os.Getenv("PORT"),
		MONGODB_URL: os.Getenv("MONGODB_URL"),
		SECRET_KEY:  os.Getenv("SECRET_KEY"),
	}
}
