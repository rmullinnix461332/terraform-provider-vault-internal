package vaultclient

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rmullinnix461332/logger"
)

type VaultLogin struct {
	Role string `json:"role"`
	Jwt  string `json:"jwt"`
}

type VaultClient struct {
    Server string
    ClientToken string
}

type VaultAuth struct {
	Auth AuthRecord `json:"auth"`
}

type AuthRecord struct {
	ClientToken   string   `json:"client_token"`
	Accessor      string   `json:"accessor"`
	Policies      []string `json:"polcies"`
	LeaseDuration int      `json:"lease_duration"`
	Renewable     bool     `json:"renewable"`
}

type VaultSecret struct {
	RequestID     string    `json:"request_id"`
	LeaseID       string    `json:"lease_id"`
	Renewable     bool      `json:"renewable"`
	LeaseDuration int       `json:"lease_duration"`
	Data          VaultData `json:"data"`
}

type VaultData struct {
	Data     map[string]string `json:"data"`
	Metadata struct {
		CreatedTime    time.Time `json:"created_time"`
		CustomMetadata any       `json:"custom_metadata"`
		DeletionTime   string    `json:"deletion_time"`
		Destroyed      bool      `json:"destroyed"`
		Version        int       `json:"version"`
	} `json:"metadata"`
}

var transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

func NewVaultClient(server string, role string, jwt_path string) (VaultClient, error) {
    var client VaultClient

    client.Server = server
    auth, err := vaultLogin(server, role, jwt_path)

    if err != nil {
        return client, err
    }

    client.ClientToken = auth.ClientToken

    return client, nil
}

func (vc *VaultClient) GetSecrets(secretPath string) (map[string]string, error) {
	secrets := make(map[string]string, 0)
    logger.Info.Println("path", secretPath)
	client := &http.Client{Transport: transport}

	reqUrl := vc.Server +  "/v1/" + secretPath

	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		logger.Error.Println("Could not connect to vault:", reqUrl)
		return secrets, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Vault-Token", vc.ClientToken)

	response, err := client.Do(req)

	if err != nil {
		logger.Error.Println("Cound not process vault request:", err.Error())
		return secrets, err
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {
		var vaultSecret VaultSecret
		err = json.Unmarshal(body, &vaultSecret)
		if err != nil {
			logger.Error.Println("Vault secrets unmarshal error:", err.Error())
			return secrets, err
		}
		return vaultSecret.Data.Data, nil
	}

	return secrets, errors.New("Could not retrieve secrets from vault: " + response.Status)
}
