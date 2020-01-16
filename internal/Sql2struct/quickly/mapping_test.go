package quickly

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var sql = `CREATE TABLE pro_account_group (
  id int unsigned AUTO_INCREMENT NOT NULL  COMMENT 'ddd',
  title varchar(50)   COMMENT '分组备注',
  weight tinyint(1) DEFAULT '0' COMMENT '权重（0高级 1理事 2普通级）',
  PRIMARY KEY (id)
) engine=InnoDB DEFAULT charset=utf8mb4 COMMENT='会员分组';`

func Test_MatchStmt(t *testing.T) {
	fmt.Println(sql)
	blocks, _ := MatchStmt(strings.NewReader(sql))
	for i := range blocks {
		t := HandleStmtBlock(blocks[i])
		t.GenType(os.Stdout)
	}
}
