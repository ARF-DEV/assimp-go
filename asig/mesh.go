package asig

import (
	"github.com/go-gl/mathgl/mgl32"
)

const (
	MaxColorSets = 8
	MaxTexCoords = 8
)

type Mesh struct {

	//Bitwise combination of PrimitiveType enum
	PrimitiveTypes PrimitiveType
	Vertices       []mgl32.Vec3
	Normals        []mgl32.Vec3
	Tangents       []mgl32.Vec3
	BitTangents    []mgl32.Vec3

	//ColorSets vertex color sets where each set is either empty or has length=len(Vertices), with max number of sets=MaxColorSets
	ColorSets [MaxColorSets][]mgl32.Vec4

	//TexCoords (aka UV channels) where each TexCoords[i] has NumUVComponents[i] channels, and is either empty or has length=len(Vertices), with max number of TexCoords per vertex = MaxTexCoords
	TexCoords            [MaxTexCoords][]mgl32.Vec3
	TexCoordChannelCount [MaxTexCoords]uint

	Faces       []Face
	Bones       []*Bone
	AnimMeshes  []*AnimMesh
	AABB        AABB
	MorphMethod MorphMethod

	MaterialIndex uint
	Name          string
}

type Face struct {
	Indices []uint
}

type AnimMesh struct {
	Name string

	/** Replacement for Mes.Vertices. If this array is non-NULL,
	 *  it *must* contain mNumVertices entries. The corresponding
	 *  array in the host mesh must be non-NULL as well - animation
	 *  meshes may neither add or nor remove vertex components (if
	 *  a replacement array is NULL and the corresponding source
	 *  array is not, the source data is taken instead)*/
	Vertices    []mgl32.Vec3
	Normals     []mgl32.Vec3
	Tangents    []mgl32.Vec3
	BitTangents []mgl32.Vec3
	Colors      [MaxColorSets][]mgl32.Vec4
	TexCoords   [MaxTexCoords][]mgl32.Vec3

	Weight float32
}

type AABB struct {
	Min mgl32.Vec3
	Max mgl32.Vec3
}

type Bone struct {
	Name string
	//The influence weights of this bone
	Weights []VertexWeight

	/** Matrix that transforms from bone space to mesh space in bind pose.
	 *
	 * This matrix describes the position of the mesh
	 * in the local space of this bone when the skeleton was bound.
	 * Thus it can be used directly to determine a desired vertex position,
	 * given the world-space transform of the bone when animated,
	 * and the position of the vertex in mesh space.
	 *
	 * It is sometimes called an inverse-bind matrix,
	 * or inverse bind pose matrix.
	 */
	OffsetMatrix mgl32.Mat4
}

type VertexWeight struct {
	VertIndex uint
	//The strength of the influence in the range (0...1). The total influence from all bones at one vertex is 1
	Weight float32
}
