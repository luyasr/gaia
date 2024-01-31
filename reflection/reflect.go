package reflection

func SetUp(obj any) error {
	if err := SetDefaultTag(obj); err != nil {
		return err
	}

	return nil
}
