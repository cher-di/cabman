package edemrf

import "fmt"

type Car struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	TotalPlaces CustomUint32 `json:"totalCount"`
	PlateNumber string       `json:"plateNumber"`
}

type User struct {
	Id     CustomUint32  `json:"id"`
	Name   string        `json:"name"`
	Rating CustomFloat32 `json:"rating"`
	Cars   []Car
}

type UserInfo struct {
	reponseStatus
	Data struct {
		User User  `json:"user"`
		Cars []Car `json:"userCars"`
	} `json:"data"`
}

func GetUserInfo(userId string) (UserInfo, error) {
	parsedUrl := GetEndpoint("/users")
	parsedUrl.Path += fmt.Sprintf("/%s", userId)

	var userInfo UserInfo
	if err := sendGetRequest(parsedUrl.String(), &userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("failed to get user with user id %s: %v", userId, err)
	}
	return userInfo, nil
}
