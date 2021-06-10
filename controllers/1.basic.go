package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

/// declare stuct ///
type AnimalName string
type AnimalColor int
type Animal struct {
	Name  *AnimalName  `json:"name,omitempty"`
	Color *AnimalColor `json:"color,omitempty"`
}

const (
	Dog AnimalName = "dog"
	Cat AnimalName = "cat"
)
const (
	Red   AnimalColor = iota // 0
	Green                    // 1
)

/// function ///
func FnBasic(c echo.Context) (err error) {
	mAnimal := Animal{}

	err = c.Bind(&mAnimal)
	if err != nil {
		log.Println(err)
		return err
	}

	if !mAnimal.Name.verifyEnum() {
		return fmt.Errorf("verify name enum is fail")
	}
	if !mAnimal.Color.verifyEnum() {
		return fmt.Errorf("verify color enum is fail")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

/// service ///
func (pName *AnimalName) verifyEnum() bool {
	switch *pName {
	case Dog, Cat:
		return true
	}
	return false
}
func (pColor *AnimalColor) verifyEnum() bool {
	switch *pColor {
	case Red, Green:
		return true
	}

	return false
}
