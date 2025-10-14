package dto

import (
	"context"

	"github.com/anonychun/bibit/internal/entity"
	"github.com/anonychun/bibit/internal/storage"
)

type AttachmentBlueprint struct {
	Id       string `json:"id"`
	FileName string `json:"file_name"`
	Url      string `json:"url"`
}

func NewAttachmentBlueprint(ctx context.Context, s3 *storage.S3, attachment *entity.Attachment) (*AttachmentBlueprint, error) {
	u, err := s3.PresignedGetObject(ctx, attachment.ObjectName)
	if err != nil {
		return nil, err
	}

	return &AttachmentBlueprint{
		Id:       attachment.Id.String(),
		FileName: attachment.FileName,
		Url:      u.String(),
	}, nil
}
