package dal

import (
	"survey/graph/model"

	"gorm.io/gorm"
)

type SurveyRepository interface {
	CreateSurvey(name string, description *string, questions []string) (*model.Survey, error)
	GetSurveys() ([]*model.Survey, error)
	GetSurvey(id string) (*model.Survey, error)
	UpdateSurvey(id string, name string, description *string, questions []string) (*model.Survey, error)
}

type Survey struct {
	db *gorm.DB
}

func NewSurveyRepository(db *gorm.DB) SurveyRepository {
	return &Survey{db: db}
}

func (s *Survey) CreateSurvey(name string, description *string, questions []string) (*model.Survey, error) {
	// Create survey
	survey := &model.Survey{
		Name:        name,
		Description: description,
	}

	if err := s.db.Create(survey).Error; err != nil {
		return nil, err
	}

	// Create questions
	for _, questionText := range questions {
		var question model.Question

		// Check if question already exists
		if err := s.db.Where("question_text = ?", questionText).First(&question).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create question
				question = model.Question{
					QuestionText: questionText,
				}

				if err := s.db.Create(&question).Error; err != nil {
					return nil, err
				}

			} else {
				return nil, err
			}
		}

		// Create survey_question
		surveyQuestion := &model.SurveyQuestion{
			SurveyID:   survey.ID,
			QuestionID: question.ID,
		}

		if err := s.db.Create(surveyQuestion).Error; err != nil {
			return nil, err
		}
	}

	return survey, nil
}

func (s *Survey) GetSurveys() ([]*model.Survey, error) {
	var surveys []*model.Survey

	if err := s.db.Find(&surveys).Error; err != nil {
		return nil, err
	}

	return surveys, nil
}

func (s *Survey) GetSurvey(id string) (*model.Survey, error) {
	var survey *model.Survey

	if err := s.db.Where("id = ?", id).First(&survey).Error; err != nil {
		return nil, err
	}

	return survey, nil
}

func (s *Survey) UpdateSurvey(id string, name string, description *string, questions []string) (*model.Survey, error) {
	var survey *model.Survey

	if err := s.db.Model(&model.Survey{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        name,
		"description": description,
	}).Error; err != nil {
		return nil, err
	}

	if err := s.db.First(&survey, id).Error; err != nil {
		return nil, err
	}

	// Question handling
	for _, questionText := range questions {
		var question model.Question

		// Check if question already exists
		if err := s.db.Where("question_text = ?", questionText).First(&question).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create question
				question = model.Question{
					QuestionText: questionText,
				}

				if err := s.db.Create(&question).Error; err != nil {
					return nil, err
				}

			} else {
				return nil, err
			}
		}

		var surveyQuestion model.SurveyQuestion

		// Check if survey question already exists
		if err := s.db.Where("survey_id = ? AND question_id = ?", survey.ID, question.ID).First(&surveyQuestion).Error; err == nil {
			if err == gorm.ErrRecordNotFound {
				// Create survey question
				surveyQuestion := &model.SurveyQuestion{
					SurveyID:   survey.ID,
					QuestionID: question.ID,
				}

				if err := s.db.Create(surveyQuestion).Error; err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
	}

	return survey, nil
}
