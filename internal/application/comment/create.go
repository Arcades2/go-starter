package comment

import "app/internal/domain/comment"

func (s *commentService) Create(cmd CreateCommentCommand) (*comment.Comment, error) {
	if cmd.AuthorID != nil {
		_, err := s.userReader.GetByID(*cmd.AuthorID)
		if err != nil {
			return nil, err
		}
	}

	_, err := s.postReader.GetByID(cmd.PostID)
	if err != nil {
		return nil, err
	}

	comment, err := comment.NewComment(cmd.Content, cmd.AuthorID, cmd.PostID)
	if err != nil {
		return nil, err
	}

	err = s.repository.Create(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

type CreateCommentCommand struct {
	Content  string
	AuthorID *uint
	PostID   uint
}
