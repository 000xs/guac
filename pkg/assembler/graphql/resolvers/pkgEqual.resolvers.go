package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestPkgEqual is the resolver for the ingestPkgEqual field.
func (r *mutationResolver) IngestPkgEqual(ctx context.Context, pkg model.IDorPkgInput, otherPackage model.IDorPkgInput, pkgEqual model.PkgEqualInputSpec) (string, error) {
	return r.Backend.IngestPkgEqual(ctx, pkg, otherPackage, pkgEqual)
}

// IngestPkgEquals is the resolver for the ingestPkgEquals field.
func (r *mutationResolver) IngestPkgEquals(ctx context.Context, pkgs []*model.IDorPkgInput, otherPackages []*model.IDorPkgInput, pkgEquals []*model.PkgEqualInputSpec) ([]string, error) {
	funcName := "IngestPkgEquals"
	ingestedHashEqualsIDS := []string{}
	if len(pkgs) != len(otherPackages) {
		return ingestedHashEqualsIDS, gqlerror.Errorf("%v :: uneven packages and other packages for ingestion", funcName)
	} else if len(pkgs) != len(pkgEquals) {
		return ingestedHashEqualsIDS, gqlerror.Errorf("%v :: uneven packages and pkgEquals for ingestion", funcName)
	}

	return r.Backend.IngestPkgEquals(ctx, pkgs, otherPackages, pkgEquals)
}

// PkgEqual is the resolver for the PkgEqual field.
func (r *queryResolver) PkgEqual(ctx context.Context, pkgEqualSpec model.PkgEqualSpec) ([]*model.PkgEqual, error) {
	if len(pkgEqualSpec.Packages) > 2 {
		return nil, gqlerror.Errorf("PkgEqual :: too many packages in query, max 2, got: %v", len(pkgEqualSpec.Packages))
	}
	return r.Backend.PkgEqual(ctx, &pkgEqualSpec)
}
