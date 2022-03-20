package cognito

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var (
	poolID       = os.Getenv("COGNITO_POOL_ID")
	clientID     = os.Getenv("COGNITO_CLIENT_ID")
	clientSecret = os.Getenv("COGNITO_CLIENT_SECRET")
)

// https://dev.to/mcharytoniuk/using-aws-cognito-app-client-secret-hash-with-go-8ld
// https://dev.classmethod.jp/articles/change-cognito-user-force_change_passwore-to-confirmed/
func CreateUser() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	cognitoClient := cognitoidentityprovider.New(sess)

	userName := flag.String("u", "", "user name")
	password := time.Now().String()

	flag.Parse()

	if *userName == "" {
		fmt.Println("invalid parameter")
		os.Exit(1)
	}

	newUserData := &cognitoidentityprovider.AdminCreateUserInput{}
	newUserData.SetUserPoolId(poolID)
	newUserData.SetUsername(*userName)
	newUserData.SetTemporaryPassword(password)
	newUserData.SetUserAttributes([]*cognitoidentityprovider.AttributeType{
		{
			Name:  aws.String("email"),
			Value: userName,
		},
	})

	_, err := cognitoClient.AdminCreateUser(newUserData)
	if err != nil {
		fmt.Println("Got error creating user:", err)
		os.Exit(1)
	}

	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(*userName + clientID))
	secretHash := base64.StdEncoding.EncodeToString(mac.Sum(nil))

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
		fmt.Println(err)
		os.Exit(1)
	}

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
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(*adminRespondToAuthChallengeOutput.AuthenticationResult.AccessToken)
}
