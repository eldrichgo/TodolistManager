package dal

import (
	"errors"
	"strings"
	"survey/graph/model"
)

type SurveyService struct {
	repo SurveyRepository
}

func NewSurveyService(repo SurveyRepository) *SurveyService {
	return &SurveyService{repo: repo}
}

func FilterQuestions(questions []string) []string {
	filteredQuestions := []string{}

	for _, question := range questions {
		if question != "" {
			if strings.TrimSpace(question) != "" {
				filteredQuestions = append(filteredQuestions, question)
			}
		}
	}

	return filteredQuestions
}

func (s *SurveyService) CreateSurvey(name string, description *string, questions []string) (*model.Survey, error) {
	if name == "" {
		return nil, errors.New("Survey name cannot be blank")
	}

	if len(questions) == 0 {
		return nil, errors.New("Survey must have at least one question")
	}

	questions = FilterQuestions(questions)

	return s.repo.CreateSurvey(name, description, questions)
}

func (s *SurveyService) GetSurveys() ([]*model.Survey, error) {
	return s.repo.GetSurveys()
}

func (s *SurveyService) GetSurvey(id string) (*model.Survey, error) {
	return s.repo.GetSurvey(id)
}

func (s *SurveyService) UpdateSurvey(id string, name string, description *string, questions []string) (*model.Survey, error) {
	return s.repo.UpdateSurvey(id, name, description, questions)
}
