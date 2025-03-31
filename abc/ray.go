package abc

// Ray is a geometric shape that represents a set of points lying on a straiht line
// from the beginning (begin Point2D), and to the infinity in the
// direction of the point (direction Point2D)
type Ray struct {
	begin     Point2D
	direction Point2D
}

func NewRay(begin, direction Point2D) Ray {
	return Ray{begin: begin, direction: direction}
}
