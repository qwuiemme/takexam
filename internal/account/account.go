package account

import (
	"crypto/sha256"
	"encoding/base64"
	"log"

	"github.com/hnnngn/take-exam/pkg/client"
)

type Account struct {
	Login      string
	Password   string
	AvatarLink string
}

func GetSHA256Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (a *Account) InsertIntoDatabase() {
	conn := client.Connect()
	defer conn.Close()

	res, err := conn.Query("INSERT INTO `accounts` (`Login`, `Password`, `AvatarLink`) VALUES ('" + a.Login + "', '" + a.Password + "', '" + a.AvatarLink + "')")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()
}

func GetFromDatabase(login string) (accounts []Account) {
	conn := client.Connect()
	defer conn.Close()

	res, err := conn.Query("SELECT `Login`, `Password`, `AvatarLink` FROM `accounts` WHERE Login = '" + login + "'")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	for res.Next() {
		var acc Account
		err = res.Scan(&acc.Login, &acc.Password, &acc.AvatarLink)

		if err != nil {
			log.Fatal(err)
		}

		accounts = append(accounts, acc)
	}

	return
}
