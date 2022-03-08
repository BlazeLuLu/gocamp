package apis

type PersonalInformation struct {
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
}
