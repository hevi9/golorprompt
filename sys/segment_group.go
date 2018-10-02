package sys

type groupSegment struct{}

func init() {
	Register(
		"group",
		"Group sub segments",
		func() Segment {
			return &groupSegment{}
		},
	)
}

func (*groupSegment) Render(env Environment) []Chunk {
	return nil /* dummy */
}
