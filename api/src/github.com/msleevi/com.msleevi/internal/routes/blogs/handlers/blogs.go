package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/msleevi/com.msleevi/internal/routes/blogs/formats"

	"github.com/gorilla/mux"
)

// BlogPostsHandler is reponsible for collecting the blog posts based on a series of parameters
// Parameters taken into account include:
// * id: ID of the blog
// * sort: Field to sort on, includes the following: views, date, stats (sum of stats.votes)
// * search: Field to search on, does a fuzzy search on all fields
func BlogPostsHandler(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	params := mux.Vars(r)
	results, err := handleBlogs(params)

	if err != nil {
		return nil, 500, errors.New("Unable to find blog posts")
	}

	ret, err = json.Marshal(results)
	if err != nil {
		return nil, 504, errors.New("Unable to parse data into return body")
	}

	return ret, 200, nil
}

func handleBlogs(params map[string]string) (results []formats.Blog, err error) {
	return []formats.Blog{}, nil
}
