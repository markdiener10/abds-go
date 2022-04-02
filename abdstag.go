package abds

import "strings"

type AbdsTags struct {
	vals []string
}

func (g *AbdsTags) tags(ntag uint) string {
	//We are a slot 1 based system, slot 0 is not a tag
	if ntag == 0 {
		return ""
	}
	if uint(len(g.vals)) <= ntag {
		return ""
	}
	return g.vals[ntag]
}

func (g *AbdsTags) tagi(stag string) uint {
	//We are a slot 1 based system, slot 0 is not a tag
	for idx := 1; idx < len(g.vals); idx++ {
		if g.vals[idx] != stag {
			continue
		}
		return uint(idx)
	}
	return 0
}

func (g *AbdsTags) tag(ntag uint) string {
	if g == nil {
		return ""
	}
	if ntag == 0 {
		return ""
	}
	if uint(len(g.vals)) < ntag {
		return ""
	}
	return g.vals[ntag]
}

func (g *AbdsTags) add(stag string) uint {
	stag = strings.TrimSpace(stag)
	if stag == "" {
		return 0
	}
	g.vals = append(g.vals, stag)

	//We already created the vals array with len==1, slot 0 is empty
	//So our first addition will be len(g.vals) == 2, but really slot 1
	return uint(len(g.vals) - 1)
}

func NewTags() *AbdsTags {
	pTags := &AbdsTags{vals: make([]string, 1)}
	pTags.vals[0] = ""
	return pTags
}

///ABDS Tag operations
func (g *Abds) TagExists(stag string) bool {

	if g == nil {
		return false
	}

	ntag := g.GetTagi(stag)
	if ntag == 0 {
		return false
	}
	return true
}

func (g *Abds) GetTag(ntag uint) string {

	if g == nil {
		return ""
	}
	if ntag == 0 {
		return ""
	}
	var pAbds *Abds = nil

	if !g.root {
		if g.link == nil {
			g.adderr("ABDS GetTag link is null but should have value")
			return ""
		}
		pAbds = g.link.(*Abds)
	} else {
		pAbds = g
	}
	if pAbds.tags == nil {
		//No tags yet? Means they have not added any tags yet to this abds
		//Tags are an on-demand feature (totally optional)
		//Tags can be injected
		return ""
	}
	return pAbds.tags.tags(ntag)
}

func (g *Abds) GetTagi(stag string) uint {

	if g == nil {
		return 0
	}
	var pAbds *Abds = nil
	stag = strings.TrimSpace(stag)
	if stag == "" {
		return 0
	}
	if !g.root {
		if g.link == nil {
			g.adderr("ABDS GetTag link is null but should have value")
			return 0
		}
		pAbds = g.link.(*Abds)
	} else {
		pAbds = g
	}
	if pAbds.tags == nil {
		return 0
	}
	return pAbds.tags.tagi(stag)
}

func (g *Abds) GetSetTag(stag string) uint {

	if g == nil {
		return 0
	}

	var pAbds *Abds = nil
	stag = strings.TrimSpace(stag)
	if stag == "" {
		return 0
	}
	if !g.root {
		if g.link == nil {
			g.adderr("ABDS GetTag link is null but should have value")
			return 0
		}
		pAbds = g.link.(*Abds)
	} else {
		pAbds = g
	}
	if pAbds.tags == nil {
		pAbds.tags = NewTags()
	}
	ntag := pAbds.tags.tagi(stag)
	if ntag != 0 {
		return ntag
	}
	return pAbds.tags.add(stag)
}

/*

for idx := 0; idx < len(g.vals); idx++ {
	pItem = g.vals[idx]
	if pItem.tag != ntag {
		continue
	}
	return true
}

return true
}


tag = strings.TrimSpace(tag)
if tag == "" {
return false
}
pItem := g.find(tag, parmbool(parms...))
if pItem == nil {
return false
}
return true


func (g *Abds) tags(tag string) AbdsTag {

	if g.root {
		return g.tags.Tagi(tag)
	}
	if g.link == nil {
		g.adderr("ABDS non-root called with null link")
		return 0
	}
	return (g.link.(*Abds)).tags(tag)
}

func (g *Abds) tagi(AbdsTag) string {

	ntag := g.tags.Tagi(tag)
	if ntag == 0 {



}

*/
