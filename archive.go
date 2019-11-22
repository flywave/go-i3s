package i3s

type Mesh struct{}

type Cloud struct{}

type Archive interface {
}

type MeshArchive interface {
	Archive
	LoadNodeIndex(dir string, path *string) (error, *Node)
	LoadRootNodeIndex() (error, *Node)
	LoadTree() (error, *Tree)
	LoadNodes() (error, []Node)
	LoadMesh(n *Node) (error, *Mesh)
	LoadTexture(n *Node, index int) (error, []byte)
	GetTexturePath(n *Node, index int) (error, string)
}

type CloudArchive interface {
	LoadNodePage(index int) (error, *NodePage)
	LoadTree() (error, *PcslTree)
	LoadNodes() (error, []int)
	LoadGeometry(n *NodePage) (error, *Cloud)
}
