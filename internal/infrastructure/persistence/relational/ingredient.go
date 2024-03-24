package relational

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/microsoft/go-mssqldb/integratedauth/krb5"
	"log"
)

type IngredientRepository struct {
	db *sql.DB
}

func NewIngredientRepository(connString string) IngredientRepository {
	sqlDb, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("could not create connection: %s", err)
	}

	defer sqlDb.Close()

	return IngredientRepository{db: sqlDb}
}
