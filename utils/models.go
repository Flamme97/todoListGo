package utils

import (
	"Documents/todoListgo/internal/database"
	"time"

	"github.com/flamme97/todolistgo/internal/database"
	"github.com/google/uuid"
)


type ToDoList struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string 		`json:"name"`
	APIkey 		string 		`json:"api_key"`
}


func DatabaselistToList(dbToDoList database.Todolist) ToDoList {
	return ToDoList{
		List: dbToDoList.list,
		CreatedAt: dbToDoList.CreatedAt,
		UpdatedAt: dbToDoList.UpdatedAt,
		Complete: dbToDoList.complete,
		}
}


type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string 		`json:"name"`
	Url 			string    `json:"URL"`
	UserID		uuid.UUID `json:"user_id"`
	
}


func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: 			 dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: 		 dbFeed.Name,
		Url: 			 dbFeed.Url,
		UserID: 	 dbFeed.UserID,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(dbFeed))
	}
	return feeds
}


type Post struct {
	ID       	  uuid.UUID `json:"id"`
	FeedID      uuid.UUID `json:"feed_id"`
	Description *string 	`json:"description"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
	Title      	string 		`json:"title"`
	Url 				string    `json:"URL"`
}

func DatabasePostsToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid{
		description = &dbPost.Description.String
	}
	return Post{
		ID:          dbPost.ID,
		FeedID:      dbPost.FeedID,
		Description: description,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.CreatedAt,
		PublishedAt: dbPost.PublishedAt,
		Title:       dbPost.Title,
		Url:         dbPost.Title,
	}
}

func DatabasePostsToPosts(dbposts []database.Post) []Post{
	posts := []Post{}

	for _, post := range dbposts {
		posts = append(posts, DatabasePostsToPost(post))
	}
	return posts

}