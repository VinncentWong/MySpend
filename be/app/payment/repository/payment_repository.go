package repository

import (
	"context"
	"fmt"
	"mime/multipart"
	"module/domain"
	"module/infrastructure"
	"strconv"

	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentRepository struct {
	db          *gorm.DB
	minioClient *minio.Client
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		db:          infrastructure.GetDb(),
		minioClient: infrastructure.GetClient(),
	}
}

func (r *PaymentRepository) CreatePaymentHistory(id string, payment domain.PaymentHistory, file *multipart.FileHeader) (*domain.PaymentHistory, error) {
	var user domain.User
	result := r.db.
		Model(&domain.User{}).
		Preload(clause.Associations).
		Where("id = ?", id).
		Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	payment.LinkPaymentProof = fmt.Sprintf("%s/%d", file.Filename, user.ID)
	r.minioClient.FPutObject(context.Background(),
		"payment_proof",
		payment.LinkPaymentProof,
		fmt.Sprintf("/tmp/%s", payment.LinkPaymentProof),
		minio.PutObjectOptions{
			ContentType: file.Header["Content-Type"][0],
		},
	)
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	payment.User = user
	payment.UserID = uint(idNumber)
	result = r.db.Create(&payment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &payment, nil
}
