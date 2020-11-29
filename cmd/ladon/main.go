package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		exit("Just 1 arg required, but %d provided\n", len(os.Args)-1)
	}
	bytes, err := ioutil.ReadFile("examples/users.json")
	if err != nil {
		exit("Failed to read the registry, %v\n", err)
	}
	var users Users
	if err := json.Unmarshal(bytes, &users); err != nil {
		exit("Failed to parse the registry, %v\n", err)
	}
	user, ok := users[os.Args[1]]
	if !ok {
		exit("No user with name %q found\n", os.Args[1])
	}
	isVerified := false
	for i := 0; i < 3 && !isVerified; i++ {
		if i == 0 {
			fmt.Print("Password: ")
		} else {
			fmt.Print("\nIncorrect value, try again: ")
		}
		password, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			exit("\nFailed to read the password, %v", err)
		}
		if sha256.Sum256(password) == sha256.Sum256([]byte(user.Password)) {
			isVerified = true
		}
	}
	if !isVerified {
		exit("\nAccess denied\n")
	}
	file, err := os.OpenFile("ladon.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		exit("Failed to open the log file, %v", err)
	} else {
		fmt.Println()
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(file)
	log.WithField("user", os.Args[1]).Info("User signed in")
	for range time.Tick(5 * time.Second) {
		fmt.Println("Oh, fuck")
	}
}

func exit(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

type Users map[string]User

func (users *Users) UnmarshalJSON(bytes []byte) error {
	var customers map[string]User
	if err := json.Unmarshal(bytes, &customers); err != nil {
		return err
	}
	for name, user := range customers {
		customers[name] = User{
			name,
			user.Password,
			user.LikesCakes,
			user.WashesHands,
			user.DrivesCar,
			user.EatsJunkFood,
		}
	}
	*users = customers
	return nil
}

type User struct {
	Name         string `json:"-"`
	Password     string `json:"password"`
	LikesCakes   bool   `json:"likes_cakes"`
	WashesHands  bool   `json:"washes_hands"`
	DrivesCar    bool   `json:"drives_car"`
	EatsJunkFood bool   `json:"eats_junk_food"`
}

func (user User) Exists() bool {
	return user.Name != ""
}
