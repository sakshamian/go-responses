# go API common response struct

A package for handling most common success responses in golang

## Installation

```
go get github.com/sakshamian/go-responses   
```

## Usage

This package makes response handling in go easy. Just write the type of response, data(if it returns one) and custom message for consumption on frontend.
```
http.HandleFunc("/health", func (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  // Write response as JSON
  json.NewEncoder(w).Encode(responsepkg.ResponseSuccess("Health OK!", nil))
})
```

Or when using with gin:
```
r.GET("/health", func(c *gin.Context) {
  c.JSON(http.StatusOK, responsepkg.ResponseSuccess("Health OK!", nil))
})
```

## Methods

| S. No. | Method name                     | Parameters(required)      | Parameters(optional)
| :-:    | ------------------------------- | ------------------------- | --------------------|
| 01     | ResponseSuccess                 | Response message, data(can be nil) |                     |
| 02     | ResponseCreated                 | Response message, data(can be nil) |        -            |
| 03     | ResponseNoContent               | Response message          |       data          |
