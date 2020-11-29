package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strings"
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
	file, err := os.OpenFile("ladon.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		exit("Failed to open the log file, %v", err)
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(file)
	entry := log.WithField("user", os.Args[1])
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
		suffix := fmt.Sprintf("|%.6f", math.Sin(1/(7*rand.Float64()+0.001)))
		if sha256.Sum256(append(password, suffix...)) ==
			sha256.Sum256([]byte(user.Password+suffix)) {
			isVerified = true
		}
	}
	if !isVerified {
		entry.Info("Identification failed")
		exit("\nAccess denied\n")
	} else {
		fmt.Println()
	}
	entry.Info("Identification succeeded")
	reader := bufio.NewReader(os.Stdin)
	authenticate(reader, user, entry)
	for range time.Tick(time.Second * 30) {
		authenticate(reader, user, entry)
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

func authenticate(reader *bufio.Reader, user User, entry log.FieldLogger) {
	likesCakes := ask(reader, "Do you like cakes (y/n): ")
	washesHands := ask(reader, "Do you wash hands (y/n): ")
	drivesCar := ask(reader, "Do you drive a car (y/n): ")
	eatsJunkFood := ask(reader, "Do you eat junk food (y/n): ")
	if likesCakes != user.LikesCakes ||
		washesHands != user.WashesHands ||
		drivesCar != user.DrivesCar ||
		eatsJunkFood != user.EatsJunkFood {
		entry.Info("Authentication failed")
		exit("Access denied\n")
	} else {
		entry.Info("Authentication succeeded")
	}
}

func ask(reader *bufio.Reader, question string) bool {
	fmt.Print(question)
	line, err := reader.ReadString('\n')
	if err != nil {
		exit("Failed to read the line, %v", err)
	}
	switch strings.ToLower(line[:len(line)-1]) {
	case "y":
		return true
	case "n":
		return false
	default:
		exit("Got malformed input %q\n", line)
		return false
	}
}
