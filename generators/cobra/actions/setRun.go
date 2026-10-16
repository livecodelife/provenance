	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return run{{command|pascal}}(cmd.OutOrStdout(), cmd, args)
	}
