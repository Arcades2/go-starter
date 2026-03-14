package post

import "app/internal/domain/post"

func (s *postService) UpdateTitle(cmd UpdateTitleCommand) (*post.Post, error) {
	post, err := s.reader.GetByID(cmd.ID)
	if err != nil {
		return nil, ErrPostNotFound
	}

	post.Title = cmd.Title

	err = s.repository.Update(post)

	return post, err
}

type UpdateTitleCommand struct {
	ID    uint
	Title string
}
