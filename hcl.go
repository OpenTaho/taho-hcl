package hcl

type HclType int

const (
	HclTypeBlock = iota
	HclTypeComment
	HclTypeDir
	HclTypeDoc
	HclTypeDocWithIndent
	HclTypeOther
	HclTypePair
	HclTypePairGroup
	HclTypeSpace
	HclTypeSpan
	HclTypeString
	HclTypeToken
)

var hclElementName = map[HclType]string{
	HclTypeBlock:         "block",
	HclTypeComment:       "comment",
	HclTypeDir:           "dir",
	HclTypeDoc:           "doc",
	HclTypeDocWithIndent: "doc-with-indent",
	HclTypeOther:         "other",
	HclTypePair:          "pair",
	HclTypePairGroup:     "pair-group",
	HclTypeSpace:         "space",
	HclTypeSpan:          "span",
	HclTypeString:        "string",
	HclTypeToken:         "token",
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

type HclCommentAccessor interface {
	Comment() string
}

type HclCommentMutator interface {
	SetComment(value string)
}

type HclPairAccessor interface {
	Pair() HclNode
}

type HclPairMutator interface {
	SetPair(value HclNode)
}

type HclDir interface {
	HclNameAccessor
	HclBodyAccessor
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
type HclNode interface {
	HclBody
	HclPair
	HclValue
	HclCommentMutator

	AddNode(element HclNode)
	File() HclFile
	IsSimplePair() bool
	Operator() string
	SetDocIndentation(value int)
	SetDocTag(value string)
	SetFileName(value string)
	SetType(value HclType)
	String() string
	Type() HclType
}

type HclFile interface {
	HclNameAccessor
	HclBodyAccessor
	Format() HclFile
	String() string
	SaveAs(fileName string)
}

type HclParser interface {
	NewFile(name string) HclFile
	NewDir(name string) HclDir
}

func (id HclType) String() string {
	return hclElementName[id]
}
