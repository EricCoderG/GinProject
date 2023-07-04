package dto

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(name string, telephone string) UserDto {
	return UserDto{
		Name:      name,
		Telephone: telephone,
	}
}
