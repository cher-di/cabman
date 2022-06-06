package edemrf

import "fmt"

type Car struct {
	Id          CustomUint32 `json:"id"`
	Name        string       `json:"name"`
	TotalPlaces CustomUint32 `json:"totalCount"`
	PlateNumber string       `json:"plateNumber"`
}

type User struct {
	Id     CustomUint32  `json:"id"`
	Name   string        `json:"name"`
	Rating CustomFloat32 `json:"rating"`
	Thumbs struct {
		Maxres string `json:"maxres"`
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"thumbs"`
}

func (user *User) GetBestQualityThumbUrl() (string, error) {
	thumbsRelUrls := [...]string{user.Thumbs.Maxres, user.Thumbs.Large, user.Thumbs.Medium, user.Thumbs.Small}
	for _, thumbRelUrl := range thumbsRelUrls {
		if thumbRelUrl != "" {
			return getImageFullUrl(thumbRelUrl), nil
		}
	}
	return "", fmt.Errorf("no thumbs for user %v", user)
}

type UserInfo struct {
	reponseStatus
	Data struct {
		User User  `json:"user"`
		Cars []Car `json:"userCars"`
	} `json:"data"`
}

func GetUserInfo(userId uint32) (UserInfo, error) {
	parsedUrl := getEndpoint("/users")
	parsedUrl.Path += fmt.Sprintf("/%d", userId)

	var userInfo UserInfo
	if err := sendGetRequest(parsedUrl.String(), &userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("failed to get user with user id %d: %v", userId, err)
	}
	return userInfo, nil
}
