package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	protonmail "github.com/shehackedyou/protonmail"
	//totp "github.com/pquerna/otp/totp"
)

func main() {
	manager := protonmail.New(protonmail.Config{
		HostURL:    "https://api.protonmail.ch",
		AppVersion: "web-account@4.28.2",
	})

	username := os.Getenv("PROTONMAIL_USERNAME")
	password := os.Getenv("PROTONMAIL_PASSWORD")
	//totpSecret := os.Getenv("PROTONMAIL_TOTP_SECRET")

	client, _, err := manager.NewClientWithLogin(context.Background(), username, []byte(password))
	if err != nil {
		panic(err)
	}

	//otp, err := totp.GenerateCode(totpSecret, time.Now())
	//if err != nil {
	//	panic(err)
	//}

	//err = client.Auth2FA(context.Background(), otp)
	//if err != nil {
	//	panic(err)
	//}

	salt, err := client.AuthSalt(context.Background())
	if err != nil {
		panic(err)
	}

	hashedMboxPassword, err := protonmail.HashMailboxPassword([]byte(password), salt)
	if err != nil {
		panic(err)
	}

	err = client.Unlock(context.Background(), hashedMboxPassword)
	if err != nil {
		panic(err)
	}

	addressList, err := client.GetAddresses(context.Background())
	if err != nil {
		panic(err)
	}

	mainAddressId := addressList.Main().ID
	messageCounts, err := client.CountMessages(context.Background(), mainAddressId)
	if err != nil {
		panic(err)
	}

	var allMessages []*protonmail.Message
	allMailMessageCount := messageCounts[5]
	page := 0
	for count := allMailMessageCount.Total; count > 0; count -= 100 {
		messages, _, err := client.ListMessages(context.Background(), &protonmail.MessagesFilter{
			Page:      page,
			PageSize:  100,
			AddressID: mainAddressId,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("adding (%v) messages\n", len(messages))
		allMessages = append(allMessages, messages...)
		page++
	}

	fmt.Println("how many messages?", len(allMessages))

	if len(allMessages) > 0 {
		latestMessage, err := client.GetMessage(context.Background(), allMessages[0].ID)
		if err != nil {
			panic(err)
		}

		keyRing, err := client.KeyRingForAddressID(mainAddressId)
		if err != nil {
			panic(err)
		}

		content, err := latestMessage.Decrypt(keyRing)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(content))

		err = client.AuthDelete(context.Background())
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("could not find any messages!")
	}
}

// TODO: Send mail is something like this; I think the ideal solution atm would
// be to not run a webpage but instead just resend the code to a persons email
//func sendMail() {
//	c, privateKeys, err := auth.NewManager(newClient).Auth(username, bridgePassword)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	u, err := c.GetCurrentUser()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	addrs, err := c.ListAddresses()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = smtpbackend.SendMail(c, u, privateKeys, addrs, rcpt, os.Stdin)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func printStruct(s interface{}) {
	structJson, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(structJson))
}
