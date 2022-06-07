package cts

import (
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Id string `json:"id"`
    Domain Domain `json:"domain"`
}

func (u *User) String() string {
    return fmt.Sprintf(`User{
                                 name=%v,
                                 id=%v,
                                 domain=%+v
                               }`, u.Name, u.Id, u.Domain)
}