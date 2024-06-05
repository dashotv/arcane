// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	"github.com/dashotv/grimoire"
	"github.com/kamva/mgm/v3"
)

func init() {
	initializers = append(initializers, setupDb)
	healthchecks["db"] = checkDb
}

func setupDb(app *Application) error {
	db, err := NewConnector(app)
	if err != nil {
		return err
	}

	app.DB = db
	return nil
}

func checkDb(app *Application) (err error) {
	// TODO: Check DB connection
	return nil
}

type Connector struct {
	Log             *zap.SugaredLogger
	File            *grimoire.Store[*File]
	Library         *grimoire.Store[*Library]
	LibraryTemplate *grimoire.Store[*LibraryTemplate]
	LibraryType     *grimoire.Store[*LibraryType]
}

func connection[T mgm.Model](name string) (*grimoire.Store[T], error) {
	s, err := app.Config.ConnectionFor(name)
	if err != nil {
		return nil, err
	}
	c, err := grimoire.New[T](s.URI, s.Database, s.Collection)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewConnector(app *Application) (*Connector, error) {
	file, err := connection[*File]("file")
	if err != nil {
		return nil, err
	}

	grimoire.Indexes[*File](file, &File{})

	library, err := connection[*Library]("library")
	if err != nil {
		return nil, err
	}

	grimoire.Indexes[*Library](library, &Library{})

	library_template, err := connection[*LibraryTemplate]("library_template")
	if err != nil {
		return nil, err
	}

	grimoire.Indexes[*LibraryTemplate](library_template, &LibraryTemplate{})

	library_type, err := connection[*LibraryType]("library_type")
	if err != nil {
		return nil, err
	}

	grimoire.Indexes[*LibraryType](library_type, &LibraryType{})

	c := &Connector{
		Log:             app.Log.Named("db"),
		File:            file,
		Library:         library,
		LibraryTemplate: library_template,
		LibraryType:     library_type,
	}

	return c, nil
}

type File struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Type       string             `bson:"type" json:"type" grimoire:"index"`
	Path       string             `bson:"path" json:"path" grimoire:"index"`
	Size       int64              `bson:"size" json:"size"`
	ModifiedAt int64              `bson:"modified_at" json:"modified_at"`
	MediumID   primitive.ObjectID `bson:"medium_id" json:"medium_id" grimoire:"index"`
}

type Library struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Name              string             `bson:"name" json:"name" grimoire:"index"`
	Path              string             `bson:"path" json:"path" grimoire:"index"`
	LibraryTypeID     primitive.ObjectID `bson:"library_type_id" json:"library_type_id"`
	LibraryTemplateID primitive.ObjectID `bson:"library_template_id" json:"library_template_id"`
	LibraryType       *LibraryType       `bson:"-" json:"library_type"`
	LibraryTemplate   *LibraryTemplate   `bson:"-" json:"library_template"`
}

type LibraryTemplate struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Name     string `bson:"name" json:"name"`
	Template string `bson:"template" json:"template"`
}

type LibraryType struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Name string `bson:"name" json:"name"`
}