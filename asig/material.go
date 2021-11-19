package asig

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L ./libs -l assimp_windows_amd64 -l IrrXML_windows_amd64 -l zlib_windows_amd64

#include <stdlib.h> //Needed for C.free

#include <assimp/scene.h>

//Functions
unsigned int aiGetMaterialTextureCount(const struct aiMaterial* pMat, enum aiTextureType type);

enum aiReturn aiGetMaterialTexture(
	const struct aiMaterial* mat,
    enum aiTextureType type,
    unsigned int  index,
    struct aiString* path,
    enum aiTextureMapping* mapping,
    unsigned int* uvindex,
    ai_real* blend,
    enum aiTextureOp* op,
    enum aiTextureMapMode* mapmode,
    unsigned int* flags);
*/
import "C"
import (
	"errors"
	"fmt"
)

type Material struct {
	cMat *C.struct_aiMaterial

	/** List of all material properties loaded. */
	Properties []*MaterialProperty

	/** Storage allocated */
	AllocatedStorage uint
}

type MaterialProperty struct {

	//Specifies the name of the property (aka key). Keys are generally case insensitive.
	name string

	/** Textures: Specifies their exact usage semantic.
	 * For non-texture properties, this member is always 0 (aka TextureTypeNone).
	 */
	Semantic TextureType

	/** Textures: Specifies the index of the texture.
	 *  For non-texture properties, this member is always 0.
	 */
	Index uint

	/** Type information for the property.
	 *
	 * Defines the data layout inside the data buffer. This is used
	 * by the library internally to perform debug checks and to
	 * utilize proper type conversions.
	 * (It's probably a hacky solution, but it works.)
	 */
	TypeInfo MatPropertyTypeInfo

	//Binary buffer to hold the property's value.
	Data []byte
}

func GetMaterialTextureCount(m *Material, texType TextureType) int {
	return int(C.aiGetMaterialTextureCount(m.cMat, uint32(texType)))
}

type GetMatTexInfo struct {
	Path string
}

func GetMaterialTexture(m *Material, texType TextureType, texIndex uint) (*GetMatTexInfo, error) {

	outCPath := &C.struct_aiString{}
	status := aiReturn(C.aiGetMaterialTexture(m.cMat, uint32(texType), C.uint(texIndex), outCPath, nil, nil, nil, nil, nil, nil))
	if status == aiReturnSuccess {
		return &GetMatTexInfo{
			Path: parseAiString(*outCPath),
		}, nil
	}

	if status == aiReturnFailure {
		return nil, errors.New("get texture failed: " + getAiErr().Error())
	}

	if status == aiReturnOutofMemory {
		return nil, errors.New("get texture failed: out of memory")
	}

	return nil, errors.New("get texture failed: unknown error with code " + fmt.Sprintf("%v", status))
}