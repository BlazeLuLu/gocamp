module gocamp

go 1.17

require (
	github.com/armstrongli/go-bmi v0.0.1
	google.golang.org/grpc v1.45.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

replace github.com/armstrongli/go-bmi v0.0.1 => ./staging/src/github.com/armstrongli/go-bmi
