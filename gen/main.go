package main

import (
	"context"
	"fmt"
	"go-gorm/gen/dal"
	"go-gorm/gen/dal/query"
	"log"
)

func main() {
	// biz.InsertPerson()

	conn := dal.Connect()
	query := query.Use(conn)
	people, err := query.Person.WithContext(context.Background()).Debug().GetMaxVersionCount()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("people: %#v\n", people)
}
