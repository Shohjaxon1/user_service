package repo

import (
	pb "user_service/genproto/user_service"
)

// UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	GetById(*pb.GetByIdRequest) (*pb.User, error)
	GetByFirstName(*pb.GetByFirstNameRequest) (*pb.User, error)
	GetByLastName(*pb.GetByLastNameRequest) (*pb.User, error)
	GetByPhoneNumber(*pb.GetByPhoneNumberRequest) (*pb.User, error)
	GetAll(*pb.GetAllRequest) (*pb.AllUsers, error)
	Update(*pb.User) (*pb.User, error)
	Delete(*pb.GetByIdRequest) (*pb.User, error)
}