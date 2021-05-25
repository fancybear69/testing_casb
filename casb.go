package main
import (
	"fmt"
	"github.com/casbin/casbin/v2"
)
e := casbin.NewEnforcer("model.conf", "model.csv")
a = e.HasPermissionForUser(alice, []string{'read'})
fmt.Printf("The  value is %d", a)


