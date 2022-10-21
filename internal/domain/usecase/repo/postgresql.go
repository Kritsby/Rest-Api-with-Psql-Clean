package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"rest-api/internal/domain/entity"
)

type userStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *userStorage {
	return &userStorage{db: db}
}

func (us *userStorage) GetAllAdapter() []entity.FirstTable {
	var data []entity.FirstTable

	query := `select id, first_name, second_name from first_table`

	fmt.Println(query)
	rows, err := us.db.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("error while executing query: %s", err)
	}

	for rows.Next() {
		var id int
		var firstName string
		var secondName string

		err := rows.Scan(&id, &firstName, &secondName)
		if err != nil {
			log.Fatalf("can't scan all data : %s", err)
		}

		data = append(data, entity.FirstTable{
			ID:         id,
			FirstName:  firstName,
			SecondName: secondName,
		})
	}
	return data
}
