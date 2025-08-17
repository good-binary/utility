package uuid

import (
	"time"

	"github.com/good-binary/utility/random"
)

func Newuuid() string {
	t := time.Now()
	x := t.Format("20060102150405")
	y := random.RandomString(8, "", "")
	uuid := x + y
	return uuid
}
