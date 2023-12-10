package installer

var classes = []map[string]string{
	{"class": "altacv.cls", "url": "https://raw.githubusercontent.com/liantze/AltaCV/main/altacv.cls"},
	{"class": "extarticle.cls", "url": "http://tug.ctan.org/tex-archive/macros/latex/contrib/extsizes/extarticle.cls"},
}

var packageList = []string{
	"pgf", "fontawesome5", "koma-script", "cmap", "ragged2e", "everysel",
	"tcolorbox", "enumitem", "ifmtarg", "dashrule", "changepage", "multirow",
	"environ", "paracol", "lato", "fontaxes", "accsupp", "tikzfill", "hyperref",
}

func GetPackageList() []string {
	return append([]string{}, packageList...)
}

func GetLatexClass() []map[string]string {
	return classes
}
