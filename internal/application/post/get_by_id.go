package post

import "app/internal/domain/post"

func (r *postReaderService) GetByID(ID uint) (*post.Post, error) {
	p, err := r.PostRepository.FindByID(ID)
	if err != nil {
		return nil, post.ErrPostNotFound
	}
	return p, nil
}
