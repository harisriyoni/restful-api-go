package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/harisriyoni/restful-api-go/exception"
	"github.com/harisriyoni/restful-api-go/helper"
	"github.com/harisriyoni/restful-api-go/model/domain"
	"github.com/harisriyoni/restful-api-go/model/web"
	"github.com/harisriyoni/restful-api-go/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}
	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// aksi validasi request terlebih dahulu
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	// buka transaksi ke databaset
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// jika terjadi satu saja kesalahan maka akan di rollback Oleh helper yang telah kita buat CommitOrRollback
	// kalau semua berlajan tanpa ada error maka akan dicommit dan berhasil untuk transaksi ke databasenya dan statusnya berhasil.
	defer helper.CommitOrRollback(tx)
	// ambil id lewat category repository lalu dari category repository memanggil fungsi FindById lalu mengqueri dan mengambil Id nya
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	// jika sudah berhasil dan tidak mendapatkan error maka category.Name si Name nya itu di ganti dengan request yang baru
	category.Name = request.Name
	// lalu melakukan update lewat category repository dan memanggil fungsi update lalu mengqueri untuk mengubah data ke database
	category = service.CategoryRepository.Update(ctx, tx, category)
	// lalu tampilkan data yang telah berhasil ter ubah
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	categories := service.CategoryRepository.FindAll(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponses(categories)
}
