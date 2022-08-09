package greet

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func (h handler) To(cmd *cobra.Command, args []string) {
	h.logger.Info("命令调用成功", zap.String("use", cmd.Use))
	fmt.Printf("Hello, %s\n", args[0])
}
