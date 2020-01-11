module gormat

go 1.13

require (
	fyne.io/fyne v1.2.1
	github.com/buger/jsonparser v0.0.0-20191204142016-1a29609e0929
	github.com/chenhg5/collection v0.0.0-20191118032303-cb21bccce4c3
	github.com/denisenkom/go-mssqldb v0.0.0-20191128021309-1d7a30a10f73
	github.com/fatih/astrewrite v0.0.0-20191207154002-9094e544fcef // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/structtag v1.2.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/lib/pq v1.3.0
	github.com/liudanking/gorm2sql v0.0.0-20180430122841-1fe6dd257bc8
	github.com/liudanking/goutil v0.0.0-20190930035630-6830c8a1be22
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
	github.com/pinzolo/casee v0.0.0-20191019093852-17765ba5eb57
	xorm.io/core v0.7.2
	xorm.io/xorm v0.8.1
)

replace fyne.io/fyne => ./vendor/fyne.io/fyne
