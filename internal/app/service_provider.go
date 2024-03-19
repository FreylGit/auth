package app

import (
	"context"
	"github.com/FreylGit/auth/internal/api/auth"
	"github.com/FreylGit/auth/internal/api/user"
	"github.com/FreylGit/auth/internal/config"
	"github.com/FreylGit/auth/internal/repository"
	"github.com/FreylGit/auth/internal/repository/refreshToken"
	"github.com/FreylGit/auth/internal/repository/role"
	userRepository "github.com/FreylGit/auth/internal/repository/user"
	"github.com/FreylGit/auth/internal/service"
	auth2 "github.com/FreylGit/auth/internal/service/auth"
	userService "github.com/FreylGit/auth/internal/service/user"
	"github.com/FreylGit/platform_common/pkg/closer"
	"github.com/FreylGit/platform_common/pkg/db"
	"github.com/FreylGit/platform_common/pkg/db/pg"
	"github.com/FreylGit/platform_common/pkg/db/transaction"
	"log"
)

type serviceProvider struct {
	grpcConfig             config.GRPCConfig
	pgConfig               config.PGConfig
	userImpl               *user.Implementation
	authImpl               *auth.Implementation
	dbClient               db.Client
	userService            service.UserService
	authService            service.AuthService
	userRepository         repository.UserRepository
	roleRepository         repository.RoleRepository
	refreshTokenRepository repository.RefreshTokenRepository

	txManager db.TxManager
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("Failed to get grpc config: %w", err.Error())
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		pgConfig, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("Failed to get pg config")
		}
		s.pgConfig = pgConfig
	}

	return s.pgConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = auth2.NewService(s.TxManager(ctx), s.UserRepository(ctx), s.RefreshTokenRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx),
			s.RoleRepository(ctx),
			s.TxManager(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) RefreshTokenRepository(ctx context.Context) repository.RefreshTokenRepository {
	if s.refreshTokenRepository == nil {
		s.refreshTokenRepository = refreshToken.NewRepository(s.DBClient(ctx))
	}

	return s.refreshTokenRepository
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) RoleRepository(ctx context.Context) repository.RoleRepository {
	if s.roleRepository == nil {
		s.roleRepository = role.NewRepository(s.DBClient(ctx))
	}

	return s.roleRepository
}
