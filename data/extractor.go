package data

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ExtractCSVFromZIP extracts the CSV file from a ZIP archive
func ExtractCSVFromZIP(zipFilePath, csvFilePath string) error {
	zipFile, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	for _, file := range zipFile.File {
		if filepath.Ext(file.Name) == ".CSV" || filepath.Ext(file.Name) == ".csv" {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			dst, err := os.Create(csvFilePath)
			if err != nil {
				return err
			}
			defer dst.Close()

			_, err = io.Copy(dst, src)
			if err != nil {
				return err
			}

			fmt.Printf("CSV file extracted: %s\n", csvFilePath)
			return nil
		}
	}

	return fmt.Errorf("CSV file not found in the ZIP archive")
}
