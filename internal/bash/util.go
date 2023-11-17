package bash

// TODO should happen in basm
func funcName(namespace, name string) string {
	if namespace == "" {
		return name
	}

	return namespace + "__" + name
}
