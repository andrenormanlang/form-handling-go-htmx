package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andrenormanlang/common"
)

type Database struct {
	Address string
	Port int
	User string
	Connection *sql.DB
}

/// This function gets all the posts from the current
/// database connection.
func (db Database) GetPosts() ([]common.Post, error) {
	rows, err := db.Connection.Query("SELECT title, content FROM posts")
	if err != nil {
		return make([]common.Post, 0), err
	}

	all_posts := make([]common.Post, 0)
	for rows.Next() {
		var post common.Post
		if err = rows.Scan(&post.Title, &post.Content); err != nil {
			return make([]common.Post, 0), err;
		}
		all_posts = append(all_posts, post)
	}

	return all_posts, nil
}

func MakeSqlConnection(user string, password string, address string, port int) (Database, error) {
	/// Checking the DB connection
	connection_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/gocms", user, password, address, port)
	db, err := sql.Open("mysql", connection_str)
	if err != nil {
		return Database{}, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return Database{
		Address: address,
		Port: port,
		User: user,
		Connection: db,
	}, nil
}