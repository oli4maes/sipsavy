package relational

import (
	"database/sql"
	"log"
	"time"
)

type CocktailRepository struct {
	db *sql.DB
}

func NewCocktailRepository(connString string) CocktailRepository {
	sqlDb, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("could not create connection: %s", err)
	}

	return CocktailRepository{db: sqlDb}
}

type Cocktail struct {
	Id             int
	Name           string
	Created        time.Time
	CreatedBy      string
	LastModified   time.Time
	LastModifiedBy string
	Ingredients    []CocktailIngredient
}

type CocktailIngredient struct {
	CocktailId     int
	IngredientId   int
	Amount         int
	IngredientUnit string
}

func (repo *CocktailRepository) GetAll() ([]Cocktail, error) {
	var cocktails []Cocktail

	query := `SELECT * FROM Mixology.Mixology.Cocktails`

	rows, err := repo.db.Query(query)
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

		var cocktail = Cocktail{
			Id:             id,
			Name:           name,
			Created:        created,
			CreatedBy:      createdBy,
			LastModified:   lastModified,
			LastModifiedBy: lastModifiedBy,
		}

		cocktails = append(cocktails, cocktail)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cocktails, nil
}

func (repo *CocktailRepository) CreateCocktail(cocktail Cocktail, ingredients []CocktailIngredient) (Cocktail, error) {
	return Cocktail{}, nil
}
