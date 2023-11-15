package consts

import "github.com/ericchiang/css"

var (
	ByCountryLinks = css.MustParse(`li>ul>li>a[href*="rankings_by_country"]`)
	RegionLinks    = css.MustParse(`span>a.region_link`)

	TitleOption  = css.MustParse(`select[name="title"]>option`)
	ColumnOption = css.MustParse(`select[name="displayColumn"]>option`)

	TableRows    = css.MustParse(`tr[style]`)
	TableRowName = css.MustParse(`td.cityOrCountryInIndicesTable`)
	TableColumns = css.MustParse(`td[style]`)
)
