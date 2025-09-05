package employeeDto

import (
	"api/app/models/model/employeeModel"
	"fmt"

	"github.com/jinzhu/copier"
)

type EmployeeDocumentResponse struct {
	ID                  int64   `json:"id"`
	EmployeeID          int64   `json:"employee_id"`
	PhotoURL            *string `json:"photo_url"`
	KtpURL              *string `json:"ktp_url"`
	KkURL               *string `json:"kk_url"`
	CertificateURL      *string `json:"certificate_url"`
	GradeTranscriptURL  *string `json:"grade_transcript_url"`
	CertificateSkillURL *string `json:"certificate_skill_url"`
	BankAccountURL      *string `json:"bank_account_url"`
	NpwpURL             *string `json:"npwp_url"`
	BpjsKtnURL          *string `json:"bpjs_ktn_url"`
	BpjsKesURL          *string `json:"bpjs_kes_url"`
}

func ToEmployeeDocumentResponse(model employeeModel.EmployeeDocumentModel) EmployeeDocumentResponse {
	var photoURL *string
	var KtpURL *string
	var kkURL *string
	var certificateURL *string
	var gradeTranscriptURL *string
	var certificateSkillURL *string
	var bankAccountURL *string
	var npwpURL *string
	var bpjsKtnURL *string
	var bpjsKesURL *string
	if model.Photo != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.Photo)
		photoURL = &url
	}

	if model.Ktp != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.Ktp)
		KtpURL = &url
	}

	if model.Kk != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.Kk)
		kkURL = &url
	}

	if model.Certificate != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.Certificate)
		certificateURL = &url
	}

	if model.GradeTranscript != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.GradeTranscript)
		gradeTranscriptURL = &url
	}

	if model.CertificateSkillURL != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.CertificateSkillURL)
		certificateSkillURL = &url
	}

	if model.BankAccount != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.BankAccount)
		bankAccountURL = &url
	}

	if model.Npwp != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.Npwp)
		npwpURL = &url
	}

	if model.BpjsKtn != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.BpjsKtn)
		bpjsKtnURL = &url
	}

	if model.BpjsKes != nil { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *model.BpjsKes)
		bpjsKesURL = &url
	}

	var data EmployeeDocumentResponse

	copier.Copy(&data, &model)
	data.PhotoURL = photoURL
	data.KtpURL = KtpURL
	data.KkURL = kkURL
	data.CertificateURL = certificateURL
	data.GradeTranscriptURL = gradeTranscriptURL
	data.CertificateSkillURL = certificateSkillURL
	data.BankAccountURL = bankAccountURL
	data.NpwpURL = npwpURL
	data.BpjsKtnURL = bpjsKtnURL
	data.BpjsKesURL = bpjsKesURL

	return data

}
