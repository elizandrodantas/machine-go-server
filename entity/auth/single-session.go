package auth

func (y *Auth) EnsureSingleSession(id int) {
	res, err := y.FindByUserIdAndActive(id)

	if err == nil {
		for _, value := range res {
			if value.Status {
				y.UpdateStatus(value.ID, false)
			}
		}
	}
}
