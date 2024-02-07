package usecases

import "belobetty-queue-manager/domain"

type UseCaseInterface interface {
	Exec(entity domain.Entity, action, functionality, user string) error
}
