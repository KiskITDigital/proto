package questionanswer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *QuestionAnswerStore) Get(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.QuestionWithAnswer, error) {
	builder := squirrel.
		Select(
			"q.id",
			"q.tender_id",
			"q.type",
			"q.content",
			"q.author_organization_id",

			"a.id ",
			"a.tender_id",
			"a.type",
			"a.content",
			"a.parent_id",
			"a.author_organization_id",
		).
		From("question_answer q").
		LeftJoin("question_answer a ON a.parent_id = q.id").
		Where(squirrel.Eq{"q.tender_id": tenderID}).
		Where(squirrel.Eq{"q.type": models.QuestionAnswerTypeQuestion}).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var questionWithAnswers []models.QuestionWithAnswer

	for rows.Next() {
		var (
			qa                         models.QuestionWithAnswer
			answerID                   sql.NullInt64
			answerTenderID             sql.NullInt64
			answerType                 sql.NullInt16
			answerContent              sql.NullString
			answerParentID             sql.NullInt64
			answerAuthorOrganizationID sql.NullInt64
		)

		err = rows.Scan(
			&qa.Question.ID,
			&qa.Question.TenderID,
			&qa.Question.Type,
			&qa.Question.Content,
			&qa.Question.AuthorOrganizationID,
			&answerID,
			&answerTenderID,
			&answerType,
			&answerContent,
			&answerParentID,
			&answerAuthorOrganizationID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		if answerParentID.Valid {
			qa.Answer = models.NewOptional(models.QuestionAnswer{
				ID:                   int(answerID.Int64),
				TenderID:             int(answerTenderID.Int64),
				AuthorOrganizationID: int(answerAuthorOrganizationID.Int64),
				ParentID:             models.NewOptional(int(answerParentID.Int64)),
				Type:                 models.QuestionAnswerType(answerType.Int16),
				Content:              answerContent.String,
			})
		}

		questionWithAnswers = append(questionWithAnswers, qa)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration: %w", rows.Err())
	}

	return questionWithAnswers, nil
}
