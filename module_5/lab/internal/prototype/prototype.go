package prototype

type IPrototype interface {
	Clone() IPrototype
}

type Image struct {
	URL string
	Alt string
}

func (i Image) Clone() IPrototype {
	return Image{URL: i.URL, Alt: i.Alt}
}

type Section struct {
	Title   string
	Content string
	Images  []Image
}

func (s Section) Clone() IPrototype {
	newImages := make([]Image, len(s.Images))
	for idx, img := range s.Images {
		newImages[idx] = img.Clone().(Image)
	}
	return Section{Title: s.Title, Content: s.Content, Images: newImages}
}

type Document struct {
	Title    string
	Content  string
	Sections []Section
}

func (d Document) Clone() IPrototype {
	newSections := make([]Section, len(d.Sections))
	for i, sec := range d.Sections {
		newSections[i] = sec.Clone().(Section)
	}
	return Document{Title: d.Title, Content: d.Content, Sections: newSections}
}

type DocumentManager struct{}

func (DocumentManager) CreateDocument(proto IPrototype) IPrototype {
	return proto.Clone()
}
