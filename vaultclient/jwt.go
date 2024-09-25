package vaultclient

import (
	"io/ioutil"
	"os"

    "github.com/rmullinnix461332/logger"
) 

func getToken(path string) (string, error) {
	tokenFile, err := os.Open(path)

	if err != nil {
		logger.Error.Println("[error] Unable to open service account token file", path, err)
		return "", err
    }

    byteToken, _ := ioutil.ReadAll(tokenFile)

	return string(byteToken), nil
}
