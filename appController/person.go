package appController

import (
	"fmt"
	"net/http"
	"strconv"

	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
)

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PersonController struct {
	model        appModel.PersonModel
	profileModel appModel.ProfileModel
	jwtSecret    string
}

func NewPersonController(jwtSecret string, m appModel.PersonModel, profileModel appModel.ProfileModel) PersonController {
	return PersonController{
		jwtSecret:    jwtSecret,
		model:        m,
		profileModel: profileModel,
	}
}

func (pc PersonController) Login(c echo.Context) error {
	loginInfo := LoginInfo{}
	c.Bind(&loginInfo)
	person, err := pc.model.GetByEmail(loginInfo.Email)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot login")
	}

	if !appMiddleware.VerifyPassword(loginInfo.Password, person.Password) {
		return c.String(http.StatusUnauthorized, "invalid credentials")
	}

	token, err := appMiddleware.CreateToken(int(person.ID), person.Role, pc.jwtSecret)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot create token")
	}
	person.Token = token

	person, err = pc.model.Edit(int(person.ID), person)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot add token")
	}
	return c.JSON(http.StatusOK, person)
}

func (pc PersonController) Register(c echo.Context) error {
	person := appModel.Person{}
	c.Bind(&person)

	// Enkripsi kata sandi sebelum menyimpan ke database
	encryptedPassword, err := appMiddleware.EncryptPassword(person.Password)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot register")
	}

	person.Password = encryptedPassword

	newPerson, err := pc.model.Add(person)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot register")
	}

	userID := int(newPerson.ID)

	profile := appModel.Profile{
		IdUser: userID,
		// Set other profile data here if needed
	}

	_, err = pc.profileModel.Add(profile)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot create profile")
	}

	// Generate token
	token, err := appMiddleware.CreateToken(int(newPerson.ID), person.Role, pc.jwtSecret)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot generate token")
	}
	newPerson.Token = token

	return c.JSON(http.StatusOK, newPerson)
}

func (pc PersonController) GetAll(c echo.Context) error {
	currentLoginPersonId := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("😸 Current user id: ", currentLoginPersonId)
	allPersons, err := pc.model.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get persons")
	}
	return c.JSON(http.StatusOK, allPersons)
}

func (pc PersonController) Add(c echo.Context) error {
	var person appModel.Person
	if err := c.Bind(&person); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person data")
	}
	person, err := pc.model.Add(person)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add person")
	}
	return c.JSON(http.StatusOK, person)
}

func (pc PersonController) Edit(c echo.Context) error {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person id")
	}
	var person appModel.Person
	if err := c.Bind(&person); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person data")
	}
	person, err = pc.model.Edit(personId, person)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot edit person")
	}
	return c.JSON(http.StatusOK, person)
}
