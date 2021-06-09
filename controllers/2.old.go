package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

/// declare stuct ///
type Today struct {
	Feeling     *string `json:"feeling,omitempty"`
	FeelingEnum *FeelingEnum
}
type FeelingEnum struct {
	Good string
	Bad  string
}

/// function ///
func FnOld(c echo.Context) (err error) {
	mToday := Today{}

	err = c.Bind(&mToday)
	if err != nil {
		log.Println(err)
		return err
	}

	if !verifyEnumOld(mToday.Feeling) {
		return fmt.Errorf("verify feeling enum is fail")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

/// service ///
func verifyEnumOld(pFeeling *string) bool {
	var feeling = getFeeling()
	switch *pFeeling {
	case feeling.Good, feeling.Bad:
		return true
	}
	return false
}

func getFeeling() *FeelingEnum {
	return &FeelingEnum{
		Good: "good",
		Bad:  "bad",
	}
}
