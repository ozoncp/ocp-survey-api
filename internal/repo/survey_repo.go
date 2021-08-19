package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-survey-api/internal/models"
)

type surveyRepo struct {
	db *sqlx.DB
}

func NewSurveyRepo(db *sqlx.DB) Repo {
	return &surveyRepo{
		db: db,
	}
}

func (r *surveyRepo) AddSurvey(ctx context.Context, surveys []models.Survey) error {
	query := `INSERT INTO surveys (user_id, link) 
			VALUES ($1, $2);`

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prep.Close()

	for _, survey := range surveys {
		_, err := prep.ExecContext(ctx, survey.UserId, survey.Link)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *surveyRepo) ListSurveys(ctx context.Context, limit, offset uint64) ([]models.Survey, error) {
	query := `SELECT id, user_id, link 
			FROM surveys LIMIT $1 OFFSET $2;`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []models.Survey{}
	for rows.Next() {
		survey := models.Survey{}
		err := rows.Scan(&survey.Id, &survey.UserId, &survey.Link)
		if err != nil {
			return nil, err
		}
		res = append(res, survey)
	}

	return res, nil
}

func (r *surveyRepo) DescribeSurvey(ctx context.Context, surveyId uint64) (*models.Survey, error) {
	query := `SELECT id, user_id, link
			FROM surveys WHERE id=$1;`

	rows, err := r.db.QueryContext(ctx, query, surveyId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		survey := &models.Survey{}
		err = rows.Scan(&survey.Id, &survey.UserId, &survey.Link)
		if err != nil {
			return nil, err
		}
		return survey, nil
	}

	return nil, ErrNotFound
}

func (r *surveyRepo) RemoveSurvey(ctx context.Context, surveyId uint64) error {
	query := `DELETE FROM surveys WHERE id=$1;`

	res, err := r.db.ExecContext(ctx, query, surveyId)
	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num == 0 {
		return ErrNotFound
	}

	return nil
}
