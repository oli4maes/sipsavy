package relational

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"time"
)

type CocktailRepository struct {
	db *gorm.DB
}

func NewCocktailRepository(connString string) CocktailRepository {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not create connection: %s", err)
	}

	return CocktailRepository{db: db}
}

type Cocktail struct {
	CocktailId     uuid.UUID `gorm:"primaryKey"`
	Name           string
	Created        time.Time
	CreatedBy      string
	LastModified   time.Time
	LastModifiedBy string
	Ingredients    []CocktailIngredient
}

type CocktailIngredient struct {
	CocktailId     uuid.UUID `gorm:"primaryKey"`
	IngredientId   uuid.UUID `gorm:"primaryKey"`
	Amount         int
	IngredientUnit string
}

func (repo *CocktailRepository) GetAll(ctx context.Context) ([]Cocktail, error) {
	var cocktails []Cocktail
	result := repo.db.Find(&cocktails)
	if result.Error != nil {
		return nil, result.Error
	}

	return cocktails, nil
}

func (repo *CocktailRepository) Create(ctx context.Context, cocktail Cocktail) (Cocktail, error) {
	cocktail.CocktailId = uuid.New()
	result := repo.db.Create(cocktail)
	if result.Error != nil {
		return Cocktail{}, result.Error
	}

	return cocktail, nil
}

func (repo *CocktailRepository) GetByIngredientIds(ctx context.Context, ingredientIds []int) ([]Cocktail, error) {
	var cocktails []Cocktail

	// TODO: implement this func

	return cocktails, nil
}
