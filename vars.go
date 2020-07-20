package main

import "fmt"

var (
	domain = "https://buntu.ga"

	dbHost = "localhost"
	dbPort = "3306"
	dbUsername = "root"
	dbPassword = "Abc123^^"
	dbName = "buntunet"

	pubApikey = "1f8cdf07622716653b9968bf07d0df10181ceb370a7b7a37bac837bd6cc3262f"
	priApikey = "91D3414C30a079dCf5022B1fc7F062243835eb486b1aA23aba66a88aEa250077"

	dbQuery = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
)
