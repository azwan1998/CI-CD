package appController

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	model     appModel.ProfileModel
	jwtSecret string
}

func NewProfileController(m appModel.ProfileModel, jwtSecret string) ProfileController {
	return ProfileController{
		model:     m,
		jwtSecret: jwtSecret,
	}
}

func (pc ProfileController) GetAll(c echo.Context) error {
	allProfiles, err := pc.model.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get Profiles")
	}

	// Convert each ProfileResponse to Profile to be able to use FillFilePaths
	for i := range allProfiles {
		profile := appModel.Profile{
			Id:         allProfiles[i].Id,
			Alamat:     allProfiles[i].Alamat,
			Institusi:  allProfiles[i].Institusi,
			Foto:       allProfiles[i].Foto,
			FotoIjazah: allProfiles[i].FotoIjazah,
			FotoKTP:    allProfiles[i].FotoKTP,
			Surat:      allProfiles[i].Surat,
			IsApprove:  allProfiles[i].IsApprove,
			CreatedAt:  allProfiles[i].CreatedAt,
			UpdatedAt:  allProfiles[i].UpdatedAt,
		}

		// Generate links for the profile's files
		appModel.FillFilePaths(&profile)

		// Replace the ProfileResponse with the modified Profile
		allProfiles[i] = appModel.ProfileResponse{
			Id:         allProfiles[i].Id,
			Alamat:     profile.Alamat,
			Institusi:  profile.Institusi,
			Foto:       profile.Foto,
			FotoIjazah: profile.FotoIjazah,
			FotoKTP:    profile.FotoKTP,
			Surat:      profile.Surat,
			IsApprove:  profile.IsApprove,
			CreatedAt:  profile.CreatedAt,
			UpdatedAt:  profile.UpdatedAt,
			UserName:   allProfiles[i].UserName,
			Email:      allProfiles[i].Email,
			Role:       allProfiles[i].Role,
			IdUser:     allProfiles[i].IdUser,
			IsActive:   allProfiles[i].IsActive,
		}
	}

	return c.JSON(http.StatusOK, allProfiles)
}

func (pc ProfileController) GetById(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid user ID")
	}

	profileResponse, err := pc.model.GetById(userID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get profile by ID")
	}

	// Convert the ProfileResponse to Profile
	profile := appModel.Profile{
		Id:         profileResponse.Id,
		Alamat:     profileResponse.Alamat,
		Institusi:  profileResponse.Institusi,
		Foto:       profileResponse.Foto,
		FotoIjazah: profileResponse.FotoIjazah,
		FotoKTP:    profileResponse.FotoKTP,
		Surat:      profileResponse.Surat,
		IsApprove:  profileResponse.IsApprove,
		CreatedAt:  profileResponse.CreatedAt,
		UpdatedAt:  profileResponse.UpdatedAt,
	}

	// Populate the photo URLs using the FillFilePaths function
	appModel.FillFilePaths(&profile)

	response := appModel.ProfileResponse{
		Id:         profileResponse.Id,
		Alamat:     profile.Alamat,
		Institusi:  profile.Institusi,
		Foto:       profile.Foto,
		FotoIjazah: profile.FotoIjazah,
		FotoKTP:    profile.FotoKTP,
		Surat:      profile.Surat,
		IsApprove:  profile.IsApprove,
		CreatedAt:  profile.CreatedAt,
		UpdatedAt:  profile.UpdatedAt,
		UserName:   profileResponse.UserName,
		Email:      profileResponse.Email,
		Role:       profileResponse.Role,
		IdUser:     profileResponse.IdUser,
		IsActive:   profileResponse.IsActive,
	}

	return c.JSON(http.StatusOK, response)
}

func (pc ProfileController) Add(c echo.Context) error {
	var Profile appModel.Profile
	if err := c.Bind(&Profile); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid Profile data")
	}
	Profile, err := pc.model.Add(Profile)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot add Profile")
	}
	return c.JSON(http.StatusOK, Profile)
}

func generateUniqueFileName() string {
	uuidObj := uuid.New()
	return uuidObj.String()
}

func (pc ProfileController) uploadFile(c echo.Context, fieldname string) (string, error) {
	err := c.Request().ParseMultipartForm(10 << 20) // Max 10 MB
	if err != nil {
		return "", err
	}

	file, handler, err := c.Request().FormFile(fieldname)
	if err != nil && err != http.ErrMissingFile {
		return "", err
	}

	if handler != nil {
		// Buat nama file yang unik menggunakan UUID
		fileExt := filepath.Ext(handler.Filename)
		fileName := generateUniqueFileName() + fileExt
		filePath := "C:/laragon/www/GoApp/IslamicNews/storage/file/" + fileName

		// Simpan file di lokasi yang diinginkan
		dst, err := os.Create(filePath)
		if err != nil {
			return "", err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			return "", err
		}

		return fileName, nil
	}

	return "", nil
}

func (pc ProfileController) Edit(c echo.Context) error {
	userInfo := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("Current user id: ", userInfo.IdUser)
	profileID, err := pc.model.GetById(userInfo.IdUser)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "ID profil tidak valid")
	}

	var existingProfile appModel.Profile // Simpan data existing profile dari database
	err = pc.model.GetByID(profileID.Id, &existingProfile)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Tidak dapat mengambil data profil")
	}

	// Bind data dari request ke variabel profile
	if err := c.Bind(&existingProfile); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Data profil tidak valid")
	}

	// Upload dan simpan file foto
	fotoFileName, err := pc.uploadFile(c, "foto")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Tidak dapat mengunggah file foto")
	}
	if fotoFileName != "" {
		existingProfile.Foto = fotoFileName
	}

	// Upload dan simpan file fotoIjazah
	fotoIjazahFileName, err := pc.uploadFile(c, "fotoIjazah")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Tidak dapat mengunggah file fotoIjazah")
	}
	if fotoIjazahFileName != "" {
		existingProfile.FotoIjazah = fotoIjazahFileName
	}

	// Upload dan simpan file fotoKTP
	fotoKTPFileName, err := pc.uploadFile(c, "fotoKTP")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Tidak dapat mengunggah file fotoKTP")
	}
	if fotoKTPFileName != "" {
		existingProfile.FotoKTP = fotoKTPFileName
	}

	// Upload dan simpan file surat
	suratFileName, err := pc.uploadFile(c, "surat")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Tidak dapat mengunggah file surat")
	}
	if suratFileName != "" {
		existingProfile.Surat = suratFileName
	}

	// Panggil fungsi Edit di model untuk mengupdate data
	updatedProfile, err := pc.model.Edit(profileID.Id, existingProfile)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Tidak dapat mengedit profil")
	}

	// Isi URL file untuk bidang foto, fotoIjazah, fotoKTP, dan surat
	appModel.FillFilePaths(&updatedProfile)

	return c.JSON(http.StatusOK, updatedProfile)
}

func (nc ProfileController) ApproveProfile(c echo.Context) error {
	ProfileId, err := strconv.Atoi(c.Param("id"))
	userInfo := appMiddleware.ExtractTokenUserId(c)
	if userInfo.Role != "admin" {
		return c.String(http.StatusForbidden, "You are not allowed")
	}

	var profile appModel.Profile
	if err := c.Bind(&profile); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid Profile data")
	}

	profile.IsApprove = true

	// Ubah status berita menjadi "accept"
	profile, err = nc.model.ApproveProfile(ProfileId, profile)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot approve Profile")
	}

	return c.JSON(http.StatusOK, profile)
}
