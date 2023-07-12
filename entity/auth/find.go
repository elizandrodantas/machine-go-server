package auth

func (y *Auth) FindAll(page, limit int) (res []AuthResponse, err error) {
	err = y.Client.Select(&res, `SELECT * FROM oauth ORDER BY id LIMIT $1 OFFSET $2`, limit, page)

	return
}

func (y *Auth) FindById(id int) (res AuthResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM oauth WHERE id=$1`, id)

	return
}

func (y *Auth) FindByToken(token string) (res AuthAndUserResponse, err error) {
	err = y.Client.Get(&res, `SELECT * FROM oauth INNER JOIN users ON oauth.user_id = users.id WHERE token=$1`, token)

	return
}

type AuthFindStruct struct {
	data []AuthAndUserResponse
}

func (y *Auth) FindByUserId(id int) (output *AuthFindStruct, err error) {
	var res []AuthAndUserResponse

	err = y.Client.Select(&res, `SELECT * FROM oauth INNER JOIN users ON oauth.user_id = users.id WHERE user_id=$1`, id)

	if err != nil {
		return &AuthFindStruct{}, err
	}

	return &AuthFindStruct{
		data: res,
	}, nil
}

func (y *Auth) FindByUserIdAndActive(id int) (output []AuthResponse, err error) {
	var res []AuthResponse

	err = y.Client.Select(&res, `SELECT * FROM oauth WHERE user_id=$1 AND status=true`, id)

	if err != nil {
		return []AuthResponse{}, err
	}

	return res, nil
}

func (y *Auth) FindByUserIdAndActiveWithUserData(id int) (output *AuthFindStruct, err error) {
	var res []AuthAndUserResponse

	err = y.Client.Select(&res, `SELECT * FROM oauth INNER JOIN users ON oauth.user_id = users.id WHERE user_id=$1 AND status=true`, id)

	if err != nil {
		return &AuthFindStruct{}, err
	}

	return &AuthFindStruct{
		data: res,
	}, nil
}

func (y *AuthFindStruct) First() AuthAndUserResponse {
	return y.data[0]
}

func (y *AuthFindStruct) Last() AuthAndUserResponse {
	return y.data[len(y.data)-1]
}

func (y *AuthFindStruct) Get() []AuthAndUserResponse {
	return y.data
}
