package relational

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"time"
)

type IngredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(connString string) IngredientRepository {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not create connection: %s", err)
	}

	return IngredientRepository{db: db}
}

type Ingredient struct {
	IngredientId   uuid.UUID `gorm:"primaryKey"`
	Name           string
	Created        time.Time
	CreatedBy      string
	LastModified   time.Time
	LastModifiedBy string
}

func (repo *IngredientRepository) GetAll(ctx context.Context) ([]Ingredient, error) {
	var ingredients []Ingredient
	result := repo.db.Find(&ingredients)
	if result.Error != nil {
		return nil, result.Error
	}

	return ingredients, nil
}

func (repo *IngredientRepository) Create(ctx context.Context, ingredient Ingredient) (Ingredient, error) {
	ingredient.IngredientId = uuid.New()
	result := repo.db.Create(&ingredient)
	if result.Error != nil {
		return Ingredient{}, result.Error
	}

	return ingredient, nil
}
