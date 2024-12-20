package repository

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sync"
)

type CollectionTmpFile struct{}

func NewCollectionTmpFile() *CollectionTmpFile {
	return &CollectionTmpFile{}
}

const (
	environment = "environment"
	collection  = "collection"
)

func (c CollectionTmpFile) SaveFile(ctx context.Context, collectionFilePath, envFilePath string, collectionFile, environmentFile multipart.File) error {
	await := sync.WaitGroup{}
	mapErrors := sync.Map{}

	await.Add(2)

	go func() {
		defer await.Done()
		if err := c.saveFile(collectionFilePath, collectionFile); err != nil {
			fmt.Errorf("unable to save collectionFile")
			mapErrors.Store(collection, err)
		}

	}()

	go func() {
		defer await.Done()
		if err := c.saveFile(envFilePath, environmentFile); err != nil {
			fmt.Errorf("unable to save environmentFile")
			mapErrors.Store(environment, err)
		}
	}()

	await.Wait()

	var (
		msgError string
		idx      int
	)

	mapErrors.Range(func(key, value interface{}) bool {
		idx++
		msgError += fmt.Sprintf("error_%d: %s", idx, value.(error).Error())
		return true
	})

	return nil

}

func (c CollectionTmpFile) UnlinkFiles(ctx context.Context, files []string) error {
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("error removing file: %w", err)
		}
	}

	return nil
}

func (c CollectionTmpFile) saveFile(path string, file multipart.File) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
