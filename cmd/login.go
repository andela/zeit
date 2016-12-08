package cmd

import (
	"crypto/tls"
	"encoding/json"
	"github.com/andela/zeit/lib"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"fmt"
)

var user lib.User

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Print the version number of Oge",
	Long:  `All software has versions. This is Oge's`,
	Run:   login,
}

func authenticateUser(config *lib.Config) {
	cmdString := "https://api-staging.andela.com/login?redirect_url=http://localhost:8089"
	open.Run(cmdString)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query()["token"][0]
		user, _ := retrieveUserInfo(token)
		config := lib.NewConfigFromFile()
		config.CurrentUser = lib.User{
			Id:    user.Id,
			Token: token,
		}
		config.Save()
	})
	http.ListenAndServe(":8089", nil)
}

func retrieveUserInfo(token string) (lib.User, error) {
	req, err := http.NewRequest("GET", "https://api-staging.andela.com/api/v1/users/me", nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	if err != nil {
		panic(err)
	} else {
		bearerToken := "Bearer " + token
		req.Header.Add("Authorization", bearerToken)
		if res, err := client.Do(req); err != nil {
			panic(err)
		} else {
			bytes, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			err = json.Unmarshal(bytes, &user)
		}
	}
	return user, err
}

func login(cmd *cobra.Command, args []string) {
	config := lib.NewConfigFromFile()
	if config.CurrentUser.Id != "" {
		_, err := retrieveUserInfo(config.CurrentUser.Token)
		if err != nil {
			authenticateUser(config)
		}
		fmt.Println("You are already logged in")
	} else {
		authenticateUser(config)
	}
}

func init() {
	RootCmd.AddCommand(loginCmd)
}
