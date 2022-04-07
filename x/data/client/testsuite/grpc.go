package testsuite

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/regen-network/regen-ledger/x/data"
)

func (s *IntegrationTestSuite) TestQueryByIRI() {
	val := s.network.Validators[0]

	iri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"invalid IRI",
			fmt.Sprintf("%s/regen/data/v1/by-iri/%s", val.APIAddress, "foo"),
			true,
			"not found",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/by-iri/%s", val.APIAddress, iri),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var entry data.QueryByIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &entry)

			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
				require.NotNil(entry.Entry)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryByAttestor() {
	val := s.network.Validators[0]

	acc1, err := val.ClientCtx.Keyring.Key("acc1")
	s.Require().NoError(err)

	addr := acc1.GetAddress().String()

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid attestor",
			fmt.Sprintf("%s/regen/data/v1/by-attestor/%s", val.APIAddress, "foo"),
			true,
			"invalid bech32 string",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/by-attestor/%s", val.APIAddress, addr),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/by-attestor/%s?pagination.limit=1", val.APIAddress, addr),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var entries data.QueryByAttestorResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &entries)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(entries.Entries)
				require.Len(entries.Entries, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryHashByIRI() {
	val := s.network.Validators[0]

	iri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	testCases := []struct {
		name   string
		url    string
		expErr bool
		errMsg string
	}{
		{
			"invalid IRI",
			fmt.Sprintf("%s/regen/data/v1/hash/%s", val.APIAddress, "foo"),
			true,
			"invalid IRI",
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/hash/%s", val.APIAddress, iri),
			false,
			"",
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var contentHash data.QueryHashByIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &contentHash)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(contentHash.ContentHash)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAttestors() {
	val := s.network.Validators[0]

	iri := "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid attestor",
			fmt.Sprintf("%s/regen/data/v1/attestors/%s", val.APIAddress, "foo"),
			true,
			"not found",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/attestors/%s", val.APIAddress, iri),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/attestors/%s?pagination.limit=1", val.APIAddress, iri),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var attestors data.QueryAttestorsByIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &attestors)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(attestors.Attestors)
				require.Len(attestors.Attestors, tc.expItems)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolverInfo() {
	val := s.network.Validators[0]

	url := "https://foo.bar"

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid url",
			fmt.Sprintf("%s/regen/data/v1/resolver?url=%s", val.APIAddress, "foo"),
			true,
			"not found",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/resolver?url=%s", val.APIAddress, url),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/resolver?url=%s&pagination.limit=1", val.APIAddress, url),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var resolver data.QueryResolverInfoResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &resolver)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(resolver.Id)
				require.NotNil(resolver.Manager)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryResolvers() {
	val := s.network.Validators[0]

	iri := s.iri

	testCases := []struct {
		name     string
		url      string
		expErr   bool
		errMsg   string
		expItems int
	}{
		{
			"invalid iri",
			fmt.Sprintf("%s/regen/data/v1/resolvers/%s", val.APIAddress, "foo"),
			true,
			"not found",
			0,
		},
		{
			"valid request",
			fmt.Sprintf("%s/regen/data/v1/resolvers/%s", val.APIAddress, iri),
			false,
			"",
			2,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/regen/data/v1/resolvers/%s?pagination.limit=1", val.APIAddress, iri),
			false,
			"",
			1,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var resolvers data.QueryResolversByIRIResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &resolvers)

			if tc.expErr {
				require.Error(err)
				require.Contains(string(resp), tc.errMsg)
			} else {
				require.NoError(err)
				require.NotNil(resolvers.ResolverUrls)
				require.Len(resolvers.ResolverUrls, tc.expItems)
			}
		})
	}
}
