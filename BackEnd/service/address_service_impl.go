package service

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
)

type AddressServiceImpl struct {
	AddressRepository repository.AddressRepository
	UserRepository    repository.UserRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewAddressService(addressRepository repository.AddressRepository, useRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AddressService {
	return &AddressServiceImpl{
		AddressRepository: addressRepository,
		UserRepository:    useRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *AddressServiceImpl) AddAddress(ctx context.Context, request web.AddressCreateRequest, userId int) web.AddressResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	address := domain.Address{
		UserId:       user.Id,
		NamaPenerima: request.NamaPenerima,
		Kabupaten:    request.Kabupaten,
		Kecamatan:    request.Kecamatan,
		Kelurahan:    request.Kelurahan,
		Alamat:       request.Alamat,
		NoHP:         request.NoHP,
		Note:         request.Note,
	}

	address = service.AddressRepository.AddAddress(ctx, tx, address)

	return helper.ToAddressResponse(address)
}

func (service *AddressServiceImpl) UpdateAddress(ctx context.Context, request web.AddressUpdateRequest, userId int) web.AddressResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	address := domain.Address{
		Id:           request.Id,
		UserId:       user.Id,
		NamaPenerima: request.NamaPenerima,
		Kabupaten:    request.Kabupaten,
		Kecamatan:    request.Kecamatan,
		Kelurahan:    request.Kelurahan,
		Alamat:       request.Alamat,
		NoHP:         request.NoHP,
		Note:         request.Note,
	}

	address = service.AddressRepository.UpdateAddress(ctx, tx, address)

	return helper.ToAddressResponse(address)
}

func (service *AddressServiceImpl) FindByUserId(ctx context.Context, userId int) []web.AddressResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	addresses, err := service.AddressRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAddressResponses(addresses)
}
