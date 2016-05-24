package authsdk

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

//ApigeeJWTToken the apigee impelmentation of the JWT token auth sdk
type ApigeeJWTToken struct {
	originalToken string
	tokenPayload  tokenPayload
}

//NewApigeeJWTToken create a new Apigee JWT token.  Return the instance or an error if one cannot be created
func NewApigeeJWTToken(token string) (JWTToken, error) {

	parts := strings.Split(token, ".")

	length := len(parts)

	if length != 3 {
		return nil, fmt.Errorf("Expected JWT token to contain a header, a payload, and a signature.  Only received %d parts", length)
	}

	partToDecode := parts[1]

	// fmt.Printf("Decoding base64 value of %s\n", partToDecode)

	decodedPayload, err := base64.RawURLEncoding.DecodeString(partToDecode)

	if err != nil {
		return nil, fmt.Errorf("Unable to decode payload.  Decode error: %s", err)
	}

	// fmt.Printf("Decoded payload is %s\n", decodedPayload)

	var tokenPayload tokenPayload

	err = json.Unmarshal([]byte(decodedPayload), &tokenPayload)

	if err != nil {
		return nil, err
	}

	return &ApigeeJWTToken{
		originalToken: token,
		tokenPayload:  tokenPayload,
	}, nil
}

//GetSubject get the subject claim from the token
func (token *ApigeeJWTToken) GetSubject() string {
	return token.tokenPayload.Subject
}

//GetUsername return the username if possible for the subject
func (token *ApigeeJWTToken) GetUsername() string {
	return token.tokenPayload.Username
}

//GetEmail return the username if possible for the subject
func (token *ApigeeJWTToken) GetEmail() string {
	return token.tokenPayload.Email
}

//IsOrgAdmin is the current JWTToken subject an organization admin.  If so return true, if not, return false.  An error is returned if the check cannot be performed
func (token *ApigeeJWTToken) IsOrgAdmin(orgName string) (bool, error) {
	//TODO, implement this
	return true, nil
}

type tokenPayload struct {
	Subject  string `json:"sub"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
}
