package controllers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
)

/// declare stuct ///
type Date struct {
	Day   *string `json:"day,omitempty"`
	Month *int    `json:"month,omitempty"`
}

// type DateFace struct {
// 	Day   *string `json:"day,omitempty"`
// 	Month *int    `json:"month,omitempty"`
// }
// type DateFaceFace struct {
// 	Day   *string `json:"day,omitempty"`
// 	Month *int    `json:"month,omitempty"`
// }

const (
	Sunday    string = "sunday"
	Monday    string = "monday"
	Tuesday   string = "tuesday"
	Wednesday string = "wednesday"
	Thursday  string = "thursday"
	Friday    string = "friday"
	Saturday  string = "saturday"
)
const (
	Januray int = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func FnIota(c echo.Context) (err error) {
	mDate := Date{}

	err = c.Bind(&mDate)
	if err != nil {
		log.Println(err)
		return err
	}

	// getType() to get struct name of variable
	if !validateEnumTip(mDate.Day, mDate.Month, getType(mDate)) {
		return fmt.Errorf("verify day/month enum is fail")
	}

	// stringBySlice() convert month index to string name
	toMonthName := stringBySlice(*mDate.Month)
	fmt.Println(toMonthName)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})

}

func validateEnumTip(pStr *string, pInt *int, pType string) bool {
	if pStr != nil {
		switch *pStr {
		case Sunday, Monday, Tuesday, Thursday, Friday, Saturday:
			return pType == "Date"
		}
	}
	if pInt != nil {
		switch *pInt {
		case Januray, February, March, April, May, June, July, August, September, October, November, December:
			return pType == "Date"
		}
	}
	return false
}

func getType(pStruct interface{}) string {
	t := reflect.TypeOf(pStruct)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func stringByArray(pKey int) string {
	arrayJa := [...]string{"januray", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	// arrayJa = append(arrayJa, "januray")
	return arrayJa[pKey-1]
}
func stringBySlice(pKey int) string {
	sliceJa := []string{"januray", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	// sliceJa = append(sliceJa, "januray")
	return sliceJa[pKey-1]
}
