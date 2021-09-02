package company

import (
	"apus-sample/common/database"
	"apus-sample/common/vo"
)

func listAllCompanies(conn database.Conn) (companies []Company, err error) {
	return companies, conn.List(&companies)

}

func listCompanies(conn database.Conn, companies []Company, pageReq vo.PageRequest) {

}

func retrieveCompany(conn database.Conn, comp *Company, id uint64) error {
	return conn.Retrieve(comp, id)
}

func createCompany(conn database.Conn, comp *Company) error {
	return conn.Create(comp)
}

func updateCompany(conn database.Conn, comp *Company) error {
	return conn.Update(comp)
}

func deleteCompany(conn database.Conn, comp *Company) error {
	return conn.Delete(comp)
}
