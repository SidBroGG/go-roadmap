package main

// Structures

type AssetBase struct {
	ID   int
	Path string
}

type Texture struct {
	AssetBase
	Resolution string
	Format     string
}

type Mesh struct {
	AssetBase
	VertexCount int
	IsAnimated  bool
}

type SoundEffect struct {
	AssetBase
	Duration float32
	Loopable bool
}

// Interfaces

type Loadable interface {
	Load(path string) error
}

type Renderable interface {
	Render(engine string) string
}

type Updatable interface {
	UpdateVersion(v float32)
}

type DebugInfo interface {
	GetInfo() string
}

type EngineAsset interface {
	Loadable
	Renderable
}
