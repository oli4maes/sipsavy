package relational

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type IngredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository() IngredientRepository {
	connString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		panic("connection string env variable not set")
	}

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not create connection: %s", err)
	}

	return IngredientRepository{db: db}
}

type Ingredient struct {
	Id             uuid.UUID `gorm:"primaryKey"`
	Name           string    `gorm:"not null"`
	Created        time.Time `gorm:"not null"`
	CreatedBy      string    `gorm:"not null"`
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
	ingredient.Id = uuid.New()
	result := repo.db.Create(&ingredient)
	if result.Error != nil {
		return Ingredient{}, result.Error
	}

	return ingredient, nil
}
