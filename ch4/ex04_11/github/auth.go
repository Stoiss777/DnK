/**
 * The authorization functions.
 * 
 * - Client ID must be stored in ".client_id" file in the app directory.
 * - The authorization token is saved in the "token" file.
 * 
 * Github manual:
 * https://docs.github.com/en/apps/creating-github-apps/writing-code-for-a-github-app/building-a-cli-with-a-github-app
 */
package github

import (
	"encoding/json"
	"fmt"
	"os"
	"net/http"
	"net/url"
	"time"
	"strings"
)


/**
 * The primary function of authorization.
 * 
 * Checks whether the user is already logged in. If not, it offers them to log in.
 * Upon successful login, an "access code" is returned for Bearer authentication.
 * 
 **/
func Login() (*User, error) {
	clientID, err := GetClientID(); 
	if err != nil {
		return nil, err
	}
	accessToken := GetLastAccessToken(); 
	if accessToken != "" {
		if user, err := Whoami(accessToken); err == nil {
			return user, nil
		}
	}

	deviceCodeResponse, err := RequestDeviceCode(clientID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Login URL: %s\nCode: %s\n",
		deviceCodeResponse.VerificationUri, deviceCodeResponse.UserCode)

	var requestInterval time.Duration = AccessTokenRequestInterval
	var accessTokenResponse *AccessTokenResponse
	for true {
		time.Sleep(requestInterval * time.Second)
		accessTokenResponse, err = RequestAccessToken(clientID, deviceCodeResponse.DeviceCode)
		if err != nil {
			return nil, err
		}		
		if accessTokenResponse.AccessToken != "" {
			break
		}
		// fmt.Println(accessTokenResponse.Error)
		switch accessTokenResponse.Error {
		case "authorization_pending":
			fmt.Println("Waiting for authorization...")
		case "slow_down":
			requestInterval += 5
		default:
			return nil, fmt.Errorf("Access token request failed: %s", 
				accessTokenResponse.ErrorDescription)
		}
	}

	if user, err := Whoami(accessTokenResponse.AccessToken); err == nil {
		saveAccessToken(accessTokenResponse.AccessToken)
		return user, nil
	}

	return nil, fmt.Errorf("No access token found.")
}


/**
 * Retrieves the authorized user's data.
 * 
 */
func Whoami(accessToken string) (*User, error) {
	req, err := http.NewRequest("GET",	WhoamiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		resp.Body.Close()
		return nil, fmt.Errorf("Bad credentials")
	} else if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Response status: %s", resp.Status)
	}

	res := &User{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return res, nil
}


/**
 * The 1st authorization step - retrieve a device code.
 * 
 */
func RequestDeviceCode(clientID string) (*DeviceCodeResponse, error) {
	form := url.Values{
		"client_id": {clientID},
	}
	req, err := http.NewRequest(
		"POST",
		DeviceCodeURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		resp.Body.Close()
		return nil, fmt.Errorf("Device not found")
	} else if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Device code request failed: %s", resp.Status)
	}

	res := &DeviceCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return res, nil
}


/**
 * The 2nd authorization step - change a device code to an access token.
 * 
 */
func RequestAccessToken(clientID, deviceCode string) (*AccessTokenResponse, error) {
	form := url.Values{
		"client_id": {clientID},
		"device_code": {deviceCode},
		"grant_type": {"urn:ietf:params:oauth:grant-type:device_code"},
	}
	req, err := http.NewRequest(
		"POST",
		AccessTokenURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Access code request failed: %s", resp.Status)
	}

	res := &AccessTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return res, nil
}


func GetClientID() (string, error) {
	clientID, err := os.ReadFile(ClientIDFile)
	if err != nil {
		return "", err
	}

	return string(clientID), nil
}


func GetLastAccessToken() string {
	token, err := os.ReadFile(AccessTokenFile)
	if err == nil {
		return string(token)
	}

	return ""
}

func saveAccessToken(accessToken string) error {
	if err := os.WriteFile(AccessTokenFile, []byte(accessToken), 0666); err != nil {
		return err
	}

	return nil
}
