package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/yusronhanan-m/generate-jwt/config"
	"github.com/yusronhanan-m/generate-jwt/global"
	"golang.org/x/crypto/bcrypt"
)

func printJson(objectPublic interface{}) {

	b, err := json.Marshal(objectPublic)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println("===")
	fmt.Println(string(b))
}
func GenerateJWTToken(claims map[string]interface{}) (*string, error) {
	jwtauth.SetExpiry(claims, time.Now().Add(time.Hour*48))
	tokenAuth := jwtauth.New("HS256", []byte(config.GetValue(config.JWT_SECRET)), nil)
	_, jwtToken, err := tokenAuth.Encode(claims)
	if err != nil {
		return nil, err
	}
	return &jwtToken, nil
}
func GetClaimsFromContext(ctx context.Context) (map[string]interface{}, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
func generateJWT(name string, role string) {
	newUUID, _ := uuid.NewRandom()
	newUUID = uuid.MustParse("7ef91ba9-dae7-4d66-af78-e2607b61f3d9")
	claims := map[string]interface{}{
		"uuid": newUUID,
		"name": name,
		"role": role,
	}
	loginToken, err := GenerateJWTToken(claims)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newUUID, name, role)
		fmt.Printf("%s", *loginToken)
		fmt.Println("")
	}
}

func trialMapInterface() {
	newListMap := make([]map[string]interface{}, 0)

	for i := 0; i < 5; i++ {
		newMap := map[string]interface{}{
			strconv.Itoa(i): "i",
			"height":        i,
			"unit":          "cm",
		}
		newListMap = append(newListMap, newMap)
	}

	newMap2 := map[string]interface{}{
		"height": 100,
		"unit":   "cm",
	}
	newListMap = append(newListMap, newMap2)

	fmt.Println(newListMap)
}

func makehttpReq() {
	/*
	* Http (curl) request in golang
	* @author Shashank Tiwari
	 */

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://api.medicplus.co.id//doctor/id?id=d35e1823-a2ed-4783-83f7-0ec17688169b&language_code=id")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

}
func main() {
	date := time.Date(2000, 5, 5, 0, 0, 0, 0, time.Local).UTC()
	fmt.Println(date)
	pass := global.GeneratePassword()
	pass = "admin123"
	fmt.Println(pass)
	bcryptHash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	fmt.Println(string(bcryptHash))

	generateJWT("Dr Vanya", "doctor")
	//admin
	generateJWT("Karen", "admin")
	//patient
	generateJWT("Josh", "patient")

	// fmt.Println("=======")
	// printJson(public.UpdateDoctorApplicantRequest{})
	// fmt.Println("=======")
	// trialMapInterface()

	// makehttpReq()
}
