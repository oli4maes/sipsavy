package relational

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/microsoft/go-mssqldb/integratedauth/krb5"
)

type IngredientRepository struct {
	db *sql.DB
}

func NewIngredientRepository(connString string) IngredientRepository {
	sqlDb, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("could not create connection: %s", err)
	}

	return IngredientRepository{db: sqlDb}
}

type Ingredient struct {
	Id             int
	Name           string
	Created        time.Time
	CreatedBy      string
	LastModified   time.Time
	LastModifiedBy string
}

func (repo *IngredientRepository) GetAll(ctx context.Context) ([]Ingredient, error) {
	var ingredients []Ingredient

	query := `SELECT *
			FROM Mixology.Mixology.Ingredients`

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer repo.db.Close()

	var id int
	var name, createdBy, lastModifiedBy string
	var created, lastModified time.Time
	for rows.Next() {
		if err := rows.Scan(&id, &name, &created, &createdBy, &lastModified, &lastModifiedBy); err != nil {
			return nil, err
		}

		var ingredient = Ingredient{
			Id:             id,
			Name:           name,
			Created:        created,
			CreatedBy:      createdBy,
			LastModified:   lastModified,
			LastModifiedBy: lastModifiedBy,
		}

		ingredients = append(ingredients, ingredient)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (repo *IngredientRepository) Create(ctx context.Context, ingredient Ingredient) (Ingredient, error) {
	query := `INSERT INTO Mixology.Mixology.Ingredients (Name, Created, CreatedBy, LastModified, LastModifiedBy) 
				OUTPUT inserted.Id
				VALUES(@name, @created, @createdBy, @lastModified, @lastModifiedBy)`

	defer repo.db.Close()

	var id int
	rows, err := repo.db.QueryContext(
		ctx,
		query,
		sql.Named("name", ingredient.Name),
		sql.Named("created", ingredient.Created),
		sql.Named("createdBy", ingredient.CreatedBy),
		sql.Named("lastModified", ingredient.LastModified),
		sql.Named("lastModifiedBy", ingredient.LastModifiedBy),
	)
	if err != nil {
		return Ingredient{}, err
	}

	for rows.Next() {
		err = rows.Scan(&id)
	}
	ingredient.Id = id

	return ingredient, nil
}
