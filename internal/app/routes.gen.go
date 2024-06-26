// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dashotv/fae"
	"github.com/dashotv/golem/plugins/router"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers = append(initializers, setupRoutes)
	healthchecks["routes"] = checkRoutes
	starters = append(starters, startRoutes)
}

func checkRoutes(app *Application) error {
	// TODO: check routes
	return nil
}

func startRoutes(ctx context.Context, app *Application) error {
	go func() {
		app.Routes()
		app.Log.Info("starting routes...")
		if err := app.Engine.Start(fmt.Sprintf(":%d", app.Config.Port)); err != nil {
			app.Log.Errorf("routes: %s", err)
		}
	}()
	return nil
}

func setupRoutes(app *Application) error {
	logger := app.Log.Named("routes").Desugar()
	e, err := router.New(logger)
	if err != nil {
		return fae.Wrap(err, "router plugin")
	}
	app.Engine = e
	// unauthenticated routes
	app.Default = app.Engine.Group("")
	// authenticated routes (if enabled, otherwise same as default)
	app.Router = app.Engine.Group("")

	// TODO: fix auth
	if app.Config.Auth {
		clerkSecret := app.Config.ClerkSecretKey
		if clerkSecret == "" {
			app.Log.Fatal("CLERK_SECRET_KEY is not set")
		}
		clerkToken := app.Config.ClerkToken
		if clerkToken == "" {
			app.Log.Fatal("CLERK_TOKEN is not set")
		}

		app.Router.Use(router.ClerkAuth(clerkSecret, clerkToken))
	}

	return nil
}

type Setting struct {
	Name  string `json:"name"`
	Value bool   `json:"value"`
}

type SettingsBatch struct {
	IDs   []string `json:"ids"`
	Name  string   `json:"name"`
	Value bool     `json:"value"`
}

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Total   int64       `json:"total,omitempty"`
}

func (a *Application) Routes() {
	a.Default.GET("/", a.indexHandler)
	a.Default.GET("/health", a.healthHandler)

	file := a.Router.Group("/file")
	file.GET("/", a.FileIndexHandler)
	file.POST("/", a.FileCreateHandler)
	file.GET("/:id", a.FileShowHandler)
	file.PUT("/:id", a.FileUpdateHandler)
	file.PATCH("/:id", a.FileSettingsHandler)
	file.DELETE("/:id", a.FileDeleteHandler)

	library := a.Router.Group("/library")
	library.GET("/", a.LibraryIndexHandler)
	library.POST("/", a.LibraryCreateHandler)
	library.GET("/:id", a.LibraryShowHandler)
	library.PUT("/:id", a.LibraryUpdateHandler)
	library.PATCH("/:id", a.LibrarySettingsHandler)
	library.DELETE("/:id", a.LibraryDeleteHandler)

	library_template := a.Router.Group("/library_template")
	library_template.GET("/", a.LibraryTemplateIndexHandler)
	library_template.POST("/", a.LibraryTemplateCreateHandler)
	library_template.GET("/:id", a.LibraryTemplateShowHandler)
	library_template.PUT("/:id", a.LibraryTemplateUpdateHandler)
	library_template.PATCH("/:id", a.LibraryTemplateSettingsHandler)
	library_template.DELETE("/:id", a.LibraryTemplateDeleteHandler)

	library_type := a.Router.Group("/library_type")
	library_type.GET("/", a.LibraryTypeIndexHandler)
	library_type.POST("/", a.LibraryTypeCreateHandler)
	library_type.GET("/:id", a.LibraryTypeShowHandler)
	library_type.PUT("/:id", a.LibraryTypeUpdateHandler)
	library_type.PATCH("/:id", a.LibraryTypeSettingsHandler)
	library_type.DELETE("/:id", a.LibraryTypeDeleteHandler)

}

func (a *Application) indexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, router.H{
		"name": "arcane",
		"routes": router.H{
			"file":             "/file",
			"library":          "/library",
			"library_template": "/library_template",
			"library_type":     "/library_type",
		},
	})
}

