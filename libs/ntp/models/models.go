package models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

// # User

var _ models.Model = (*User)(nil)

type User struct {
	models.BaseModel

	Email        string `json:"email" db:"email"`
	Name         string `json:"name" db:"name"`
	PasswordHash string `json:"passwordHash" db:"password_hash"`
	Description  string `json:"description" db:"description"`
	Avatar       string `json:"avatar" db:"avatar"`
}

func (m *User) TableName() string {
	return "users"
}

func UserQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&User{})
}

// # Presentation

var _ models.Model = (*Presentation)(nil)

type Presentation struct {
	models.BaseModel

	NotionPageUrl  string        `json:"notionPageUrl" db:"notion_page_url"`
	Html           string        `json:"html" db:"html"`
	CustomCss      string        `json:"customCss" db:"custom_css"`
	UserID         string        `json:"userId" db:"user_id"`
	Title          string        `json:"title" db:"title"`
	Description    string        `json:"description" db:"description"`
	Public         bool          `json:"public" db:"public"`
	Theme          string        `json:"theme" db:"theme"`
	Customizations types.JsonMap `json:"customizations" db:"customizations"`
}

func (m *Presentation) TableName() string {
	return "presentation"
}

func PresentationQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Presentation{})
}
