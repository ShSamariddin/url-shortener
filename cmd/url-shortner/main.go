package url_shortner

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad("/cmd/config/local.yaml")
	fmt.Println(cfg)
}
