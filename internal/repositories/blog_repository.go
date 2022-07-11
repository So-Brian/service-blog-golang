package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

type Blog struct {
	ID           int
	Title        string
	AuthorID     int
	Path         string
	GraphID      string
	WeChatUrl    string
	CreatedDate  time.Time
	ModifiedDate time.Time
}

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository() (*BlogRepository, error) {
	r := &BlogRepository{}

	// Build connection string
	var server = ""
	var dbPort = 0
	var user = ""
	var password = ""
	var database = ""

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, dbPort, database)

	var err error
	// Create connection pool
	r.db, err = sql.Open("sqlserver", connString)
	if err != nil {

		return nil, err
		// log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = r.db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *BlogRepository) GetBlog(id int) (*Blog, error) {
	ctx := context.Background()

	// Check if database is alive
	err := r.db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := "SELECT ID, Title FROM Blog WHERE ID = @1;"

	// Execute query
	rows, err := r.db.QueryContext(ctx, tsql, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Iterate through the result set.
	rows.Next()
	var title string

	// Get values from row.
	err = rows.Scan(&id, &title)
	if err != nil {
		return nil, err
	}

	fmt.Printf("ID: %d, Title: %s\n", id, title)

	blog := &Blog{ID: id, Title: title}
	return blog, nil
}

func (r *BlogRepository) GetBlogs() ([]Blog, error) {
	ctx := context.Background()

	// Check if database is alive
	err := r.db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := "SELECT ID, Title FROM Blog;"

	// Execute query
	rows, err := r.db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	blogs := make([]Blog, 0, 0)
	// Iterate through the result set.
	for rows.Next() {

		var id int
		var title string

		// Get values from row.
		err = rows.Scan(&id, &title)
		if err != nil {
			return nil, err
		}
		blog := &Blog{ID: id, Title: title}
		blogs = append(blogs, *blog)

	}

	return blogs, nil
}
