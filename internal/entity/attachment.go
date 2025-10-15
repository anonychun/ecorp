package entity

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/oklog/ulid/v2"
)

type Attachment struct {
	Base

	ObjectName string
	FileName   string
	ByteSize   int64
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
