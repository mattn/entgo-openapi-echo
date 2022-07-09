package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mattn/entgo-openapi-echo/ent"
	"github.com/mattn/entgo-openapi-echo/ent/entry"
	_ "github.com/mattn/go-sqlite3"
)

type Api struct {
	client *ent.Client
}

type ListEntryParams struct {
	Page         *int
	ItemsPerPage *int
}

func (a *Api) ListEntry(ctx echo.Context, params ListEntryParams) error {
	page := 0
	if params.Page != nil {
		page = *params.Page
	}
	itemsPerPage := 5
	if params.ItemsPerPage != nil {
		itemsPerPage = *params.ItemsPerPage
	}
	ees, err := a.client.Entry.Query().
		Order(ent.Desc(entry.FieldCreatedAt)).
		Offset(page * itemsPerPage).
		Limit(itemsPerPage).
		All(context.Background())
	if err != nil {
		return err
	}
	return ctx.JSON(200, ees)
}

func (a *Api) CreateEntry(ctx echo.Context) error {
	var ee ent.Entry
	err := json.NewDecoder(ctx.Request().Body).Decode(&ee)
	if err != nil {
		return err
	}
	e := a.client.Entry.Create()
	e.SetContent(ee.Content)
	if !ee.CreatedAt.IsZero() {
		e.SetCreatedAt(ee.CreatedAt)
	}
	if ee2, err := e.Save(context.Background()); err != nil {
		return err
	} else {
		ee = *ee2
	}
	return ctx.JSON(200, ee)
}

func (a *Api) DeleteEntry(ctx echo.Context, id int32) error {
	e := a.client.Entry.DeleteOneID(int(id))
	return e.Exec(context.Background())
}

func (a *Api) ReadEntry(ctx echo.Context, id int32) error {
	e, err := a.client.Entry.Get(context.Background(), int(id))
	if err != nil {
		return echo.ErrNotFound
	}
	return ctx.JSON(200, e)
}

func (a *Api) UpdateEntry(ctx echo.Context, id int32) error {
	var ee ent.Entry
	err := json.NewDecoder(ctx.Request().Body).Decode(&ee)
	if err != nil {
		return err
	}
	e := a.client.Entry.UpdateOneID(int(id))
	e.SetContent(ee.Content)
	if ee2, err := e.Save(context.Background()); err != nil {
		return err
	} else {
		ee = *ee2
	}
	return ctx.JSON(200, ee)
}

func main() {
	client, err := ent.Open("sqlite3", "file:entry.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	e := echo.New()
	myApi := &Api{client: client}
	RegisterHandlers(e, myApi)
	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":8989"))
}
