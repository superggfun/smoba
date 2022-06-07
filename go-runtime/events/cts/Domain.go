package cts

import (
    "fmt"
)

type Domain struct {
    Name string `json:"name"`
    Id string `json:"id"`
}

func (d *Domain) String() string {
    return fmt.Sprintf(`Domain{
                                 name='%v',
                                 id='%v'
                               }`, d.Name, d.Id)
}