func (a *Application) healthHandler(c echo.Context) error {
	health, err := a.Health()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, router.H{"name": "arcane", "health": health})
}

// File (/file)
func (a *Application) FileIndexHandler(c echo.Context) error {
	page := router.QueryParamIntDefault(c, "page", "1")
	limit := router.QueryParamIntDefault(c, "limit", "25")
	return a.FileIndex(c, page, limit)
}
func (a *Application) FileCreateHandler(c echo.Context) error {
	subject := &File{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.FileCreate(c, subject)
}
func (a *Application) FileShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.FileShow(c, id)
}
func (a *Application) FileUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	subject := &File{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.FileUpdate(c, id, subject)
}
func (a *Application) FileSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	setting := &Setting{}
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.FileSettings(c, id, setting)
}
func (a *Application) FileDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.FileDelete(c, id)
}

// Library (/library)
func (a *Application) LibraryIndexHandler(c echo.Context) error {
	page := router.QueryParamIntDefault(c, "page", "1")
	limit := router.QueryParamIntDefault(c, "limit", "25")
	return a.LibraryIndex(c, page, limit)
}
func (a *Application) LibraryCreateHandler(c echo.Context) error {
	subject := &Library{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.LibraryCreate(c, subject)
}
func (a *Application) LibraryShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.LibraryShow(c, id)
}
func (a *Application) LibraryUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	subject := &Library{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.LibraryUpdate(c, id, subject)
}
func (a *Application) LibrarySettingsHandler(c echo.Context) error {
	id := c.Param("id")
	setting := &Setting{}
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.LibrarySettings(c, id, setting)
}
func (a *Application) LibraryDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.LibraryDelete(c, id)
}

// LibraryTemplate (/library_template)
func (a *Application) LibraryTemplateIndexHandler(c echo.Context) error {
	page := router.QueryParamIntDefault(c, "page", "1")
	limit := router.QueryParamIntDefault(c, "limit", "25")
	return a.LibraryTemplateIndex(c, page, limit)
}
func (a *Application) LibraryTemplateCreateHandler(c echo.Context) error {
	subject := &LibraryTemplate{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.LibraryTemplateCreate(c, subject)
}
func (a *Application) LibraryTemplateShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.LibraryTemplateShow(c, id)
}
func (a *Application) LibraryTemplateUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	subject := &LibraryTemplate{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.LibraryTemplateUpdate(c, id, subject)
}
func (a *Application) LibraryTemplateSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	setting := &Setting{}
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.LibraryTemplateSettings(c, id, setting)
}
func (a *Application) LibraryTemplateDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.LibraryTemplateDelete(c, id)
}

// LibraryType (/library_type)
func (a *Application) LibraryTypeIndexHandler(c echo.Context) error {
	page := router.QueryParamIntDefault(c, "page", "1")
	limit := router.QueryParamIntDefault(c, "limit", "25")
	return a.LibraryTypeIndex(c, page, limit)
}
func (a *Application) LibraryTypeCreateHandler(c echo.Context) error {
	subject := &LibraryType{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.LibraryTypeCreate(c, subject)
}
func (a *Application) LibraryTypeShowHandler(c echo.Context) error {
	id := c.Param("id")
	return a.LibraryTypeShow(c, id)
}
func (a *Application) LibraryTypeUpdateHandler(c echo.Context) error {
	id := c.Param("id")
	subject := &LibraryType{}
	if err := c.Bind(subject); err != nil {
		return err
	}
	return a.LibraryTypeUpdate(c, id, subject)
}
func (a *Application) LibraryTypeSettingsHandler(c echo.Context) error {
	id := c.Param("id")
	setting := &Setting{}
	if err := c.Bind(setting); err != nil {
		return err
	}
	return a.LibraryTypeSettings(c, id, setting)
}
func (a *Application) LibraryTypeDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	return a.LibraryTypeDelete(c, id)
}
