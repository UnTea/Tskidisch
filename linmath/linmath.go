package linmath

import "math"

const (
	Epsilon float64 = 1e-3
)

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Norm() Vector {
	invert := 1.0 / v.Length()
	return MulOnScalar(v, invert)
}

func (v Vector) Negative() Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func Mul(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X * v2.X,
		Y: v1.Y * v2.Y,
		Z: v1.Z * v2.Z,
	}
}

func MulOnScalar(v Vector, scalar float64) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

func Dot(v1 Vector, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Add(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func Sub(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func Div(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X / v2.X,
		Y: v1.Y / v2.Y,
		Z: v1.Z / v2.Z,
	}
}

func DivOnScalar(v Vector, scalar float64) Vector {
	return Vector{
		X: v.X / scalar,
		Y: v.Y / scalar,
		Z: v.Z / scalar,
	}
}

func (v Vector) Clamp(min, max float64) Vector {
	return Vector{
		X: math.Min(math.Max(v.X, min), max),
		Y: math.Min(math.Max(v.Y, min), max),
		Z: math.Min(math.Max(v.Z, min), max),
	}
}

func Splat(scalar float64) Vector {
	return Vector{
		X: scalar,
		Y: scalar,
		Z: scalar,
	}
}

func Radians(degrees float64) float64 {
	return math.Pi * degrees / 180
}
