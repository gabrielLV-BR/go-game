package loaders

import (
	"bufio"
	"fmt"
	"gabriellv/game/structs"
	"os"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
)

type OBJLoader struct {}

type ObjData struct {
	positions []mgl32.Vec3
	normals []mgl32.Vec3
	uvs []mgl32.Vec2
	faces []faceData

	Vertices []structs.Vertex
	Indices []uint32
}

type faceData struct {
	positions [3]uint32
	normals [3]uint32
	uvs [3]uint32
}

func parseVec3(line string) mgl32.Vec3 {
	vec3 := mgl32.Vec3{}

	formatString := "v %f %f %f"

	if line[1] == 'n' {
		formatString = "vn %f %f %f"
	}

	n, err := fmt.Sscanf(line, formatString, &vec3[0], &vec3[1], &vec3[2])

	if err != nil || n != 3 {
		panic(err)
	}

	return vec3
}

func parseVec2(line string) mgl32.Vec2 {
	vec2 := mgl32.Vec2{}

	n, err := fmt.Sscanf(line, "vt %f %f", &vec2[0], &vec2[1])

	if err != nil || n != 2 {
		panic(err)
	}

	return vec2
}

func parseFace(line string) faceData {
	data := faceData{}

	formatStr := "f %d/%d/%d %d/%d/%d %d/%d/%d"

	n, err := fmt.Sscanf(line, formatStr,
		&data.positions[0], &data.uvs[0], &data.normals[0],
		&data.positions[1], &data.uvs[1], &data.normals[1],
		&data.positions[2], &data.uvs[2], &data.normals[2],
	)

	if err != nil || n != 9 {
		panic(err)
	}

	return data
}

type triple struct {
	a, b, c uint32
}

func (data *ObjData) buildMesh() {
	index := uint32(0)
	vertexToIndex := make(map[triple]uint32)

	for _, face := range data.faces {
		for i := 0; i < 3; i++ {
			position_index := face.positions[i] - 1
			normal_index := face.normals[i] - 1
			uv_index := face.uvs[i] - 1

			v := structs.Vertex {
				Position: data.positions[position_index],
				Normal: data.normals[normal_index],
				UV: data.uvs[uv_index],
			}

			triple := triple {
				a: position_index,
				b: normal_index,
				c: uv_index,
			}

			indx, ok := vertexToIndex[triple]

			if ok {
				// already in, just add index
				data.Indices = append(data.Indices, indx)
			} else {
				vertexToIndex[triple] = index
				data.Vertices = append(data.Vertices, v)
				data.Indices = append(data.Indices, index)
				index += 1
			}
		}
	}
}

func LoadObj(path string) ObjData {
	contents, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return ParseObj(string(contents))
}

func ParseObj(contents string) ObjData {
	data := ObjData{
		positions: []mgl32.Vec3{},
		normals: []mgl32.Vec3{},
		uvs: []mgl32.Vec2{},
		faces: []faceData{},

		Vertices: []structs.Vertex{},
		Indices: []uint32{},
	}

	scanner := bufio.NewScanner(strings.NewReader(contents))

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "v "):
			data.positions = append(data.positions, parseVec3(line))
		case strings.HasPrefix(line, "vn"):
			data.normals = append(data.normals, parseVec3(line))
		case strings.HasPrefix(line, "vt"):
			data.uvs = append(data.uvs, parseVec2(line))
		case strings.HasPrefix(line, "f "):
			data.faces = append(data.faces, parseFace(line))
		}
	}

	data.buildMesh()

	data.positions = nil
	data.normals = nil
	data.uvs = nil
	data.faces = nil

	return data
}
