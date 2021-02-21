package domain

import (
	"net/http"

	"cqrses/internal/article/domain/command"
	"cqrses/internal/article/domain/entities"
	"cqrses/internal/pkg/service"
	"cqrses/pkg/cqrs"
	"cqrses/pkg/payload"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ArticleService struct {
	service.DomainService
}

func NewArticleService() *ArticleService {
	return new(ArticleService)
}

func (s ArticleService) validation(article entities.ArticleData) error {
	if err := validation.ValidateStruct(&article,
		validation.Field(&article.Title, validation.Required),
		validation.Field(&article.Content, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

func (s ArticleService) Create(data entities.ArticleData) payload.DomainPayload {
	if err := s.validation(data); err != nil {
		return s.NewPayload(http.StatusBadRequest, err)
	}

	commandMessage := command.ArticleCommandMessage{
		ArticleData: data,
	}

	cmdDescriptor := cqrs.NewCommandMessage(&commandMessage)

	res, err := CommandBus.Dispatch(cmdDescriptor)
	if err != nil {
		return s.NewPayload(http.StatusNotFound, err)
	}

	return s.NewPayload(http.StatusCreated, res)
}

func (s ArticleService) Update(data entities.ArticleData) payload.DomainPayload {
	if err := s.validation(data); err != nil {
		return s.NewPayload(http.StatusBadRequest, err)
	}

	commandMessage := command.ArticleCommandMessage{
		ArticleData: data,
	}

	cmdDescriptor := cqrs.NewCommandMessage(&commandMessage)

	res, err := CommandBus.Dispatch(cmdDescriptor)
	if err != nil {
		return s.NewPayload(http.StatusNotFound, err)
	}

	return s.NewPayload(http.StatusCreated, res)
}
