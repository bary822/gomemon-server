package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/bary822/gomemon-server/internal/entity"
)

const FilePath = "./memos.json"

type MemoFileRepository struct {
	memos []*entity.Memo
}

func NewMemoFileRepository() *MemoFileRepository {
	return &MemoFileRepository{
		memos: make([]*entity.Memo, 0),
	}
}

func (repo *MemoFileRepository) Save(memo entity.Memo) (*entity.Memo, error) {
	repo.load()

	repo.memos = append(repo.memos, &memo)
	repo.save()

	return &memo, nil
}

func (repo *MemoFileRepository) GetByID(id string) (*entity.Memo, error) {
	repo.load()

	for _, memo := range repo.memos {
		if memo.ID == id {
			return memo, nil
		}
	}

	return nil, errors.New("Memo not found")
}

func (repo *MemoFileRepository) GetAll() ([]*entity.Memo, error) {
	repo.load()
	return repo.memos, nil
}

func (repo *MemoFileRepository) save() error {
	f, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := repo.marshal()
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func (repo *MemoFileRepository) load() error {
	f, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()

	// Initialize file content if it does not exist.
	fi, err := f.Stat()
	if err != nil {
		log.Fatal("Failed to examine file: " + err.Error())
	}
	if fi.Size() == 0 {
		repo.save()
	}

	if err != nil {
		log.Fatal("Failed to open file: " + err.Error())
	}

	repo.ummarshal(f)

	return err
}

func (repo *MemoFileRepository) marshal() (io.Reader, error) {
	b, err := json.MarshalIndent(repo.memos, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (repo *MemoFileRepository) ummarshal(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(&repo.memos)

	if err != nil {
		log.Fatal("Failed to decode json: " + err.Error())
		return err
	}

	return err
}
