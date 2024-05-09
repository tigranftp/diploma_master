package api

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var allModelsPath = filepath.Join(os.ExpandEnv("$AppData"), "nn-auto", "models")

func (a *API) createModelHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseMultipartForm(1000000000); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		modelPath := filepath.Join(allModelsPath, uuid.NewString(), "custom")
		err := os.MkdirAll(modelPath, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		script := req.MultipartForm.File["script"][0]
		files := req.MultipartForm.File["files"]
		files = append(files)
		for _, fileHeader := range files {
			func() {
				err = copyFile(fileHeader, modelPath)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "Файл %s успешно загружен\n", fileHeader.Filename)
			}()
		}
		err = copyFile(script, modelPath, "__init__.py")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Файл %s успешно загружен\n", script.Filename)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}
}

func copyFile(fileHeader *multipart.FileHeader, modelPath string, specificPath ...string) (err error) {
	// Открываем файл
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Создаем новый файл для сохранения загруженного файла
	fileName := fileHeader.Filename
	if len(specificPath) > 0 {
		fileName = specificPath[0]
	}
	newFile, err := os.Create(filepath.Join(modelPath, fileName))
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Копируем содержимое загруженного файла в новый файл
	_, err = io.Copy(newFile, file)
	return err
}
