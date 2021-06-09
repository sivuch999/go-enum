package models

type Day int

type Week struct {
	Day *Day `json:"day,omitempty"`
}

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// String - Creating common behavior - give the type a String function
func (w Day) StringByArray() string {
	arrayJa := [...]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	// arrayJa = append(arrayJa, "januray")
	return arrayJa[w]
}
func (w Day) StringBySlice() string {
	sliceJa := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	// sliceJa = append(sliceJa, "januray")
	return sliceJa[w]
}

//Value validate enum when set to database
// func (t Day) Value() (driver.Value, error) {
// 	switch t {
// 	case Sunday, Monday, Tuesday: //valid case
// 		return string(t), nil
// 	}
// 	return nil, fmt.Errorf("Invalid product type value") //else is invalid
// }
