package dbconnect_test

import (
	"testing"

	config "github.com/kevinkimutai/metadata"
	"github.com/kevinkimutai/metadata/internal/adapter/db/db"
	dbconnect "github.com/kevinkimutai/metadata/internal/adapter/db/dbConnect"
	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {

	// Define test database URL
	testDBUrl := config.DATABASE_URL

	// Initialize a new DBAdapter
	adapter := dbconnect.NewDB(testDBUrl)

	// Test database initialization
	if adapter == nil {
		t.Fatal("DBAdapter initialization failed")
	}

}

func TestGetMovieById(t *testing.T) {

	testDBUrl := config.DATABASE_URL

	// Initialize a new DBAdapter
	adapter := dbconnect.NewDB(testDBUrl)

	// Test database initialization
	if adapter == nil {
		t.Fatal("DBAdapter initialization failed")
	}

	//Check no movie with ID
	movieID := int64(0)
	movie, err := adapter.GetMovieById(movieID)

	if err != nil {
		t.Fatalf("Error getting movie by id:%v", err)
	}

	assert.Equal(t, movie, db.Movie{})

	//Check Valid ID
	movieID = int64(1)
	movie, err = adapter.GetMovieById(movieID)

	if err != nil {
		t.Fatalf("Error getting movie by id:%v", err)
	}

	assert.Equal(t, movie.Title, "Dark Knight 3")

}
