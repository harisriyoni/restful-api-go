package controller

import (
	"net/http"
	"strconv"

	"github.com/harisriyoni/restful-api-go/helper"
	"github.com/harisriyoni/restful-api-go/model/web"
	"github.com/harisriyoni/restful-api-go/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	CategoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	CategoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &CategoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	Id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	CategoryUpdateRequest.Id = Id

	CategoryResponse := controller.CategoryService.Update(request.Context(), CategoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	Id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	Id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	CategoryResponse := controller.CategoryService.FindById(request.Context(), Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	CategoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
	// settup untuk informsai lebih jelas tentang apa informasi data yang akan kita berikan  ke client atau frontend
	// writer.Header().Add("Content-Type", "application/json")
	// kirim ke client dengan writer
	// encoder := json.NewEncoder(writer)
	// encoder itu alat buat ngirim ke  client atau frontend dan apa yang dikirim? ya, webresponse/atau datanya
	// err := encoder.Encode(webResponse)
	// helper.PanicIfError(err)
}
