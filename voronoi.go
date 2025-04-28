package geom

import "github.com/pzsz/voronoi"

func (c *Canvas) DrawVoronoiEdges(d *voronoi.Diagram) {
	for _, edge := range d.Edges {
		c.DrawLine(edge.Va.X, edge.Va.Y, edge.Vb.X, edge.Vb.Y)
	}
}

func (c *Canvas) DrawVoronoiPoints(d *voronoi.Diagram) {
	for _, cell := range d.Cells {
		c.DrawPoint(cell.Site.X, cell.Site.Y, 1)
	}
}

func (c *Canvas) DrawDelone(d *voronoi.Diagram) {
	for _, cell := range d.Cells {
		for _, halfedge := range cell.Halfedges {
			edge := halfedge.Edge
			if edge.LeftCell == nil || edge.RightCell == nil {
				continue
			}
			c.DrawLine(edge.LeftCell.Site.X, edge.LeftCell.Site.Y, edge.RightCell.Site.X, edge.RightCell.Site.Y)
		}
	}
}
