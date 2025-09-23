package internal

import (
	"fmt"
	"path/filepath"
)

func GenerateEntity(name string) error {
	data := TemplateData{PackageName: "entity"}
	filePath := filepath.Join("internal/entity", fmt.Sprintf("%s.go", name))

	return generateFile(filePath, emptyTemplate, data)
}
