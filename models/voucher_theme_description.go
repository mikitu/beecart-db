package models

type VoucherThemeDescription struct {
	VoucherThemeId *VoucherTheme `orm:"column(voucher_theme_id);rel(fk)"`
	LanguageId     *Language     `orm:"column(language_id);rel(fk)"`
	Name           string        `orm:"column(name);size(32)"`
}
