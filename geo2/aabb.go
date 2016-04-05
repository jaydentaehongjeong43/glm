package geo2

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/math"
)

// AABB is an axis-aligned bounding box.
type AABB struct {
	Center glm.Vec2
	Radius glm.Vec2
}

// TestAABBAABB returns true if these AABB overlap.
func TestAABBAABB(a *AABB, b *AABB) bool {
	if math.Abs(a.Center[0]-b.Center[0]) > (a.Radius[0] + b.Radius[0]) {
		return false
	}

	if math.Abs(a.Center[1]-b.Center[1]) > (a.Radius[1] + b.Radius[1]) {
		return false
	}

	return true
}

// UpdateAABB computes an enclosing AABB base transformed by t and puts the
// result in fill.
func UpdateAABB(base, fill *AABB, t *glm.Mat2x3) {
	for i := 0; i < 2; i++ {
		fill.Center[i] = t[i+4]
		fill.Radius[i] = 0
		for j := 0; j < 2; j++ {
			fill.Center[i] += t[j*2+i] * base.Center[j]
			fill.Radius[i] += math.Abs(t[j*2+i]) * base.Radius[j]
		}
	}
}

// ClosestPointAABBPoint returns the point in or on the AABB closest to 'p'
func ClosestPointAABBPoint(a *AABB, p *glm.Vec2) glm.Vec2 {
	return glm.Vec2{
		math.Clamp(p[0], a.Center[0]-a.Radius[0], a.Center[0]+a.Radius[0]),
		math.Clamp(p[1], a.Center[1]-a.Radius[1], a.Center[1]+a.Radius[1]),
	}
}

// SqDistAABBPoint returns the square distance of p to the AABB
func SqDistAABBPoint(a *AABB, p *glm.Vec2) float32 {
	var sqDist float32

	// For each axis count any excess distance outside box extents
	v := p[0]
	min := a.Center[0] - a.Radius[0]
	max := a.Center[0] + a.Radius[0]
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	v = p[1]
	min = a.Center[1] - a.Radius[1]
	max = a.Center[1] + a.Radius[1]
	if v < min {
		sqDist += (min - v) * (min - v)
	}
	if v > max {
		sqDist += (v - max) * (v - max)
	}

	return sqDist
}