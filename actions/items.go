package actions

import (

  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop/v5"
  "github.com/gobuffalo/x/responder"
  "github.com/todos-api/coke/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Item)
// DB Table: Plural (items)
// Resource: Plural (Items)
// Path: Plural (/items)
// View Template Folder: Plural (/templates/items/)

// ItemsResource is the resource for the Item model
type ItemsResource struct{
  buffalo.Resource
}

// List gets all Items. This function is mapped to the path
// GET /items
func (v ItemsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  items := &models.Items{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Items from the DB
  if err := q.All(items); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    c.Set("items", items)
    return c.Render(http.StatusOK, r.HTML("/items/index.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(items))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(items))
  }).Respond(c)
}

// Show gets the data for one Item. This function is mapped to
// the path GET /items/{item_id}
func (v ItemsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Item
  item := &models.Item{}

  // To find the Item the parameter item_id is used.
  if err := tx.Find(item, c.Param("item_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    c.Set("item", item)

    return c.Render(http.StatusOK, r.HTML("/items/show.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(item))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(item))
  }).Respond(c)
}

// Create adds a Item to the DB. This function is mapped to the
// path POST /items
func (v ItemsResource) Create(c buffalo.Context) error {
  // Allocate an empty Item
  item := &models.Item{}

  // Bind item to the html form elements
  if err := c.Bind(item); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(item)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the new.html template that the user can
      // correct the input.
      c.Set("item", item)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/items/new.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "item.created.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/items/%v", item.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.JSON(item))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.XML(item))
  }).Respond(c)
}

// Update changes a Item in the DB. This function is mapped to
// the path PUT /items/{item_id}
func (v ItemsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Item
  item := &models.Item{}

  if err := tx.Find(item, c.Param("item_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Item to the html form elements
  if err := c.Bind(item); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(item)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the edit.html template that the user can
      // correct the input.
      c.Set("item", item)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/items/edit.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "item.updated.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/items/%v", item.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(item))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(item))
  }).Respond(c)
}

// Destroy deletes a Item from the DB. This function is mapped
// to the path DELETE /items/{item_id}
func (v ItemsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Item
  item := &models.Item{}

  // To find the Item the parameter item_id is used.
  if err := tx.Find(item, c.Param("item_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(item); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a flash message
    c.Flash().Add("success", T.Translate(c, "item.destroyed.success"))

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "/items")
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(item))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(item))
  }).Respond(c)
}
