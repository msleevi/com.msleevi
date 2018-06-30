package blog

/*
Blog structure
Blog test
*/
type Blog struct {
	Date     uint64
	Post     BlogPost
	Stats    []BlogStat
	Author   uuid
	Comments []Comment
	Views    int
}

/*
Stat structure
*/
type Stat struct {
	Vote []Vote
}

/*
Comment structure
*/
type Comment struct {
	Author  uuid
	Date    uint64
	Comment string
	Replies []Comment
	Stats   CommentStats
}

/*
Post structure
*/
type Post struct {
	Content string
}

/*
CommentStats structure
*/
type CommentStats struct {
	Vote []Vote
}

/*
Vote structure
*/
type Vote struct {
	Vote  int
	Voter uuid
	Date  uint64
}
