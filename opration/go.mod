module movie_opration

go 1.12

replace git.woa.com/crotaliu/pb-hub => ./stub/git.woa.com/crotaliu/pb-hub

require (
	git.code.oa.com/trpc-go/trpc-config-rainbow v0.1.24
	git.code.oa.com/trpc-go/trpc-database/gorm v0.2.1
	git.code.oa.com/trpc-go/trpc-filter/debuglog v0.1.5
	git.code.oa.com/trpc-go/trpc-filter/recovery v0.1.3
	git.code.oa.com/trpc-go/trpc-go v0.9.4
	git.code.oa.com/trpc-go/trpc-log-atta v0.1.14
	git.code.oa.com/trpc-go/trpc-metrics-m007 v0.5.1
	git.code.oa.com/trpc-go/trpc-metrics-runtime v0.3.3
	git.code.oa.com/trpc-go/trpc-naming-polaris v0.3.3
	git.woa.com/crotaliu/pb-hub v0.0.0-00010101000000-000000000000
	github.com/frankban/quicktest v1.11.3 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.6.0
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)
