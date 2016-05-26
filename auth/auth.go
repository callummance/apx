package auth

import "errors"

type User struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
}

const defaultAvatar string = ""

//Generates a new user struct from data fetched from Facebook
func BuildUser(details FbDetails) User {
	userDetails := getUserAdvancedDetails(details.User_id)
	var newUser User
	newUser.UID = details.User_id
	newUser.Name = userDetails.Name
	newUser.Email = userDetails.Email
	newUser.Avatar = defaultAvatar

	return newUser
}

//Authenticates an existing user session and retrieves the
//user ID it was assigned to
func AuthenticateSession(sid string) string {
	//Lookup session id
	//Check if it is valid and has not expired
	//Return uid if true
	//Otherwise return empty string
	return ""
}

func AuthenticateUser(token string) (string, error) {
	userDetails := GetUserDetails(token)
	if !userDetails.Is_valid {
		return "", errors.New("Token provided is not valid or has expired.")
	}

	//Lookup the user in database
	//If the user exists, return the user id and nil
	//Otherwise, make a new user with BuildUser, and add it to the database
	return "", nil
}
