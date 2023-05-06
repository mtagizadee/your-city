package dto

type CreateUserDto struct {
  Name string
  Surname string
  Email string
  Password string
}

func (d *CreateUserDto) Validate() {
  // to be implemented
}
