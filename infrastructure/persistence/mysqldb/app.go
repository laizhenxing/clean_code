package mysqldb

import (
	"database/sql"
	"fmt"
	"time"

	"cc/domain/entity"
)

// app 的操作实例
type AppHandler struct {
	dbConn *sql.DB
}

func NewAppHandler(conn *sql.DB) *AppHandler {
	return &AppHandler{conn}
}

func (appHandler *AppHandler) FindAll() ([]*entity.App, error) {
	var apps []*entity.App
	rows, err := appHandler.dbConn.Query("SELECT * FROM apps")
	if err != nil {
		return apps, err
	}

	for rows.Next() {
		var app entity.App
		err := rows.Scan(&app.ID, &app.Name, &app.Version, &app.CreatedAt, &app.UpdatedAt)
		if err != nil {
			return nil, err
		}
		apps = append(apps, &app)
	}

	return apps, nil
}

func (appHandler *AppHandler) Save(app entity.App) error {
	stmt, err := appHandler.dbConn.Prepare("INSERT INTO apps(`name`, `version`) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(app.Name, app.Version)
	if err != nil {
		return err
	}
	return nil
}

func (appHandler *AppHandler) FindOne(id uint64) (*entity.App, error) {
	var app entity.App
	err := appHandler.dbConn.QueryRow("SELECT * FROM apps WHERE id = ?", id).Scan(&app.ID, &app.Name, &app.Version, &app.CreatedAt, &app.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (appHandler *AppHandler) Delete(id uint64) error {
	_, err := appHandler.dbConn.Exec("DELETE from apps WHERE id = ?", id)
	return err
}

func (appHandler *AppHandler) Update(app entity.App) error {
	setter := ""
	if app.Name != "" {
		setter = fmt.Sprintf("%s `name`='%s'", setter, app.Name)
	}
	if app.Version != "" {
		setter = fmt.Sprintf("%s, `version`='%s'", setter, app.Version)
	}
	// 无需更新
	if len(setter) == 0 {
		return nil
	}
	setter = fmt.Sprintf("%s, updated_at='%s'", setter, time.Now().Format("2006-01-02 15:04:05"))

	stmt, err := appHandler.dbConn.Prepare("UPDATE apps SET " + setter + " WHERE `id` = ?")

	if err != nil {
		return err
	}
	_, err = stmt.Exec(app.ID)
	return err
}
