package domains

func UserToUserResp(u User) UserResp {
	return UserResp{
		Id:          u.Id,
		UserName:    u.UserName,
		Gender:      u.Gender,
		BirthOfDate: u.BirthOfDate,
		Email:       u.Email,
		Phone:       u.Phone,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
