package main

import (
	"encoding/json"
)

type Customer struct {
	CustomerID           json.Number  `json:"customerId" csv:"customerId"`
	CustomerCode         string       `json:"customerCode" csv:"customerCode"`
	CustomerNo           *string      `json:"customerNo" csv:"customerNo"`
	Rank                 *string      `json:"rank" csv:"rank"`
	StaffRank            *string      `json:"staffRank" csv:"staffRank"`
	LastName             string       `json:"lastName" csv:"lastName"`
	FirstName            string       `json:"firstName" csv:"firstName"`
	LastKana             string       `json:"lastKana" csv:"lastKana"`
	FirstKana            string       `json:"firstKana" csv:"firstKana"`
	PostCode             *string      `json:"postCode" csv:"postCode"`
	Address              *string      `json:"address" csv:"address"`
	PhoneNumber          *string      `json:"phoneNumber" csv:"phoneNumber"`
	FaxNumber            *string      `json:"faxNumber" csv:"faxNumber"`
	MobileNumber         *string      `json:"mobileNumber" csv:"mobileNumber"`
	MailAddress          *string      `json:"mailAddress" csv:"mailAddress"`
	MailAddress2         *string      `json:"mailAddress2" csv:"mailAddress2"`
	MailAddress3         *string      `json:"mailAddress3" csv:"mailAddress3"`
	CompanyName          *string      `json:"companyName" csv:"companyName"`
	DepartmentName       *string      `json:"departmentName" csv:"departmentName"`
	ManagerialPosition   *string      `json:"managerialPosition" csv:"managerialPosition"`
	Sex                  string       `json:"sex" csv:"sex"`
	BirthDate            *string      `json:"birthDate" csv:"birthDate"`
	Mile                 *json.Number `json:"mile" csv:"mile"`
	Point                json.Number  `json:"point" csv:"point"`
	PointExpireDate      *string      `json:"pointExpireDate" csv:"pointExpireDate"`
	LastComeDateTime     *string      `json:"lastComeDateTime" csv:"lastComeDateTime"`
	EntryDate            string       `json:"entryDate" csv:"entryDate"`
	LeaveDate            *string      `json:"leaveDate" csv:"leaveDate"`
	PointGivingUnitPrice *string      `json:"pointGivingUnitPrice" csv:"pointGivingUnitPrice"`
	PointGivingUnit      *string      `json:"pointGivingUnit" csv:"pointGivingUnit"`
	PinCode              *string      `json:"pinCode" csv:"pinCode"`
	PassportNo           *string      `json:"passportNo" csv:"passportNo"`
	Nationality          *string      `json:"nationality" csv:"nationality"`
	AlphabetName         *string      `json:"alphabetName" csv:"alphabetName"`
	MailReceiveFlag      string       `json:"mailReceiveFlag" csv:"mailReceiveFlag"`
	Note                 *string      `json:"note" csv:"note"`
	Note2                *string      `json:"note2" csv:"note2"`
	FavoriteList         *string      `json:"favoriteList" csv:"favoriteList"`
	BrowsingList         *string      `json:"browsingList" csv:"browsingList"`
	Status               string       `json:"status" csv:"status"`
	StoreID              *json.Number `json:"storeId" csv:"storeID"`
	InsDateTime          string       `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime          string       `json:"updDateTime" csv:"updDateTime"`
}

type CustomerCSV struct {
	*CSVHandler
	buf []Customer
}

func NewCustomerCSV(bufSize int, output string) (*CustomerCSV, error) {
	buf := make([]Customer, bufSize)
	handler, err := NewCSVHandler([]Customer{}, output)
	if err != nil {
		return nil, err
	}

	return &CustomerCSV{
		handler,
		buf,
	}, nil
}

func (cc *CustomerCSV) Write(resp *SrRefResponse) *CSVWriter {
	for i, r := range resp.Result {
		json.Unmarshal([]byte(r.String()), &cc.buf[i])
	}
	cc.CSVWriter.Write(cc.buf[:len(resp.Result)])
	return cc.CSVWriter
}
