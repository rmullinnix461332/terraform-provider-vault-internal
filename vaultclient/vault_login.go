package vaultclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rmullinnix461332/logger"
)

func vaultLogin(server string, role string, jwt_path string) (AuthRecord, error) {
	var auth VaultAuth

	svcJwt, err := getToken(jwt_path)
    if err != nil {
        return auth.Auth, err
    }

	client := &http.Client{Transport: transport}

	vaultUrl := server + "/v1/auth/jwt/login"

	var loginReq VaultLogin

	loginReq.Role = role
	loginReq.Jwt = svcJwt

	buf, _ := json.Marshal(loginReq)

	req, err := http.NewRequest("POST", vaultUrl, bytes.NewBufferString(string(buf)))

	if err != nil {
		logger.Error.Println("Could not connect to vault", vaultUrl)
		return auth.Auth, err
	}

	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)

	if err != nil {
		logger.Error.Println("Cound not process vault request", err)
		return auth.Auth, err
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 || response.StatusCode == 201 {
		err = json.Unmarshal(body, &auth)

		if err != nil {
			logger.Error.Println("Cound not unmarshal vault response", err, string(body))
		}
	} else {
        logger.Error.Println("vault bad response", response.StatusCode, string(body))
    }

	return auth.Auth, err
}
