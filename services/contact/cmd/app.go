package main

import (
	"fmt"

	"buzukbuzuk.clean-architecture/pkg/store/postgres"
	"buzukbuzuk.clean-architecture/services/contact/internal/domain"
)

func main() {
	dcp := &postgres.DbConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "Example123+",
		DbName:   "cleanarc",
	}

	db, err := postgres.OpenDB(dcp)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	ilya := domain.NewContact("Kochemirov", "Ilya", "Alekseevich")
	anar := domain.NewContact("Daureneva", "Anar", "Alibekkyzy")
	group1 := domain.NewGroup("Group 1")

	fmt.Println(ilya, anar, group1)
}
