package infrastructure

import(
    "database/sql"
    // for mysql driver
    _ "github.com/go-sql-driver/mysql"
    "jdodge-go/interfaces"
    "fmt"
)

type MysqlHandler struct {
    Conn *sql.DB
}

type MysqlRow struct {
    Rows *sql.Rows
}

func (handler *MysqlHandler) Execute(statement string) {
    handler.Conn.Exec(statement)
}

func (handler *MysqlHandler) Query(statement string) interfaces.Row {
    rows, err := handler.Conn.Query(statement)
    if err != nil {
        fmt.Errorf("query error: %s", err)
        return new(MysqlRow)
    }
    row := new(MysqlRow)
    row.Rows = rows
    return row
}

func (row MysqlRow) Scan(dest ...interface{}) error {
    err := row.Rows.Scan(dest...)
    if err != nil {
        return err
    }
    return nil
}

func (row MysqlRow) Next() bool {
    return row.Rows.Next()
}

func NewMysqlHandler(dbfilename string) *MysqlHandler {
    db, err := sql.Open("mysql", dbfilename)
    if err != nil {
        fmt.Errorf("sql open error: %s", err)
        return nil
    }
    handler := new(MysqlHandler)
    handler.Conn = db
    return handler
}
