package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() string {
	id := uuid.Must(uuid.NewRandom())
	return id.String()
}

func HashPassword(password string) (result string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	result = string(bytes)
	return
}

func CheckPasswordHash(password, hash string) error {
	fmt.Println(password)
	fmt.Println(hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func ValidationRace(race string) bool {
	switch race {
	case "Persian", "Maine Coon", "Siamese", "Ragdoll", "Bengal", "Sphynx", "British Shorthair", "Abyssinian", "Scottish Fold", "Birman":
		return true
	}
	return false
}

func ConvertDateIso(datetime time.Time) string {
	return time.Now().Format("2006-01-02")
}

func SliceToString(s []string) (data string) {
	data = strings.Join(s, ",")
	return
}

func StringToSlice(data string) (s []string) {
	s = strings.Split(data, ",")
	return
}
