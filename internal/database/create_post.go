package database

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	"lawSite/models"
)

const defaultImage = "https://lawbox-portal.ru/api/storage/download/95af5a25367951baa2ff6cd471c483f1lawyer-with-weighing-scales-2.jpg"

func (c *Client) CreatePost(userID int64, req *models.PostData) (int64, error) {
	ctx := context.Background()

	var tags, speakers, socialMedia interface{} = "{}", "{}", "{}"
	if req.Hashtag != nil {
		tags = pq.Array(req.Hashtag.Data)
	}
	if req.Speakers != nil {
		speakers = pq.Array(req.Speakers.Data)
	}
	if req.SocialMedia != nil {
		socialMedia = pq.Array(req.SocialMedia)
	}

	q := sq.Insert(postsTable).
		Columns("title", "text", "image", "date", "place", "duration",
			"link", "event_type", "contacts", "price", "timestamp", "tags",
			"speakers", "owner", "start_time", "end_time", "social_media", "contact_phone").
		Values("title", "text", "image", "date", "place", "duration",
			"link", "event_type", "contacts", "price", "timestamp", "tags",
			"speakers", "owner", "start_time", "end_time", "social_media", "contact_phone").
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING post_id")

	qStr, _, _ := q.ToSql()

	var id int64
	if req.Image == "" {
		req.Image = defaultImage
	}

	err := c.db.QueryRow(ctx, qStr,
		req.Title, req.Body, req.Image, req.Date, req.Place, countDuration(req.TimeStart, req.TimeEnd), req.TranslationLink,
		req.EventType, req.Contacts, req.Price, "NOW()", tags, speakers, userID, req.TimeStart, req.TimeEnd, socialMedia,
		req.ContactPhone).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func countDuration(start, end string) string {
	if len(start) < 19 || len(end) < 19 {
		return ""
	}

	startTime, err := time.Parse("2006-01-02T15:04:05", start[:19])
	if err != nil {
		return ""
	}
	endTime, err := time.Parse("2006-01-02T15:04:05", end[:19])
	if err != nil {
		return ""
	}

	hoursDiff := int64(endTime.Sub(startTime).Hours())
	minutesDiff := int64(endTime.Sub(startTime).Minutes()) % 60

	var hoursText, minuteText string
	switch hoursDiff {
	case 1:
		hoursText = "час"
	case 2, 3, 4:
		hoursText = "часа"
	case 5, 6, 7, 8, 9, 10:
		hoursText = "часов"
	default:
		hoursText = "часов"
	}

	if minutesDiff > 0 {
		minuteText = fmt.Sprintf(" %d минут", minutesDiff)
	}

	return fmt.Sprintf("%d %s%s", hoursDiff, hoursText, minuteText)
}
