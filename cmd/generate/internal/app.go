package internal

import (
	"os"
	"path/filepath"
)

func GenerateApp(name string) error {
	targetDir := filepath.Join("internal/app", name)
	err := os.MkdirAll(targetDir, os.ModePerm)
	if err != nil {
		return err
	}

	data := TemplateData{
		ModuleName:  getModuleName(),
		PackageName: extractPackageName(name),
	}

	err = generateFile(filepath.Join(targetDir, "init.go"), appInitTemplate, data)
	if err != nil {
		return err
	}

	err = generateFile(filepath.Join(targetDir, "dto.go"), emptyTemplate, data)
	if err != nil {
		return err
	}

	err = generateFile(filepath.Join(targetDir, "handler.go"), emptyTemplate, data)
	if err != nil {
		return err
	}

	err = generateFile(filepath.Join(targetDir, "usecase.go"), emptyTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

const appInitTemplate = `package {{.PackageName}}

import (
	"{{.ModuleName}}/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{}, nil
}

type Handler struct {
	usecase *Usecase
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		usecase: do.MustInvoke[*Usecase](i),
	}, nil
}
`
