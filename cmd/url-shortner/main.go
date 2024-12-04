package url_shortner

import (
	"fmt"
	"url-shortener/cmd/internal/config"
)

func main() {
	cfg := config.MustLoad("/cmd/config/local.yaml")
	fmt.Println(cfg)
}
