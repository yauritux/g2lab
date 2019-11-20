package main

import (
  "fmt"
  "strconv"
  "log"
  "io/ioutil"
  "encoding/json"
  "net/http"

  "github.com/labstack/echo"
)

type product struct {
  Name string `json:"name"`
  Stock int `json:"stock"`
  Price float32 `json:"price"`
}

func index(c echo.Context) error {
  return c.String(http.StatusOK, "OK")
}

func getProducts(c echo.Context) error {
  name := c.QueryParam("name")
  stock, err := strconv.Atoi(c.QueryParam("stock"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot proceed data. Invalid type for stock!"})
  }
  price, err := strconv.ParseFloat(c.QueryParam("price"), 32)
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot proceed data. Invalid type for price!"})
  }

  switch c.Param("format") {
    case "string": return c.String(
      http.StatusOK, 
      fmt.Sprintf(
        "Product Name: %v, Stock: %d, Price: %.3g\n", name, stock, price,
      ))
    case "json": return c.JSON(
      http.StatusOK, map[string]interface{}{
        "name": name,
        "stock": stock,
        "price": price,
      })
  }


  return c.JSON(
    http.StatusBadRequest,
    map[string]string{
      "error": "You need to let us know which data type you'd like to see. Please provide either 'string' or 'json'",
    },
  )
}

func addProduct(c echo.Context) error {
  newProduct := product{}

  defer c.Request().Body.Close()

  b, err := ioutil.ReadAll(c.Request().Body)
  if err != nil {
    log.Printf("Failed reading the request body: %s", err)
    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
  }

  err = json.Unmarshal(b, &newProduct)
  if err != nil {
    log.Printf("Failed unmarshaling product payload: %s", err)
    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
  }

  log.Printf("Successfully added product %#v", newProduct)
  return c.String(http.StatusCreated, fmt.Sprintf("We've got your new product: %v\n", newProduct))
}

func mainAdmin(c echo.Context) error {
  return c.String(http.StatusOK, "Horray,,, you are on the secret admin page!")
}

func main() {
  e := echo.New()

  g := e.Group("/admin")

  g.GET("/main", mainAdmin)

  e.GET("/", index)
  e.GET("/products/:format", getProducts)
  e.POST("/products", addProduct)

  e.Logger.Fatal(e.Start(":8000"))
}
