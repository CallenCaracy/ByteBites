package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
    var err error
    connStr := "postgresql://postgres.jwltksifgtedjisubvxa:QLOVRo32Wv2eiHFp@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("error opening DB: %v", err)
    }

    // Test the connection
    if err := DB.Ping(); err != nil {
        return fmt.Errorf("error pinging DB: %v", err)
    }
    return nil
}
