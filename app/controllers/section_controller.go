package controllers

import (
	"errors"
	"fmt"
	"harmonee/app/collections"
	"harmonee/app/models"
	"harmonee/app/services"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type SectionController struct {
	sectionService services.SectionService
}

func NewSectionController(sectionService *services.SectionService) SectionController {
	return SectionController{
		sectionService: *sectionService,
	}
}

// Get all data
func (sc SectionController) FindAll(w http.ResponseWriter, r *http.Request) {
	timeFormated := time.Now().Format("02 February, 2006")

	// Params Query
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}
	perPage := 10

	// Service
	sections, totalRow, err := sc.sectionService.FindAll(perPage, page)
	if err != nil {
		log.Fatal(err)
	}

	// Get Pagination Link
	pagination, _ := GetPaginationLink(&collections.AppConfig{}, collections.PaginationParams{
		Path:        "sections",
		TotalRow:    totalRow,
		PerPage:     perPage,
		CurrentPage: page,
	})

	// Render data to template
	render := render.New(render.Options{
		Layout: "layout",
	})
	_ = render.HTML(w, http.StatusOK, "sections", map[string]interface{}{
		"title":      "Section",
		"username":   "Asep",
		"date":       timeFormated,
		"sections":   sections,
		"pagination": pagination,
		"success":    GetFlashMessage(w, r, "success"),
		"error":      GetFlashMessage(w, r, "error"),
	})
}

// Create
func (sc SectionController) Create(w http.ResponseWriter, r *http.Request) {
	timeFormated := time.Now().Format("02 February, 2006")

	// Render data to template
	render := render.New(render.Options{
		Layout: "layout",
	})
	_ = render.HTML(w, http.StatusOK, "section-create", map[string]interface{}{
		"title":    "Create Section",
		"username": "Asep",
		"date":     timeFormated,
		"success":  GetFlashMessage(w, r, "success"),
		"error":    GetFlashMessage(w, r, "error"),
	})
}

// Get detail by ID
func (sc SectionController) Detail(w http.ResponseWriter, r *http.Request) {
	timeFormated := time.Now().Format("02 February, 2006")

	// Params ID
	vars := mux.Vars(r)
	if vars["id"] == "" {
		return
	}
	id, _ := strconv.Atoi(vars["id"])

	// Service
	section, err := sc.sectionService.FindById(id)
	if err != nil {
		log.Fatal(err)
	}

	// Render data to template
	render := render.New(render.Options{
		Layout: "layout",
	})
	_ = render.HTML(w, http.StatusOK, "section-detail", map[string]interface{}{
		"title":    "Section Detail",
		"username": "Asep",
		"date":     timeFormated,
		"section":  section,
		"success":  GetFlashMessage(w, r, "success"),
		"error":    GetFlashMessage(w, r, "error"),
	})
}

// Post add data
func (sc SectionController) Store(w http.ResponseWriter, r *http.Request) {
	// Body request
	var attrSectionPost models.Section
	attrSectionPost.Name = r.FormValue("name")
	attrSectionPost.Description = r.FormValue("description")

	validateError := Validate.Struct(attrSectionPost)
	if validateError != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(validateError, &invalidValidationError) {
			fmt.Println(validateError)
			return
		}

		var validateErrs validator.ValidationErrors
		if errors.As(validateError, &validateErrs) {
			for _, e := range validateErrs {
				SetFlashMessage(w, r, "error", e.Field()+" is "+e.ActualTag())
			}
		}

		http.Redirect(w, r, "/create-section", http.StatusSeeOther)
		return
	}

	// Service
	err := sc.sectionService.Store(attrSectionPost)
	if err != nil {
		SetFlashMessage(w, r, "error", "Error")
		http.Redirect(w, r, "/sections", http.StatusSeeOther)
	}

	SetFlashMessage(w, r, "success", "Data has been saved successfully")
	http.Redirect(w, r, "/sections", http.StatusSeeOther)
}

// Post update data
func (sc SectionController) Update(w http.ResponseWriter, r *http.Request) {
	// Params ID
	vars := mux.Vars(r)
	if vars["id"] == "" {
		return
	}
	id, _ := strconv.Atoi(vars["id"])

	// Body request
	var attrSectionPost models.AttrSectionPost
	attrSectionPost.ID = uint(id)
	attrSectionPost.Name = r.FormValue("name")
	attrSectionPost.Description = r.FormValue("description")

	// Service
	err := sc.sectionService.Update(attrSectionPost)
	if err != nil {
		SetFlashMessage(w, r, "error", "Error")
		http.Redirect(w, r, "/section/"+strconv.Itoa(id), http.StatusSeeOther)
	}

	SetFlashMessage(w, r, "success", "Data has been saved successfully")
	http.Redirect(w, r, "/section/"+strconv.Itoa(id), http.StatusSeeOther)
}

// Delete data
func (sc SectionController) Delete(w http.ResponseWriter, r *http.Request) {
	// Params ID
	vars := mux.Vars(r)
	if vars["id"] == "" {
		return
	}
	id, _ := strconv.Atoi(vars["id"])

	// // Body request
	var attrSectionPost models.AttrSectionPost
	attrSectionPost.ID = uint(id)

	// Service
	err := sc.sectionService.Delete(attrSectionPost)
	if err != nil {
		SetFlashMessage(w, r, "error", "Error")
		http.Redirect(w, r, "/sections", http.StatusSeeOther)
	}
	SetFlashMessage(w, r, "success", "Data has been deleted successfully")
	http.Redirect(w, r, "/sections", http.StatusSeeOther)
}
