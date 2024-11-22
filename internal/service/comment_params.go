package service

type CommentCreateParams struct {
	TenderID    int
	Content     string
	Attachments []string
}

type GetCommentParams struct {
	TenderID int
}
