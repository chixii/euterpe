// Code generated by counterfeiter. DO NOT EDIT.
package libraryfakes

import (
	"sync"

	"github.com/ironsmile/euterpe/src/library"
)

type FakeBrowser struct {
	BrowseAlbumsStub        func(library.BrowseArgs) ([]library.Album, int)
	browseAlbumsMutex       sync.RWMutex
	browseAlbumsArgsForCall []struct {
		arg1 library.BrowseArgs
	}
	browseAlbumsReturns struct {
		result1 []library.Album
		result2 int
	}
	browseAlbumsReturnsOnCall map[int]struct {
		result1 []library.Album
		result2 int
	}
	BrowseArtistsStub        func(library.BrowseArgs) ([]library.Artist, int)
	browseArtistsMutex       sync.RWMutex
	browseArtistsArgsForCall []struct {
		arg1 library.BrowseArgs
	}
	browseArtistsReturns struct {
		result1 []library.Artist
		result2 int
	}
	browseArtistsReturnsOnCall map[int]struct {
		result1 []library.Artist
		result2 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBrowser) BrowseAlbums(arg1 library.BrowseArgs) ([]library.Album, int) {
	fake.browseAlbumsMutex.Lock()
	ret, specificReturn := fake.browseAlbumsReturnsOnCall[len(fake.browseAlbumsArgsForCall)]
	fake.browseAlbumsArgsForCall = append(fake.browseAlbumsArgsForCall, struct {
		arg1 library.BrowseArgs
	}{arg1})
	stub := fake.BrowseAlbumsStub
	fakeReturns := fake.browseAlbumsReturns
	fake.recordInvocation("BrowseAlbums", []interface{}{arg1})
	fake.browseAlbumsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrowser) BrowseAlbumsCallCount() int {
	fake.browseAlbumsMutex.RLock()
	defer fake.browseAlbumsMutex.RUnlock()
	return len(fake.browseAlbumsArgsForCall)
}

func (fake *FakeBrowser) BrowseAlbumsCalls(stub func(library.BrowseArgs) ([]library.Album, int)) {
	fake.browseAlbumsMutex.Lock()
	defer fake.browseAlbumsMutex.Unlock()
	fake.BrowseAlbumsStub = stub
}

func (fake *FakeBrowser) BrowseAlbumsArgsForCall(i int) library.BrowseArgs {
	fake.browseAlbumsMutex.RLock()
	defer fake.browseAlbumsMutex.RUnlock()
	argsForCall := fake.browseAlbumsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBrowser) BrowseAlbumsReturns(result1 []library.Album, result2 int) {
	fake.browseAlbumsMutex.Lock()
	defer fake.browseAlbumsMutex.Unlock()
	fake.BrowseAlbumsStub = nil
	fake.browseAlbumsReturns = struct {
		result1 []library.Album
		result2 int
	}{result1, result2}
}

func (fake *FakeBrowser) BrowseAlbumsReturnsOnCall(i int, result1 []library.Album, result2 int) {
	fake.browseAlbumsMutex.Lock()
	defer fake.browseAlbumsMutex.Unlock()
	fake.BrowseAlbumsStub = nil
	if fake.browseAlbumsReturnsOnCall == nil {
		fake.browseAlbumsReturnsOnCall = make(map[int]struct {
			result1 []library.Album
			result2 int
		})
	}
	fake.browseAlbumsReturnsOnCall[i] = struct {
		result1 []library.Album
		result2 int
	}{result1, result2}
}

func (fake *FakeBrowser) BrowseArtists(arg1 library.BrowseArgs) ([]library.Artist, int) {
	fake.browseArtistsMutex.Lock()
	ret, specificReturn := fake.browseArtistsReturnsOnCall[len(fake.browseArtistsArgsForCall)]
	fake.browseArtistsArgsForCall = append(fake.browseArtistsArgsForCall, struct {
		arg1 library.BrowseArgs
	}{arg1})
	stub := fake.BrowseArtistsStub
	fakeReturns := fake.browseArtistsReturns
	fake.recordInvocation("BrowseArtists", []interface{}{arg1})
	fake.browseArtistsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrowser) BrowseArtistsCallCount() int {
	fake.browseArtistsMutex.RLock()
	defer fake.browseArtistsMutex.RUnlock()
	return len(fake.browseArtistsArgsForCall)
}

func (fake *FakeBrowser) BrowseArtistsCalls(stub func(library.BrowseArgs) ([]library.Artist, int)) {
	fake.browseArtistsMutex.Lock()
	defer fake.browseArtistsMutex.Unlock()
	fake.BrowseArtistsStub = stub
}

func (fake *FakeBrowser) BrowseArtistsArgsForCall(i int) library.BrowseArgs {
	fake.browseArtistsMutex.RLock()
	defer fake.browseArtistsMutex.RUnlock()
	argsForCall := fake.browseArtistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBrowser) BrowseArtistsReturns(result1 []library.Artist, result2 int) {
	fake.browseArtistsMutex.Lock()
	defer fake.browseArtistsMutex.Unlock()
	fake.BrowseArtistsStub = nil
	fake.browseArtistsReturns = struct {
		result1 []library.Artist
		result2 int
	}{result1, result2}
}

func (fake *FakeBrowser) BrowseArtistsReturnsOnCall(i int, result1 []library.Artist, result2 int) {
	fake.browseArtistsMutex.Lock()
	defer fake.browseArtistsMutex.Unlock()
	fake.BrowseArtistsStub = nil
	if fake.browseArtistsReturnsOnCall == nil {
		fake.browseArtistsReturnsOnCall = make(map[int]struct {
			result1 []library.Artist
			result2 int
		})
	}
	fake.browseArtistsReturnsOnCall[i] = struct {
		result1 []library.Artist
		result2 int
	}{result1, result2}
}

func (fake *FakeBrowser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.browseAlbumsMutex.RLock()
	defer fake.browseAlbumsMutex.RUnlock()
	fake.browseArtistsMutex.RLock()
	defer fake.browseArtistsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBrowser) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ library.Browser = new(FakeBrowser)