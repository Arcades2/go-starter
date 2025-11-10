package postservice

import "app/internal/domain/model"

func (s *postService) GetByID(ID uint) (*model.Post, error) {
	post, err := s.PostRepository.FindByID(ID)
	if err != nil {
		return nil, s.HandleError(NewPostError(PostErrors.ErrPostNotFound))
	}
	return post, nil
}
