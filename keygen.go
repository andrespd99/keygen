package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andrespd99/quoota-api/internal/signer"
	"github.com/atotto/clipboard"
	"github.com/golang-jwt/jwt/v4"
)

const (
	defaultSubject = "Unknown subject"
)

func main() {

	subj := flag.String("s", defaultSubject, fmt.Sprintf("Sets claim subject. Defaults to \"%s\"", defaultSubject))

	// flag.Bool("cloud-key", false, "Whether the token will be used for a cloud service. If true, the key will have ")

	flag.Parse()

	k := flag.Arg(0)
	if strings.TrimSpace(k) == "" {
		fmt.Println("The signing secret key must be passed as argument")
		os.Exit(1)
	}

	fmt.Println("** GENERATING NEW SECRET TOKEN **")
	fmt.Println("**")
	fmt.Printf("** SUBJECT: \"%s\" **", *subj)
	fmt.Println()

	s := signer.NewJWTHandler([]byte(k))

	claims := jwt.RegisteredClaims{
		Issuer:  signer.TokenIssuer,
		Subject: *subj,
		IssuedAt: &jwt.NumericDate{
			Time: time.Now().UTC(),
		},
	}
	secret, err := s.Sign(claims)
	if err != nil {
		fmt.Printf("could not sign claims: %s\n", err)
		os.Exit(1)
	}

	err = clipboard.WriteAll(secret)
	msg := "👆 Here's your key, we already copied it to your clipboard :)"
	if err != nil {
		msg = "👆 Here's your key."
	}

	fmt.Println()
	fmt.Println("** NEW KEY GENERATED 🎉 **")
	fmt.Println()
	fmt.Println(secret)
	fmt.Println()
	fmt.Println(msg)
	fmt.Println()
	fmt.Println("_______")
	fmt.Println()
	fmt.Println("IN CASE OF KEY COMPROMISED, GENERATE NEW SERVER \"SECRET_KEY\" AND PUSH TO CLOUD SERVER.")
	fmt.Println("_______")
}
