package burrowdb

// Node represents a node in the store.
type Node struct {
	ID         uint64         `json:"id"`
	Labels     []string       `json:"labels"`
	Properties map[string]any `json:"properties"`
}

// Edge represents a edge/relationship in the store.
type Edge struct {
	ID         uint64         `json:"id"`
	Labels     []string       `json:"labels"`
	Properties map[string]any `json:"properties"`
	SourceID   uint64         `json:"source_id"`
	TargetID   uint64         `json:"target_id"`
}
