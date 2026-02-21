package hcl

type HclElementType int

const (
	HclElementTypeBlock = iota
	HclElementTypeByte
	HclElementTypeComment
	HclElementTypeWhitespace
)

var hclElementName = map[HclElementType]string{
	HclElementTypeByte:       "byte",
	HclElementTypeWhitespace: "whitespace",
	HclElementTypeComment:    "comment",
	HclElementTypeBlock:      "block",
}

type HclCommentAccessor interface {
	Comment() bool
}

type HclNameAccessor interface {
	Name() string
}

type HclValueAccessor interface {
	Value() string
}

type HclCommentMutator interface {
	SetComment(value bool)
}

type HclNameMutator interface {
	SetName(value string)
}

type HclValueMutator interface {
	SetValue(value string)
}

type HclBodyAccessor interface {
	Body() []HclElement
}

type HclBodyMutator interface {
	SetBody(value []HclElement)
}

type HclPairAccessor interface {
	Pair() HclElement
}

type HclPairMutator interface {
	SetPair(value HclElement)
}

type HclDir interface {
	HclName
	Files() ([]HclFile, error)
}

type HclBody interface {
	HclBodyAccessor
	HclBodyMutator
}

type HclComment interface {
	HclCommentAccessor
	HclCommentMutator
}

type HclValue interface {
	HclValueAccessor
	HclValueMutator
}

type HclName interface {
	HclNameAccessor
	HclNameMutator
}

type HclPair interface {
	HclPairAccessor
	HclPairMutator
}

// In the context of this parser, a file is content that consists of elements
// which may have subelements.
//
// In the context of HCL the basic elements are blocks and attributes.
//
// Whitespace and comments are also elements.
type HclElement interface {
	HclBody
	HclComment
	HclName
	HclPair
	HclValue

	AddNestedHclElement(element HclElement)
	File() HclFile
	HclElementType() HclElementType
	NestedElements() []HclElement
	SetFileName(value string)
}

type HclElementProvider interface {
	Elements() []HclElement
}

type HclFile interface {
	HclName
	HclElementProvider
}

type HclParser interface {
	NewFile(name string) HclFile
	NewDir(name string) HclDir
}

func (id HclElementType) String() string {
	return hclElementName[id]
}
