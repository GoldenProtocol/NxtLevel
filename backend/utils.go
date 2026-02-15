package main

func FilterNodesByType(nodes []*NodeModel, nodeType string) []*NodeModel {
	out := make([]*NodeModel, 0, len(nodes))

	for _, n := range nodes {
		if n.Type == nodeType {
			out = append(out, n)
		}
	}

	return out
}

func ExtractNodeId(nodes []*NodeModel) []string {
	out := make([]string, 0, len(nodes))

	for _, n := range nodes {
		out = append(out, n.ID)
	}

	return out
}
