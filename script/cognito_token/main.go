package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
)

var (
	poolID       string
	clientID     string
	clientSecret string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	poolID = os.Getenv("AWS_COGNITO_USER_POOL_ID")
	clientID = os.Getenv("AWS_COGNITO_CLIENT_ID")
	clientSecret = os.Getenv("AWS_COGNITO_CLIENT_SECRET")
}

// https://dev.to/mcharytoniuk/using-aws-cognito-app-client-secret-hash-with-go-8ld
// https://dev.classmethod.jp/articles/change-cognito-user-force_change_passwore-to-confirmed/
func main() {
	// 引数処理
	userName := flag.String("u", "", "user name")
	password := time.Now().String()
	flag.Parse()
	if *userName == "" {
		fmt.Println("invalid parameter")
		os.Exit(1)
	}

	// クライアント生成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	cognitoClient := cognitoidentityprovider.New(sess)

	// ユーザー作成
	_, err := cognitoClient.AdminCreateUser(&cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        aws.String(poolID),
		Username:          userName,
		TemporaryPassword: aws.String(password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: userName,
			},
		},
	})
	if err != nil {
		fmt.Println("Got error creating user:", err)
		os.Exit(1)
	}

	// シークレットハッシュ生成
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(*userName + clientID))
	secretHash := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	// 認証フロー初期化
	adminInitiateAuthOutput, err := cognitoClient.AdminInitiateAuth(&cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:   aws.String("ADMIN_USER_PASSWORD_AUTH"),
		UserPoolId: aws.String(poolID),
		ClientId:   aws.String(clientID),
		AuthParameters: map[string]*string{
			"USERNAME":    userName,
			"PASSWORD":    aws.String(password),
			"SECRET_HASH": aws.String(secretHash),
		},
	})
	if err != nil {
		fmt.Println("Got error admin initiate auth:", err)
		os.Exit(1)
	}

	// 認証フロー応答
	adminRespondToAuthChallengeOutput, err := cognitoClient.AdminRespondToAuthChallenge(&cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		UserPoolId:    aws.String(poolID),
		ClientId:      aws.String(clientID),
		ChallengeName: aws.String("NEW_PASSWORD_REQUIRED"),
		ChallengeResponses: map[string]*string{
			"USERNAME":     userName,
			"NEW_PASSWORD": aws.String(password),
			"SECRET_HASH":  aws.String(secretHash),
		},
		Session: adminInitiateAuthOutput.Session,
	})
	if err != nil {
		fmt.Println("Got error admin respond to auth challenge:", err)
		os.Exit(1)
	}

	// トークン出力
	fmt.Println("AccessToken: ", *adminRespondToAuthChallengeOutput.AuthenticationResult.AccessToken)
	fmt.Println("IdToken: ", *adminRespondToAuthChallengeOutput.AuthenticationResult.IdToken)
}
