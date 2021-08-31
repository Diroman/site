package database

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"

	"lawSite/models"
)

var ErrNoPost = errors.New("Событие с таким id отсутствует ")

func (c *Client) GetPostByID(id int64) (*models.Post, error) {
	ctx := context.Background()

	q := sq.Select("posts.post_id", "title", "text", "image", "date", "place",
		"duration", "link", "event_type", "contacts", "price", "tags", "speakers",
		"owner", "youtube_link", "start_time", "end_time", "coalesce(popular, 0)",
		"social_media", "contact_phone").
		From(postsTable).
		Where(sq.Eq{"posts.post_id": id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1).
		LeftJoin(popularQuery)

	qStr, _, _ := q.ToSql()

	var p models.PostData
	var tags, speakers, socialMedia pgtype.TextArray
	err := c.db.QueryRow(ctx, qStr, id).Scan(&p.ID, &p.Title, &p.Body, &p.Image,
		&p.Date, &p.Place, &p.Duration, &p.TranslationLink, &p.EventType, &p.Contacts,
		&p.Price, &tags, &speakers, &p.Owner, &p.YoutubeLink, &p.TimeStart, &p.TimeEnd,
		&p.MemberCount, &socialMedia, &p.ContactPhone)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNoPost
		}
		return nil, err
	}

	p.Hashtag = &models.List{Data: make([]string, len(tags.Elements))}
	p.Speakers = &models.List{Data: make([]string, len(speakers.Elements))}
	p.SocialMedia = &models.List{Data: make([]string, len(socialMedia.Elements))}
	err = tags.AssignTo(&p.Hashtag.Data)
	if err != nil {
		return nil, err
	}
	err = speakers.AssignTo(&p.Speakers.Data)
	if err != nil {
		return nil, err
	}
	err = socialMedia.AssignTo(&p.SocialMedia.Data)
	if err != nil {
		return nil, err
	}
	p.EventType = translationEvent[p.EventType]

	return &models.Post{Data: &p}, nil
}

func (c *Client) GetPostsByIDs(ids []int64, byTime string) ([]*models.Post, error) {
	if len(ids) == 0 {
		return []*models.Post{}, nil
	}
	ctx := context.Background()

	var IDs = make([]interface{}, len(ids))
	for i, id := range ids {
		IDs[i] = fmt.Sprintf("%d", id)
	}

	q := sq.Select("posts.post_id", "title", "text", "image", "date", "place",
		"duration", "link", "event_type", "contacts", "price", "tags", "speakers",
		"owner", "youtube_link", "start_time", "end_time", "coalesce(popular, 0)",
		"social_media").
		From(postsTable).
		Where(sq.Eq{"posts.post_id": IDs}).
		Where(getByTimeFilter(byTime)).
		PlaceholderFormat(sq.Dollar).
		LeftJoin(popularQuery)

	qStr, _, _ := q.ToSql()
	rows, err := c.db.Query(ctx, qStr, IDs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		p := &models.PostData{}
		var tags, speakers, socialMedia pgtype.TextArray
		err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.Image, &p.Date, &p.Place, &p.Duration,
			&p.TranslationLink, &p.EventType, &p.Contacts, &p.Price, &tags, &speakers,
			&p.Owner, &p.YoutubeLink, &p.TimeStart, &p.TimeEnd, &p.MemberCount, &socialMedia)
		if err != nil {
			return nil, err
		}

		p.Hashtag = &models.List{Data: make([]string, len(tags.Elements))}
		p.Speakers = &models.List{Data: make([]string, len(speakers.Elements))}
		p.SocialMedia = &models.List{Data: make([]string, len(socialMedia.Elements))}
		err = tags.AssignTo(&p.Hashtag.Data)
		if err != nil {
			return nil, err
		}
		err = speakers.AssignTo(&p.Speakers.Data)
		if err != nil {
			return nil, err
		}
		err = socialMedia.AssignTo(&p.SocialMedia.Data)
		if err != nil {
			return nil, err
		}
		p.EventType = translationEvent[p.EventType]

		posts = append(posts, &models.Post{Data: p})
	}

	return posts, nil
}

func (c *Client) GetPostsCreatedByUserID(userID int64, byTime string) ([]*models.Post, error) {
	ctx := context.Background()

	q := sq.Select("post_id", "title", "text", "image", "date", "place",
		"duration", "link", "event_type", "contacts", "price", "tags", "speakers",
		"owner", "youtube_link", "start_time", "end_time", "social_media").
		From(postsTable).
		Where(sq.Eq{"owner": userID}).
		Where(getByTimeFilter(byTime)).
		OrderBy("date DESC").
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()
	rows, err := c.db.Query(ctx, qStr, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		p := &models.PostData{}
		var tags, speakers, socialMedia pgtype.TextArray
		err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.Image, &p.Date, &p.Place, &p.Duration,
			&p.TranslationLink, &p.EventType, &p.Contacts, &p.Price, &tags, &speakers,
			&p.Owner, &p.YoutubeLink, &p.TimeStart, &p.TimeEnd, &socialMedia)
		if err != nil {
			return nil, err
		}

		p.Hashtag = &models.List{Data: make([]string, len(tags.Elements))}
		p.Speakers = &models.List{Data: make([]string, len(speakers.Elements))}
		p.SocialMedia = &models.List{Data: make([]string, len(socialMedia.Elements))}
		err = tags.AssignTo(&p.Hashtag.Data)
		if err != nil {
			return nil, err
		}
		err = speakers.AssignTo(&p.Speakers.Data)
		if err != nil {
			return nil, err
		}
		err = socialMedia.AssignTo(&p.SocialMedia.Data)
		if err != nil {
			return nil, err
		}
		p.EventType = translationEvent[p.EventType]

		posts = append(posts, &models.Post{Data: p})
	}

	return posts, nil
}
