package template

import (
	"embed"
	"github.com/joho/godotenv"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed web
var embededFiles embed.FS

//go:embed web/assets
var embedAssets embed.FS

func init() {
	_ = godotenv.Load(".env")
}

func GetFileSystem() http.FileSystem {
	useOS := os.Getenv("TEMPLATE_LIVE") == "true"

	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("web"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "web")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func GetAssets() http.FileSystem {
	useOS := os.Getenv("TEMPLATE_LIVE") == "true"

	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("web/assets"))
	}

	log.Print("using embed mode")
	fSys, err := fs.Sub(embedAssets, "web/assets")
	if err != nil {
		panic(err)
	}

	return http.FS(fSys)
}
