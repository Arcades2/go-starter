package post

import "app/internal/domain/post"

func (r *postReaderService) GetByID(ID uint) (*post.Post, error) {
	post, err := r.PostRepository.FindByID(ID)
	if err != nil {
		return nil, ErrPostNotFound
	}
	return post, nil
}
