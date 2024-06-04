package release_note

import (
	"pension-reservation-api/mod/release_note"
	"pension-reservation-api/model"

	"github.com/pkg/errors"
)

func (q *queryImpl) findLatestReleaseNotes() (release_note.GetLatestReleaseNotesQueryResult, error) {
	rs := []*model.ReleaseNote{}
	limit := 3
	queryResult := release_note.GetLatestReleaseNotesQueryResult{}

	err := q.db.Model(&model.ReleaseNote{}).Limit(limit).Find(&rs).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, r := range rs {
		queryResult = append(queryResult, release_note.GetLatestReleaseNotesQueryResultInner{
			ID:        int(r.ID),
			Text:      r.Text,
			CreatedAt: r.CreatedAt,
		})
	}

	return queryResult, nil
}
