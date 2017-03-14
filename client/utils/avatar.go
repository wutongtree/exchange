package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

func GetAvatarSource(avatar string) string {
	if "" == avatar {
		return "/static/img/avatar/picture.jpg"
	}

	return strings.Replace(avatar, "-cropper", "", -1)
}

func GetAvatar(avatar string) string {
	if "" == avatar {
		return fmt.Sprintf("/static/img/avatar/%d.jpg", rand.Intn(5))
	}

	return fmt.Sprintf("/static/img/avatar/%v.jpg", "picture")
}
