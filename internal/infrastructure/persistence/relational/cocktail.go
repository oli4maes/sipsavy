package relational

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/microsoft/go-mssqldb/integratedauth/krb5"
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

func (repo *CocktailRepository) GetAll(ctx context.Context) ([]Cocktail, error) {
	var cocktails []Cocktail

	query := `SELECT * FROM Mixology.Mixology.Cocktails`

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

func (repo *CocktailRepository) Create(ctx context.Context, cocktail Cocktail) (Cocktail, error) {
	cocktailQuery := `INSERT INTO Mixology.Mixology.Cocktails (Name, Created, CreatedBy, LastModified, LastModifiedBy)	
						OUTPUT inserted.Id
						VALUES (@name, @created, @createdBy, @lastModified, @lastModifiedBy)`

	ingredientsQuery := `INSERT INTO Mixology.Mixology.CocktailIngredients (CocktailId, IngredientId, Amount, IngredientUnit) 
							VALUES(@cocktailId, @ingredientId, @amount, @ingredientUnit)`

	defer repo.db.Close()

	var id int
	rows, err := repo.db.QueryContext(
		ctx,
		cocktailQuery,
		sql.Named("name", cocktail.Name),
		sql.Named("created", cocktail.Created),
		sql.Named("createdBy", cocktail.CreatedBy),
		sql.Named("lastModified", cocktail.LastModified),
		sql.Named("lastModifiedBy", cocktail.LastModifiedBy))
	if err != nil {
		return Cocktail{}, err
	}

	for rows.Next() {
		err = rows.Scan(&id)
	}
	cocktail.Id = id

	for _, i := range cocktail.Ingredients {
		_, err := repo.db.QueryContext(
			ctx,
			ingredientsQuery,
			sql.Named("cocktailId", cocktail.Id),
			sql.Named("ingredientId", i.IngredientId),
			sql.Named("amount", i.Amount),
			sql.Named("ingredientUnit", i.IngredientUnit),
		)
		if err != nil {
			return Cocktail{}, err
		}
	}

	return cocktail, nil
}

func (repo *CocktailRepository) GetByIngredientIds(ctx context.Context, ingredientIds []int) ([]Cocktail, error) {
	var cocktails []Cocktail

	ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ingredientIds)),
		","), "[]")

	query := fmt.Sprintf(
		"SELECT c.Id, c.Name FROM Mixology.Mixology.Cocktails c "+
			"JOIN Mixology.Mixology.CocktailIngredients ci ON c.Id = ci.CocktailId "+
			"WHERE ci.IngredientId IN (%s) "+
			"GROUP BY c.Id, c.Name", ids)

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer repo.db.Close()

	var id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		var cocktail = Cocktail{
			Id:   id,
			Name: name,
		}

		cocktails = append(cocktails, cocktail)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cocktails, nil
}
