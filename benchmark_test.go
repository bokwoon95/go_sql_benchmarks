package go_sql_benchmarks

import (
	"database/sql"
	"embed"
	"flag"
	"io"
	"log"
	"testing"
	"time"

	"github.com/bokwoon95/sq"
	"github.com/bokwoon95/sqddl/ddl"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed *.sql *.csv
var embedFS embed.FS

var (
	sqliteDSN      = "test.db"
	postgresDSN    = flag.String("postgres", "", "postgres dsn")
	mysqlDSN       = flag.String("mysql", "", "mysql dsn")
	postgresDriver = "postgres"
	compiledFetch  *sq.CompiledFetch[Film]
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	for _, dsn := range []string{sqliteDSN, *postgresDSN} {
		if dsn == "" {
			continue
		}

		wipeCmd, err := ddl.WipeCommand("-db", dsn)
		if err != nil {
			log.Fatal(err)
		}
		err = wipeCmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		loadCmd, err := ddl.LoadCommand("-db", dsn, "schema.sql", "language.csv", "film.csv")
		if err != nil {
			log.Fatal(err)
		}
		loadCmd.DirFS = embedFS
		loadCmd.Stderr = io.Discard
		err = loadCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	var err error
	compiledFetch, err = sq.CompileFetch(sq.Queryf("SELECT {*} FROM film ORDER BY film_id"), func(row *sq.Row) (Film, error) {
		film := Film{
			FilmID:      row.Int("film_id"),
			Title:       row.String("title"),
			Description: row.String("description"),
			ReleaseYear: row.Int("release_year"),
			Rating:      row.String("rating"),
			LastUpdate:  row.Time("last_update"),
		}
		return film, nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Film struct {
	FilmID      int       `db:"film_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ReleaseYear int       `db:"release_year"`
	Rating      string    `db:"rating"`
	LastUpdate  time.Time `db:"last_update"`
}

func Benchmark_sq_sqlite(b *testing.B) {
	db, err := sql.Open("sqlite3", sqliteDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = sq.FetchAll(db, sq.Queryf("SELECT {*} FROM film ORDER BY film_id"), func(row *sq.Row) (Film, error) {
			film := Film{
				FilmID:      row.Int("film_id"),
				Title:       row.String("title"),
				Description: row.String("description"),
				ReleaseYear: row.Int("release_year"),
				Rating:      row.String("rating"),
				LastUpdate:  row.Time("last_update"),
			}
			return film, nil
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_sq_compiled_sqlite(b *testing.B) {
	db, err := sql.Open("sqlite3", sqliteDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = compiledFetch.FetchAll(db, sq.Params{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_sqlx_sqlite(b *testing.B) {
	db, err := sqlx.Open("sqlite3", sqliteDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	films := []Film{}
	for i := 0; i < b.N; i++ {
		err = db.Select(&films, "SELECT film_id, title, description, release_year, rating, last_update FROM film ORDER BY film_id")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_sq_postgres(b *testing.B) {
	if *postgresDSN == "" {
		b.SkipNow()
	}

	db, err := sql.Open(postgresDriver, *postgresDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = sq.FetchAll(db, sq.Queryf("SELECT {*} FROM film ORDER BY film_id"), func(row *sq.Row) (Film, error) {
			film := Film{
				FilmID:      row.Int("film_id"),
				Title:       row.String("title"),
				Description: row.String("description"),
				ReleaseYear: row.Int("release_year"),
				Rating:      row.String("rating"),
				LastUpdate:  row.Time("last_update"),
			}
			return film, nil
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_sq_compiled_postgres(b *testing.B) {
	if *postgresDSN == "" {
		b.SkipNow()
	}

	db, err := sql.Open(postgresDriver, *postgresDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = compiledFetch.FetchAll(db, sq.Params{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_sqlx_postgres(b *testing.B) {
	if *postgresDSN == "" {
		b.SkipNow()
	}

	db, err := sqlx.Open(postgresDriver, *postgresDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	films := []Film{}
	for i := 0; i < b.N; i++ {
		err = db.Select(&films, "SELECT film_id, title, description, release_year, rating, last_update FROM film ORDER BY film_id")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_sq_mysql(b *testing.B) {
	if *mysqlDSN == "" {
		b.SkipNow()
	}

	db, err := sql.Open("mysql", *mysqlDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = sq.FetchAll(db, sq.Queryf("SELECT {*} FROM film ORDER BY film_id"), func(row *sq.Row) (Film, error) {
			film := Film{
				FilmID:      row.Int("film_id"),
				Title:       row.String("title"),
				Description: row.String("description"),
				ReleaseYear: row.Int("release_year"),
				Rating:      row.String("rating"),
				LastUpdate:  row.Time("last_update"),
			}
			return film, nil
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_sq_compiled_mysql(b *testing.B) {
	if *mysqlDSN == "" {
		b.SkipNow()
	}

	db, err := sql.Open("mysql", *mysqlDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = compiledFetch.FetchAll(db, sq.Params{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_sqlx_mysql(b *testing.B) {
	if *mysqlDSN == "" {
		b.SkipNow()
	}

	db, err := sqlx.Open("mysql", *mysqlDSN)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	films := []Film{}
	for i := 0; i < b.N; i++ {
		err = db.Select(&films, "SELECT film_id, title, description, release_year, rating, last_update FROM film ORDER BY film_id")
		if err != nil {
			b.Fatal(err)
		}
	}
}
