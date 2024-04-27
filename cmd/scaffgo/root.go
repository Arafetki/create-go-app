package scaffgo

import "github.com/spf13/cobra"

var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "scaffgo",
	Short: "A generator for Golang based applications",
	Long: `
 ____  ____  ____  _____ _____ ____  _     ____    _____ ____    ____  ____  ____  _     _  ____  ____  _____  _  ____  _      ____ 
/ ___\/   _\/  _ \/    //    //  _ \/ \   /  _ \  /  __//  _ \  /  _ \/  __\/  __\/ \   / \/   _\/  _ \/__ __\/ \/  _ \/ \  /|/ ___\
|    \|  /  | / \||  __\|  __\| / \|| |   | | \|  | |  _| / \|  | / \||  \/||  \/|| |   | ||  /  | / \|  / \  | || / \|| |\ |||    \
\___ ||  \_ | |-||| |   | |   | \_/|| |_/\| |_/|  | |_//| \_/|  | |-|||  __/|  __/| |_/\| ||  \_ | |-||  | |  | || \_/|| | \||\___ |
\____/\____/\_/ \|\_/   \_/   \____/\____/\____/  \____\\____/  \_/ \|\_/   \_/   \____/\_/\____/\_/ \|  \_/  \_/\____/\_/  \|\____/
                                                                                                                                    

-> Focus on writing code and thinking of business logic!
<- The Scaffold Go App CLI will take care of the rest.
`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
}
