package models

import "github.com/jinzhu/gorm"

type StudentFile struct {
	gorm.Model
	Photo                 string  `json:"photo"`
	SEETranscript         string  `json:"see_transcript"`
	SEECharacter          string  `json:"see_character"`
	CertifiacteTranscript string  `json:"certificate_transcript"`
	CertificateCharacter  string  `json:"certificate_character"`
	CertificateMigration  string  `json:"certificate_migration"`
	CitizenshipFront      string  `json:"citizenship_front"`
	CitizenshipBack       string  `json:"citizenship_back"`
	Signature             string  `json:"signature"`
	StudentID             uint    `json:"student_id"`
	Student               Student `json:"student"`
}
