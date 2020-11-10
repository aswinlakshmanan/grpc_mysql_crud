package repository

import("database/sql"
		"github.com/didikprabowo/go-grpc-json-mysql/model"
		_ "github.com/go-sql-driver/mysql"
	)
type MySQL struct {
	dbCon *sql.DB
}

const dsn = "root:Hanuman98/demodb"

const NewMySQL() *MySQL {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	return &MySQL{dbCon : db}

}

func (m *MySQL) Create(ctx context.Context, a model.Article) (int64, error) {
	queryText := fmt.Sprintf("INSERT INTO articles (title,body) values ('%v','%v')", a.Title, a.Body)
	insert, err := m.dbCon.ExecContext(ctx, queryText)

	if err != nil {
		return 0, err
	}

	return insert.LastInsertId()
}

func (m *MySQL) List(ctx context.Context, start, end int64) ([]model.Article, error) {
	queryText := fmt.Sprintf("SELECT * FROM articles Order By id DESC limit %d,%d", start, end)

	rows, err := m.dbCon.QueryContext(ctx, queryText)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var articles []model.Article
	for rows.Next() {
		var article model.Article
		rows.Scan(&article.ID, &article.Title, &article.Body)
		articles = append(articles, article)
	}
	return articles, nil
}