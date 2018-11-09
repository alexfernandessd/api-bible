package main

import (
	"database/sql"
	"fmt"

	bible "github.com/alexfernandessd/api-bible/bible"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := bible.NewConfig()

	// sess, err := session.NewSession(&aws.Config{
	// 	Region:      aws.String("region"),
	// 	Credentials: credentials.NewSharedCredentials("/home/user/.aws/credentials", "personal"),
	// })

	awsCreds := NewSharedCredentials(config.FileCreds, config.ProfileCreds)

	_, err := rdsutils.BuildAuthToken(config.MySqlbEndpoint, config.AWSRegion, config.AWSUser, awsCreds)
	fmt.Println("config", config.AWSUser)

	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=false",
		"bible", config.AWSPassword, "bible.cgrmdcvj2f27.us-east-2.rds.amazonaws.com", "bible",
	)
	fmt.Println("dns", dnsStr)

	db, _ := sql.Open("mysql", dnsStr)

	rows, err := db.Query("SELECT * FROM verses where id = 4000")

	for rows.Next() {
		var id int
		var version string
		var testament int
		var book int
		var chapter int
		var verse int
		var text string
		err = rows.Scan(&id, &version, &testament, &book, &chapter, &verse, &text)
		fmt.Printf("chapter: %d, verse: %d: ", chapter, verse)
		fmt.Println(text)
	}

	if err != nil {
		fmt.Println("error: ", err)
	}
}

func NewSharedCredentials(filename, profile string) *credentials.Credentials {
	return credentials.NewCredentials(&credentials.SharedCredentialsProvider{
		Filename: filename,
		Profile:  profile,
	})
}
