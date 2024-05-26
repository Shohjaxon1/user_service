package postgres

import (
	"context"
	pb "user_service/genproto/user_service"
	l "user_service/pkg/logger"
	"user_service/storage"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewUserService(db *sqlx.DB, log l.Logger) *UserRepo {
	return &UserRepo{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}
func (s *UserRepo) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().Create(req)
}

func (s *UserRepo) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	return s.storage.User().GetById(req)
}

func (s *UserRepo) GetByFirstName(ctx context.Context, req *pb.GetByFirstNameRequest) (*pb.User, error) {
	return s.storage.User().GetByFirstName(req)
}

func (s *UserRepo) GetByLastName(ctx context.Context, req *pb.GetByLastNameRequest) (*pb.User, error) {
	return s.storage.User().GetByLastName(req)
}

func (s *UserRepo) GetByPhoneNumber(ctx context.Context, req *pb.GetByPhoneNumberRequest) (*pb.User, error) {
	return s.storage.User().GetByPhoneNumber(req)
}

func (s *UserRepo) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.AllUsers, error) {
	return s.storage.User().GetAll(req)
}

func (s *UserRepo) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().Update(req)
}
func (s *UserRepo) DeleteUser(ctx context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	return s.storage.User().Delete(req)
}
