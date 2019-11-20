package main

import (
	"finalproject_ecomerce/adapters/web"
	"finalproject_ecomerce/engine"
	"finalproject_ecomerce/providers"
	mysql2 "finalproject_ecomerce/providers/mysql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"upper.io/db.v2/mysql"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	dbURL, err := parseDBURL("mysql://root:masok12345@localhost:3306/finalproject")
	if err != nil {
		log.Fatal(err)
	}
	connURL, err := mysql.ParseURL(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	session, err := mysql.Open(connURL)
	if err != nil {
		log.Fatal(err)
	}

	// Setup storage factory
	var sf engine.StorageFactory
	sf = mysql2.NewStorage(session)

	// Setup service dependencies
	var (
		validator  engine.Validator
		mailSender engine.MailSender
		jwt        engine.JWTSignParser
	)

	validator = providers.NewValidator()
	mailSender = providers.NewFakeMail()
	jwt = providers.NewJWT()
	emitter := providers.NewEmitter()

	f := engine.New(sf, mailSender, validator, jwt, emitter)

	log.Printf("server starting port: %s", "5000")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", "5000"), web.NewWebAdapter(f)); err != nil {
		log.Fatal(err)
	}
}

func parseDBURL(s string) (string, error) {
	durl, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("cannot parse database url, err:%s", err)
	}
	user := durl.User.Username()
	password, _ := durl.User.Password()
	host := durl.Host
	dbname := durl.Path // like: /path

	return fmt.Sprintf("%s:%s@tcp(%s)%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbname), nil
}
