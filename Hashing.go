package main

import (
	"crypto/rand"
	"crypto/sha256"
	// "crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

/*
Salting addes an extra layer of security for our hashed passwords
Salting protects our password from dictonary attacks
Why Salting is Important

Prevents Rainbow Table Attacks

Ensures same passwords produce different hashes

Increases overall password security
*/

func main(){
	password := "password123"
	// hash := sha256.Sum256([]byte(password))

	// password1 := "password123"
	// hash1 := sha512.Sum512([]byte(password1))
	// fmt.Println(hash1)
	// fmt.Println()
	// fmt.Println(hash)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt", err)
	}

	//Hash the password with the salt
	hash := HashPassword(password, salt)

	//store the salt and password in database. Just printing as of now.
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt: ", saltStr)
	fmt.Println("Hash: ", hash)

	//verify the password

	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error in decoding")
	}

	LoginPass := "pass"
	logingHash := HashPassword(LoginPass, decodedSalt)

	//Compare the stored hash with loggingHash

	if hash == logingHash {
		fmt.Println("Login Succesful")
	} else {
		fmt.Println("Incorrect Password")
	}
}

func generateSalt() ([]byte, error){
	salt := make([]byte, 16)//We generate a slice of size 16 bytes 
	_, err := io.ReadFull(rand.Reader, salt)// What this line doing ReadeFull takes to arguments one is reader which generate some random numbers and then stored on to byte slice salt 

	if err != nil {
		return nil, err
	}
	return salt, nil
}

//Function to hash password

func HashPassword (password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)

	return base64.StdEncoding.EncodeToString(hash[:])
}