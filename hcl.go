package hcl

type HclType int

const (
	HclTypeBlock = iota
	HclTypeComment
	HclTypeDir
	HclTypeFile
	HclTypeHereDoc
	HclTypeHereDocWithIndent
	HclTypeMultiLinePair
	HclTypeOther
	HclTypeSingleLinePair
	HclTypeSpace
	HclTypeString
	HclTypeToken
)

var hclElementName = map[HclType]string{
	HclTypeBlock:             "block",
	HclTypeComment:           "comment",
	HclTypeDir:               "dir",
	HclTypeFile:              "file",
	HclTypeHereDoc:           "doc",
	HclTypeHereDocWithIndent: "doc-with-indent",
	HclTypeMultiLinePair:     "multi-line-pair",
	HclTypeOther:             "other",
	HclTypeSingleLinePair:    "single-line-pair",
	HclTypeSpace:             "space",
	HclTypeString:            "string",
	HclTypeToken:             "token",
}

type HclNameAccessor interface {
	Name() string
}

type HclValueAccessor interface {
	Value() string
}

type HclNameMutator interface {
	SetName(value string)
}

type HclValueMutator interface {
	SetValue(value string)
}

type HclBodyAccessor interface {
	Body() []HclNode
}

type HclBodyMutator interface {
	SetBody(value []HclNode)
}

type HclPairAccessor interface {
	Pair() HclNode
}

type HclPairMutator interface {
	SetPair(value HclNode)
}

type HclDir interface {
	HclName
	HclNodeProvider
	Files() ([]HclFile, error)
}

type HclBody interface {
	HclBodyAccessor
	HclBodyMutator
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
type HclNode interface {
	HclBody
	HclPair
	HclValue

	AddNode(element HclNode)
	File() HclFile
	Nodes() []HclNode
	Operator() string
	SetFileName(value string)
	Type() HclType
}

type HclNodeProvider interface {
	Nodes() []HclNode
}

type HclFile interface {
	HclName
	HclNodeProvider
}

type HclParser interface {
	NewFile(name string) HclFile
	NewDir(name string) HclDir
}

func (id HclType) String() string {
	return hclElementName[id]
}
