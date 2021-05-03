package linmath

import "math"

const (
	Epsilon float64 = 1e-5
)

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Norm() Vector {
	invert := 1.0 / v.Length()
	return v.MulOnScalar(invert)
}

func (v Vector) Negative() Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vector) Mul(v2 Vector) Vector {
	return Vector{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
		Z: v.Z * v2.Z,
	}
}

func (v Vector) MulOnScalar(scalar float64) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}

func (v Vector) Div(v2 Vector) Vector {
	return Vector{
		X: v.X / v2.X,
		Y: v.Y / v2.Y,
		Z: v.Z / v2.Z,
	}
}

func (v Vector) DivOnScalar(scalar float64) Vector {
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

func (v Vector) Pow(scalar float64) Vector {
	return Vector{
		X: math.Pow(v.X, scalar),
		Y: math.Pow(v.Y, scalar),
		Z: math.Pow(v.Z, scalar),
	}
}

func Radians(degrees float64) float64 {
	return math.Pi * degrees / 180
}

func Splat(scalar float64) Vector {
	return Vector{
		X: scalar,
		Y: scalar,
		Z: scalar,
	}
}
