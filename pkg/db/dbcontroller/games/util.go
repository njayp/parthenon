package games

import (
	"fmt"
	"strings"
)

const (
	USER_LOCATION_VAR_NAME = "@userLocation%s"
)

func UserLocationVarName(userid string) string {
	return fmt.Sprintf(USER_LOCATION_VAR_NAME, strings.Replace(userid, "-", "", -1))
}
