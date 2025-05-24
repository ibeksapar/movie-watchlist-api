module movie-service

go 1.23.5

replace movie-watchlist-api => ../

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/go-resty/resty/v2 v2.16.5
	github.com/golang-migrate/migrate/v4 v4.18.3
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/stretchr/testify v1.10.0
	movie-watchlist-api v0.0.0-00010101000000-000000000000
)

require (
	github.com/bytedance/sonic v1.13.2
	github.com/bytedance/sonic/loader v0.2.4
	github.com/cloudwego/base64x v0.1.5
	github.com/cloudwego/iasm v0.2.0
	github.com/davecgh/go-spew v1.1.1
	github.com/gabriel-vasile/mimetype v1.4.8
	github.com/gin-contrib/cors v1.7.5
	github.com/gin-contrib/sse v1.0.0
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.26.0
	github.com/goccy/go-json v0.10.5
	github.com/hashicorp/errwrap v1.1.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/jackc/pgpassfile v1.0.0
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a
	github.com/jackc/pgx/v5 v5.5.5
	github.com/jackc/puddle/v2 v2.2.1
	github.com/jinzhu/inflection v1.0.0
	github.com/jinzhu/now v1.1.5
	github.com/json-iterator/go v1.1.12
	github.com/klauspost/cpuid/v2 v2.2.10
	github.com/leodido/go-urn v1.4.0
	github.com/lib/pq v1.10.9
	github.com/mattn/go-isatty v0.0.20
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.2
	github.com/pelletier/go-toml/v2 v2.2.3
	github.com/pmezard/go-difflib v1.0.0
	github.com/twitchyliquid64/golang-asm v0.15.1
	github.com/ugorji/go/codec v1.2.12
	go.uber.org/atomic v1.11.0
	golang.org/x/arch v0.15.0
	golang.org/x/crypto v0.36.0
	golang.org/x/net v0.38.0
	golang.org/x/sync v0.12.0
	golang.org/x/sys v0.31.0
	golang.org/x/text v0.23.0
	google.golang.org/protobuf v1.36.6
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/postgres v1.5.11
	gorm.io/gorm v1.25.12
)
