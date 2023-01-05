package domain

import "online-learning-restful-api/core"

type MasterAccount struct {
	core.DomainModel
	Email    string
	Password string
}
