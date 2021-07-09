package main

import (
	"context"

	pb "github.com/bohdanstryber/route_guide/routeguide"
)

type routeGuideServer struct {
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (pb.Feature, error) {
	return nil
}
