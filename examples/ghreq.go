package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GhUserInfo struct {
	Login     string
	Id        int
	Followers int
	NodeId    string
}

func BuildGhUrl(login string) string {
	return fmt.Sprintf("https://api.github.com/users/%v", login)
}

func GetGhUserData(login string) GhUserInfo {
	resp, err := http.Get(BuildGhUrl(login))
	if err != nil {
		return GhUserInfo{
			Login: login,
			Id:    -1,
		}
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var ghi GhUserInfo

	uerr := json.Unmarshal(body, &ghi)

	if uerr != nil {
		return GhUserInfo{
			Login: login,
			Id:    -1,
		}
	}

	// fmt.Println(string(body))

	return ghi
}

func main() {
	fmt.Println(GetGhUserData("geekan"))
}
