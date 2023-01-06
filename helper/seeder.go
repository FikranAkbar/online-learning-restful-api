package helper

import gormseeder "github.com/kachit/gorm-seeder"

func ReverseSeeders(s []gormseeder.SeederInterface) []gormseeder.SeederInterface {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
