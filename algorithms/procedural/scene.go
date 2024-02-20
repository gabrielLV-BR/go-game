package procedural

import (
	"gabriellv/game/core"
	"gabriellv/game/structs"
	"math/rand"
	"sort"

	"github.com/go-gl/mathgl/mgl32"
)

const MAX_ROOM_WIDTH int = 10
const MAX_ROOM_HEIGHT int = 2
const MAX_ROOM_DEPTH int = 10

type LevelGenerator struct {
	RoomCount int
}

type cube struct {
	Position mgl32.Vec3
	Width    int
	Depth    int
}

func (level *LevelGenerator) Generate() core.Mesh {
	// first, create a grid to represent world

	roomGraph := structs.Graph[cube]{}
	roomGraph.New()

	for i := 0; i < level.RoomCount; i++ {
		x := rand.Float32()
		y := rand.Float32()
		z := float32(0.0)

		w := rand.Intn(MAX_ROOM_WIDTH)
		d := rand.Intn(MAX_ROOM_DEPTH)

		point := cube{
			Position: mgl32.Vec3{x, y, z},
			Width:    w,
			Depth:    d,
		}

		roomGraph.AddNode(point)
	}

	// link every node
	// very dumb and slow but will do the trick

	for i1, node1 := range roomGraph.Nodes {

		var closestDistance float32
		var closestNode int

		for i2, node2 := range roomGraph.Nodes {
			if i1 == i2 {
				continue
			}

			dist := node1.Position.Sub(node2.Position).Len()

			if closestDistance > dist {
				closestDistance = dist
				closestNode = i2
			}
		}

		roomGraph.AddDirectedEdge(i1, closestNode)
	}

	// cube-ify world
	grid := structs.Grid3D[bool]{}
	grid.New(50, 10, 50)

	// I'm gonna use the marching cubes algorithm to make the map
	// first, we place the rooms in the grid
	placeRoomsInGrid(&grid, &roomGraph)
	// placeEdgesInGrid(&grid, &roomGraph)

	// build room meesh

	return MarchingCubes(
		grid,
		mgl32.Vec3{
			float32(MAX_ROOM_WIDTH),
			float32(MAX_ROOM_HEIGHT),
			float32(MAX_ROOM_DEPTH),
		},
	)
}

func placeRoomsInGrid(grid *structs.Grid3D[bool], roomGraph *structs.Graph[cube]) {
	for _, node := range roomGraph.Nodes {
		x := int(node.Position.X())
		y := int(node.Position.Y())
		z := int(node.Position.Z())

		for xx := 0; xx < (x + node.Width); xx++ {
			for zz := 0; zz < (z + node.Depth); zz++ {
				for yy := 0; yy < (y + MAX_ROOM_HEIGHT); yy++ {
					grid.Place(true, xx, yy, zz)
				}
			}
		}
	}
}

type tuple [2]int

func placeEdgesInGrid(grid *structs.Grid3D[bool], roomGraph *structs.Graph[cube]) {
	set := make(map[tuple]bool)

	for k, v := range roomGraph.Edges {
		fromIndex := k

		for _, toIndex := range v {
			if fromIndex == toIndex {
				continue
			}

			directionTuple := [2]int{fromIndex, toIndex}

			// check to see if we've already connected these nodes
			// we gotta do this because our room edging algorithm
			// places double edges for A -> B and B -> A
			sort.Ints(directionTuple[:])
			if _, ok := set[directionTuple]; ok {
				continue
			}

			a := roomGraph.Nodes[fromIndex]
			b := roomGraph.Nodes[toIndex]

			direction := a.Position.Sub(b.Position)
			nDirection := direction.Normalize()

			for direction.Len() > 0 {
				x := int(direction.X())
				y := int(direction.Y())

				for z := 0; z < MAX_ROOM_HEIGHT; z++ {
					grid.Place(true, x, y, z)
				}

				direction = direction.Sub(nDirection)
			}

			set[directionTuple] = true
		}
	}
}
