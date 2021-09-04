package company

import (
	"apus-sample/common/database/queries"
	"gorm.io/gorm"
)

func listCompanies(db *gorm.DB, filters map[string]string) (companies []Company, err error) {
	//filterMap := map[string][]string{
	//	"title_contains": {"Con"},
	//	"id_btw":         {"0", "4"},
	//}
	query := queries.Filter(db, Company{}, filters)
	if err != nil {
		panic(err)
	}
	return companies, query.Find(&companies).Error

}

//func listCompanies(conn database.Conn, companies []Company, pageReq vo.PageRequest) {
//
//}
//
//func retrieveCompany(conn database.Conn, comp *Company, id uint64) error {
//	return conn.Retrieve(comp, id)
//}
//
//func createCompany(conn database.Conn, comp *Company) error {
//	return conn.Create(comp)
//}
//
//func updateCompany(conn database.Conn, comp *Company) error {
//	return conn.Update(comp)
//}
//
//func deleteCompany(conn database.Conn, comp *Company) error {
//	return conn.Delete(comp)
//}
