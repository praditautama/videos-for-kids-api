package middleware

import (  
	"github.com/labstack/echo"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}