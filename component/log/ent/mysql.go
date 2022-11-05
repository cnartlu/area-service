package ent

import "strings"

type mysql struct {
}

func (m *mysql) Write(p []byte) (int, error) {
	buf := strings.Builder{}
	buf.WriteString("INSERT INTO `")
	buf.WriteString("")
	buf.WriteString("`")
	buf.WriteString("()")
	buf.WriteString("")
	_ = buf.String()
	return 0, nil
}
