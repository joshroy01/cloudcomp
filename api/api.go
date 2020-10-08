package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"errors"
	"strconv"
)

//Declare a global array of Credentials
//See credentials.go

/*YOUR CODE HERE*/
var users []Credentials = []Credentials{}


func RegisterRoutes(router *mux.Router) error {

	/*
	Fill out the appropriate get methods for each of the requests, based on the nature of the request.
	Think about whether you're reading, writing, or updating for each request
	*/

	router.HandleFunc("/api/getCookie", getCookie).Methods(http.MethodGet)
	router.HandleFunc("/api/getQuery", getQuery).Methods(http.MethodGet)
	router.HandleFunc("/api/getJSON", getJSON).Methods(http.MethodGet)
	
	router.HandleFunc("/api/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/api/getIndex", getIndex).Methods(http.MethodGet)
	router.HandleFunc("/api/getpw", getPassword).Methods(http.MethodGet)
	router.HandleFunc("/api/updatepw", updatePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/deleteuser", deleteUser).Methods(http.MethodDelete)

	return nil
}

func getCookie(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "access_token" cookie's value and write it to the response
		If there is no such cookie, write an empty string to the response
	*/

		//Obtain the the cookie "access_token", this gives us all the chararcteristics of the cookie
		cookie, err := request.Cookie("access_token")

		//check if obtaining the cookie returned an error
		if err != nil {
			fmt.Fprintf(response, "")
			return
		}

		fmt.Fprintf(response, cookie.Value)
}

func getQuery(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "userID" query paramter and write it to the response
		If there is no such query parameter, write an empty string to the response
	*/

	username := request.URL.Query().Get("userID")
	
	//it will return empty if it doesn't exist anyway
	fmt.Fprint(response, username)
}

func getJSON(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:
		{
			"username" : <username>,
			"password" : <password>
		}
		Decode this json file into an instance of Credentials.
		Then, write the username and password to the response, separated by a newline.request
		
		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	credentials := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credentials.Username == "" || credentials.Password == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(response, credentials.Username + "\n")
	fmt.Fprintf(response, credentials.Password)

	return
}

func signup(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:
		{
			"username" : <username>,
			"password" : <password>
		}
		Decode this json file into an instance of Credentials.
		Then store it ("append" it) to the global array of Credentials.
		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	credentials := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credentials.Username == "" || credentials.Password == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	users = append(users, credentials)

	response.WriteHeader(201)
	return
}

func getIndex(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:
		{
			"username" : <username>
		}
		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)
		Return the array index of the Credentials object in the global Credentials array
		
		The index will be of type integer, but we can only write strings to the response. What library and function was used to get around this?
		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	credentials := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credentials.Username == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == credentials.Username {
			fmt.Fprintf(response, strconv.Itoa(i))
			return
		}
	}

	//user doesnt exist
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}

func getPassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:
		{
			"username" : <username>
		}
		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)
		Write the password of the specific user to the response
		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	credentials := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credentials.Username == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == credentials.Username {
			fmt.Fprintf(response, users[i].Password)
			return
		}
	}

	//user doesnt exist
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}



func updatePassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:
		{
			"username" : <username>,
			"password" : <password,
		}
		Decode this json file into an instance of Credentials. 
		The password in the JSON file is the new password they want to replace the old password with.
		You don't need to return anything in this.
		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	credentials := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credentials.Username == "" || credentials.Password == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == credentials.Username {
			users[i].Password = credentials.Password
			return
		}
	}

	//user doesnt exist
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}

func deleteUser(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:
		{
			"username" : <username>,
			"password" : <password,
		}
		Decode this json file into an instance of Credentials.
		Remove this user from the array. Preserve the original order. You may want to create a helper function.
		This wasn't covered in lecture, so you may want to read the following:
			- https://gobyexample.com/slices
			- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/
		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	
	credentials := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	if credentials.Username == "" || credentials.Password == "" {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == credentials.Username &&  users[i].Password == credentials.Password {
			users = remove(users, i)
			return
		}
	}

	//user doesnt exist
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return

}

func remove(slice []Credentials, s int) []Credentials {
    return append(slice[:s], slice[s+1:]...)
}