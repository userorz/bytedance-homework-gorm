package biz

import (
	"context"
	"go-gorm/gen/dal"
	"go-gorm/gen/dal/model"
	"go-gorm/gen/dal/query"
	"log"
	"math/rand"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//插入100条数据，重复的UUID则Version+1
func InsertPerson() {
	conn := dal.Connect()
	query := query.Use(conn)
	// fmt.Printf("query.Person.ALL: %v\n", query.Person.ALL)

	UUID := make([]string, 50)
	for i := range UUID {
		u, err := uuid.NewUUID()
		if err != nil {
			log.Fatal(err)
		}
		UUID[i] = u.String()
	}

	for i := 0; i < 100; i++ {
		j := rand.Intn(50)

		u, err := query.Person.WithContext(context.Background()).Debug().Where(query.Person.UUID.Eq(UUID[j])).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = query.Person.WithContext(context.Background()).Create(&model.Person{
					UUID:    UUID[j],
					Name:    "2",
					Age:     1,
					Version: 1,
				})
				if err != nil {
					log.Printf("Create fail:%v\n", err)
				}
				continue
			}
			log.Printf("Find fail:%v\n", err)
			continue
		}

		_, err = query.Person.WithContext(context.Background()).Where(query.Person.UUID.Eq(UUID[j])).Update(query.Person.Version, u.Version+1)
		if err != nil {
			log.Printf("Update fail:%v\n", err)
			continue
		}
	}
}
