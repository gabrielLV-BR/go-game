package loaders

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
)

type Object struct {
	Vertices []float32
	Indices  []uint32
}

func LoadObj(path string) (Object, error) {
	contents, err := os.ReadFile(path)

	if err != nil {
		return Object{}, err
	}

	return ParseObj(string(contents))
}

func ParseObj(contents string) (Object, error) {
	object := Object{}

	data := objData{
		positions: []mgl32.Vec3{},
		normals:   []mgl32.Vec3{},
		uvs:       []mgl32.Vec2{},

		position_indices: []uint32{},
		normal_indices:   []uint32{},
		uv_indices:       []uint32{},
	}

	scanner := bufio.NewScanner(strings.NewReader(contents))

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "v "):
			{
				vec, err := parseVec3(line)
				if err != nil {
					return object, err
				}

				data.positions = append(data.positions, vec)
			}
		case strings.HasPrefix(line, "vn"):
			{
				vec, err := parseVec3(line)
				if err != nil {
					return object, err
				}

				data.normals = append(data.normals, vec)
			}
		case strings.HasPrefix(line, "vt"):
			{
				vec, err := parseVec2(line)
				if err != nil {
					return object, err
				}

				data.uvs = append(data.uvs, vec)
			}
		case strings.HasPrefix(line, "f "):
			{
				face, err := parseFace(line)
				if err != nil {
					return object, err
				}

				data.position_indices = append(data.position_indices, face.positions[0], face.positions[1], face.positions[2])
				data.normal_indices = append(data.normal_indices, face.normals[0], face.normals[1], face.normals[2])
				data.uv_indices = append(data.uv_indices, face.uvs[0], face.uvs[1], face.uvs[2])
			}
		}
	}

	object, err := data.buildMesh()

	// try to indicate to GC that he can free these
	data.positions = nil
	data.normals = nil
	data.uvs = nil

	data.position_indices = nil
	data.normal_indices = nil
	data.uv_indices = nil

	return object, err
}

type objData struct {
	positions []mgl32.Vec3
	normals   []mgl32.Vec3
	uvs       []mgl32.Vec2

	position_indices []uint32
	normal_indices   []uint32
	uv_indices       []uint32
}

type faceData struct {
	positions [3]uint32
	normals   [3]uint32
	uvs       [3]uint32
}

func parseVec3(line string) (mgl32.Vec3, error) {
	vec3 := mgl32.Vec3{}

	formatString := "v %f %f %f"

	if line[1] == 'n' {
		formatString = "vn %f %f %f"
	}

	_, err := fmt.Sscanf(line, formatString, &vec3[0], &vec3[1], &vec3[2])

	return vec3, err
}

func parseVec2(line string) (mgl32.Vec2, error) {
	vec2 := mgl32.Vec2{}

	_, err := fmt.Sscanf(line, "vt %f %f", &vec2[0], &vec2[1])

	return vec2, err
}

func parseFace(line string) (faceData, error) {
	data := faceData{}

	formatStr := "f %d/%d/%d %d/%d/%d %d/%d/%d"

	_, err := fmt.Sscanf(line, formatStr,
		&data.positions[0], &data.uvs[0], &data.normals[0],
		&data.positions[1], &data.uvs[1], &data.normals[1],
		&data.positions[2], &data.uvs[2], &data.normals[2],
	)

	return data, err
}

// used in the mapping
type triple struct {
	a, b, c uint32
}

func (data *objData) buildMesh() (Object, error) {
	object := Object{
		Vertices: []float32{},
		Indices:  []uint32{},
	}

	index := uint32(0)
	vertexToIndex := make(map[triple]uint32)

	if len(data.position_indices) != len(data.normal_indices) ||
		len(data.position_indices) != len(data.uv_indices) {
		return object, errors.New("Invalid mesh indices")
	}

	for i := range data.position_indices {
		position_index := data.position_indices[i] - 1
		normal_index := data.normal_indices[i] - 1
		uv_index := data.uv_indices[i] - 1

		triple := triple{
			position_index, normal_index, uv_index,
		}

		indx, ok := vertexToIndex[triple]

		if ok {
			// already in, just add index
			object.Indices = append(object.Indices, indx)
		} else {

			position := data.positions[position_index]
			normal := data.normals[normal_index]
			uv := data.uvs[uv_index]

			vertexToIndex[triple] = index
			object.Vertices = append(object.Vertices,
				position[0], position[1], position[2],
				normal[0], normal[1], normal[2],
				uv[0], uv[1],
			)
			object.Indices = append(object.Indices, index)

			index += 1
		}
	}

	return object, nil
}
