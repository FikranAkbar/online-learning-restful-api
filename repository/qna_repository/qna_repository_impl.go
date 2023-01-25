package qna_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/model/domain"
)

type QnaRepositoryImpl struct {
}

func NewQnaRepositoryImpl() *QnaRepositoryImpl {
	return &QnaRepositoryImpl{}
}

func (repository *QnaRepositoryImpl) GetQnaQuestionsByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.QnaQuestion, []domain.User, error) {
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		Preload("QnaQuestions").
		Preload("QnaQuestions.User").
		Preload("QnaQuestions.Answers").
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) || !courseEntity.IsPublished {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return nil, nil, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return nil, nil, err
	}

	qnaQuestionEntities := courseEntity.QnaQuestions
	var qnaQuestions []domain.QnaQuestion
	var users []domain.User
	for _, qnaQuestionEntity := range qnaQuestionEntities {
		qnaQuestions = append(qnaQuestions, domain.QnaQuestion{
			Id:        qnaQuestionEntity.ID,
			CreatedAt: qnaQuestionEntity.CreatedAt,
			CourseId:  qnaQuestionEntity.CourseId,
			UserId:    qnaQuestionEntity.UserId,
			Question:  qnaQuestionEntity.Question,
			Responses: len(qnaQuestionEntity.Answers),
		})

		users = append(users, domain.User{
			Name:     qnaQuestionEntity.User.Name,
			PhotoURL: qnaQuestionEntity.User.PhotoURL.String,
		})
	}

	return qnaQuestions, users, nil
}

func (repository *QnaRepositoryImpl) CreateNewQnaQuestion(ctx context.Context, db *gorm.DB, courseId uint, qnaQuestion domain.QnaQuestion) (domain.QnaQuestion, error) {
	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) || !courseEntity.IsPublished {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return qnaQuestion, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return qnaQuestion, err
	}

	qnaQuestionEntity := entity.TrxCourseQnaQuestion{
		CourseId: courseId,
		Question: qnaQuestion.Question,
		UserId:   userTokenInfo.UserId,
	}
	err = db.WithContext(ctx).
		Create(&qnaQuestionEntity).
		Error
	if err != nil {
		return qnaQuestion, err
	}

	return domain.QnaQuestion{
		Id:        qnaQuestionEntity.ID,
		CreatedAt: qnaQuestionEntity.CreatedAt,
		CourseId:  qnaQuestionEntity.CourseId,
		UserId:    qnaQuestionEntity.UserId,
		Question:  qnaQuestionEntity.Question,
		Responses: 0,
	}, nil
}

func (repository *QnaRepositoryImpl) GetDetailQnaQuestionByQnaQuestionId(ctx context.Context, db *gorm.DB, courseId uint, qnaQuestionId uint) (domain.QnaQuestion, []domain.QnaAnswer, error) {
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		Preload("QnaQuestions", "id = ?", qnaQuestionId).
		Preload("QnaQuestions.User").
		Preload("QnaQuestions.Answers", "qna_question_id = ?", qnaQuestionId).
		Preload("QnaQuestions.Answers.User").
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) || !courseEntity.IsPublished {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return domain.QnaQuestion{}, nil, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.QnaQuestion{}, nil, err
	}

	if len(courseEntity.QnaQuestions) <= 0 {
		logError := fmt.Sprintf("Qna question with id %v not found", qnaQuestionId)
		return domain.QnaQuestion{}, nil, exception.GenerateHTTPError(exception.NotFound, logError)
	}

	qnaQuestionEntity := courseEntity.QnaQuestions[0]
	qnaQuestion := domain.QnaQuestion{
		Id:        qnaQuestionEntity.ID,
		CreatedAt: qnaQuestionEntity.CreatedAt,
		CourseId:  qnaQuestionEntity.CourseId,
		UserId:    qnaQuestionEntity.UserId,
		UserName:  qnaQuestionEntity.User.Name,
		UserPhoto: qnaQuestionEntity.User.PhotoURL.String,
		Question:  qnaQuestionEntity.Question,
		Responses: len(qnaQuestionEntity.Answers),
	}

	qnaAnswerEntities := qnaQuestionEntity.Answers
	var qnaAnswers []domain.QnaAnswer
	for _, qnaAnswerEntity := range qnaAnswerEntities {
		qnaAnswers = append(qnaAnswers, domain.QnaAnswer{
			Id:            qnaAnswerEntity.ID,
			CreatedAt:     qnaAnswerEntity.CreatedAt,
			QnaQuestionId: qnaAnswerEntity.QnaQuestionId,
			UserId:        qnaAnswerEntity.UserId,
			UserName:      qnaAnswerEntity.User.Name,
			UserPhoto:     qnaAnswerEntity.User.PhotoURL.String,
			Answer:        qnaAnswerEntity.Answer,
		})
	}

	return qnaQuestion, qnaAnswers, nil
}
