package pb_migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("01si6l3omlqheru")
		if err != nil {
			return err
		}

		// update
		edit_customizations := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "elrsyizz",
			"name": "customizations",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 250000
			}
		}`), edit_customizations); err != nil {
			return err
		}
		collection.Schema.AddField(edit_customizations)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("01si6l3omlqheru")
		if err != nil {
			return err
		}

		// update
		edit_customizations := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "elrsyizz",
			"name": "customizations",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 2000000
			}
		}`), edit_customizations); err != nil {
			return err
		}
		collection.Schema.AddField(edit_customizations)

		return dao.SaveCollection(collection)
	})
}
