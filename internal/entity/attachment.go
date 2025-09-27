package entity

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Attachment struct {
	Base

	ObjectName string
	FileName   string
	ByteSize   int64
}

func (a *Attachment) BeforeUpdate(tx *gorm.DB) error {
	a.UpdatedAt = time.Now()
	return nil
}

func NewAttachmentFromFile(file *os.File) (*Attachment, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &Attachment{
		ObjectName: ulid.Make().String() + filepath.Ext(fileInfo.Name()),
		FileName:   fileInfo.Name(),
		ByteSize:   fileInfo.Size(),
	}, nil
}

func NewAttachmentFromFileHeader(fileHeader *multipart.FileHeader) *Attachment {
	return &Attachment{
		ObjectName: ulid.Make().String() + filepath.Ext(fileHeader.Filename),
		FileName:   fileHeader.Filename,
		ByteSize:   fileHeader.Size,
	}
}
