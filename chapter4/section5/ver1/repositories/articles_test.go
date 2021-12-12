package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

// 全テスト共通の前処理を書く
func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// 前テスト共通の後処理を書く
func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  2,
			},
		}, {
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

// InsertArticle関数のテスト
func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

// UpdateNiceNum関数のテスト
func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get before data")
	}

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get after data")
	}

	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}
}
