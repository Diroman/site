package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	"lawSite/models"
)

func (c *Client) UpdatePost(post *models.PostData) error {
	ctx := context.Background()

	var tags, speakers, socialMedia interface{} = "{}", "{}", "{}"
	if post.Hashtag != nil {
		tags = pq.Array(post.Hashtag.Data)
	}
	if post.Speakers != nil {
		speakers = pq.Array(post.Speakers.Data)
	}
	if post.SocialMedia != nil {
		socialMedia = pq.Array(post.SocialMedia.Data)
	}

	q := sq.Update(postsTable).
		Where(sq.Eq{"post_id": post.ID}).
		Set("title", post.Title).
		Set("text", post.Body).
		Set("image", post.Image).
		Set("date", post.Date).
		Set("place", post.Place).
		Set("duration", post.Duration).
		Set("link", post.TranslationLink).
		Set("event_type", post.EventType).
		Set("contacts", post.Contacts).
		Set("price", post.Price).
		Set("tags", tags).
		Set("speakers", speakers).
		Set("start_time", post.TimeStart).
		Set("end_time", post.TimeEnd).
		Set("social_media", socialMedia).
		PlaceholderFormat(sq.Dollar)

	qStr, args, _ := q.ToSql()
	_, err := c.db.Exec(ctx, qStr, args...)
	if err != nil {
		return err
	}

	return nil
}
