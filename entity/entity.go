package entity

import "fmt"

type Entity interface {
	fmt.Stringer
	Validate() error
}
