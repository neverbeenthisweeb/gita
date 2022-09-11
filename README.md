# Gita

Prometheus middleware for Gin

## Installation

1. Install Go (**^1.17**), then install package

```
go get -u github.com/neverbeenthisweeb/gita
```

2. Import in your code

```
import (
	"github.com/gin-gonic/gin"
	"github.com/neverbeenthisweeb/gita"
)

func main() {
	r := gin.Default()
	r.Use(gita.HandleFunc())

    // Define routes

	r.Run(":8080")
}
```
