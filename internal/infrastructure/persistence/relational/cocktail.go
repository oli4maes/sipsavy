package relational

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
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
	Id             uuid.UUID `gorm:"primaryKey"`
	Name           string    `gorm:"not null"`
	Created        time.Time `gorm:"not null"`
	CreatedBy      string    `gorm:"not null"`
	LastModified   time.Time
	LastModifiedBy string
	Ingredients    []Ingredient `gorm:"many2many:cocktail_ingredients;"`
}

type CocktailIngredient struct {
	CocktailId     uuid.UUID `gorm:"primaryKey"`
	Cocktail       Cocktail
	IngredientId   uuid.UUID `gorm:"primaryKey"`
	Ingredient     Ingredient
	Amount         int    `gorm:"not null"`
	IngredientUnit string `gorm:"not null"`
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
	cocktail.Id = uuid.New()
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
