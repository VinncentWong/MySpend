package repository

import (
	"context"
	"mime/multipart"
	"module/domain"
	"module/infrastructure"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPaymentRepository interface {
	CreatePaymentHistory(id string, payment domain.PaymentHistory, file *multipart.FileHeader, month uint) (*domain.PaymentHistory, error)
}

type PaymentRepository struct {
	db          *gorm.DB
	minioClient *minio.Client
}

func NewPaymentRepository() IPaymentRepository {
	return &PaymentRepository{
		db:          infrastructure.GetDb(),
		minioClient: infrastructure.GetClient(),
	}
}

func (r *PaymentRepository) CreatePaymentHistory(id string, payment domain.PaymentHistory, file *multipart.FileHeader, month uint) (*domain.PaymentHistory, error) {
	var user domain.User
	result := r.db.
		Model(&domain.User{}).
		Preload(clause.Associations).
		Where("id = ?", id).
		Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	payment.ID = uuid.New()
	payment.LinkPaymentProof = file.Filename
	user.IsPremium = true
	user.PremiumExpire = time.Now().AddDate(0, int(month), 0)

	// get buffer from opened file
	buffer, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer buffer.Close()

	fileBuffer := buffer
	if err != nil {
		return nil, err
	}
	_, err = r.minioClient.PutObject(context.Background(),
		"paymentproof",
		payment.LinkPaymentProof,
		fileBuffer,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header["Content-Type"][0],
		},
	)
	if err != nil {
		return nil, err
	}
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
	result = r.db.Updates(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &payment, nil
}
