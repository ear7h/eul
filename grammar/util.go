package grammar

func TreeForEach(node AstNode, fn func(AstNode) error) error {
	fn(node)
	for _, v := range node.Children() {
		if err := TreeForEach(v, fn); err != nil {
			return err
		}
	}

	return nil
}
