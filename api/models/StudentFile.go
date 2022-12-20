package models

import "github.com/jinzhu/gorm"

type StudentFile struct {
	gorm.Model
	SEETranscript         string  `json:"see_transcript"`
	SEECharacter          string  `json:"see_character"`
	CertifiacteTranscript string  `json:"certificate_transcript"`
	CertificateCharacter  string  `json:"certificate_character"`
	CertificateMigration  string  `json:"certificate_migration"`
	CitizenshipFront      string  `json:"citizenship_front"`
	CitizenshipBack       string  `json:"citizenship_back"`
	StudentID             uint    `json:"student_id"`
	Student               Student `json:"student"`
}
