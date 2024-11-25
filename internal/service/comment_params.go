package service

type CommentCreateParams struct {
	TenderID    int
	Title       string
	Content     string
	Attachments []string
}

type GetCommentParams struct {
	TenderID int
}
