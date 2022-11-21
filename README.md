# go_docker_sample

## Start with...
`go mod init`

or

`go mod init github.com.YOUR_REPOSITORY/REPOSITORY_NAME`


## Creating main.go
Just for checking, make a simple page:
```
package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}


func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run the file and Check if the connection works.
```
go run main.go
```

## Migration file

```
mkdir database
```
```
touch migration.sql
```


## docker へのインストール
```
docker-compose exec go ash
```
``
go get github.com/gin-gonic/gin
```

```
go get github.com/jinzhu/gorm
```
