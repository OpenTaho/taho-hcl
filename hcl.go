package hcl

type HclElementType int

const (
	HclElementTypeWhitespace = iota
	HclElementTypeComment
	HclElementTypeBlock
)

var hclElementName = map[HclElementType]string{
	HclElementTypeWhitespace: "whitespace",
	HclElementTypeComment:    "comment",
	HclElementTypeBlock:      "block",
}

type HclNameAccessor interface {
	Name() string
}

type HclNameMutator interface {
	SetName(value string)
}

type HclName interface {
	HclNameAccessor
	HclNameMutator
}

type HclDir interface {
	HclName
	Files() []HclFile
}

type HclElement interface {
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
	ReadFile(name string) HclFile
	ReadDir(name string) HclDir
}

// Read an HCL FILE
func ReadFile(name string, parser HclParser) HclFile {
	return parser.ReadFile(name)
}

// Read an HCL FILE
func ReadDir(name string, parser HclParser) HclDir {
	return parser.ReadDir(name)
}

func (id HclElementType) String() string {
	return hclElementName[id]
}
