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
	StoreID              *json.Number `json:"storeId" csv:"storeId"`
	InsDateTime          string       `json:"insDateTime" csv:"insDateTime"`
	UpdDateTime          string       `json:"updDateTime" csv:"updDateTime"`
}

type CustomerParquetSchema struct {
	CustomerID           int64   `json:",string" parquet:"name=customer_id, type=INT64"`
	CustomerCode         string  `parquet:"name=customer_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	CustomerNo           *string `parquet:"name=customer_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Rank                 *string `parquet:"name=rank, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	StaffRank            *string `parquet:"name=staff_rank, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	LastName             string  `parquet:"name=last_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	FirstName            string  `parquet:"name=first_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	LastKana             string  `parquet:"name=last_kana, type=BYTE_ARRAY, convertedtype=UTF8"`
	FirstKana            string  `parquet:"name=first_kana, type=BYTE_ARRAY, convertedtype=UTF8"`
	PostCode             *string `parquet:"name=post_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Address              *string `parquet:"name=address, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PhoneNumber          *string `parquet:"name=phone_number, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	FaxNumber            *string `parquet:"name=fax_number, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	MobileNumber         *string `parquet:"name=mobile_number, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	MailAddress          *string `parquet:"name=mail_address, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	MailAddress2         *string `parquet:"name=mail_address2, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	MailAddress3         *string `parquet:"name=mail_address3, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	CompanyName          *string `parquet:"name=company_name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	DepartmentName       *string `parquet:"name=department_name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	ManagerialPosition   *string `parquet:"name=managerial_position, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Sex                  string  `parquet:"name=sex, type=BYTE_ARRAY, convertedtype=UTF8"`
	BirthDate            *string `parquet:"name=birth_date, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Mile                 *int64  `json:",string" parquet:"name=mile, type=INT64"`
	Point                int64   `json:",string" parquet:"name=point, type=INT64"`
	PointExpireDate      *string `parquet:"name=point_expire_date, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	LastComeDateTime     *string `parquet:"name=last_come_date_time, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	EntryDate            string  `parquet:"name=entry_date, type=BYTE_ARRAY, convertedtype=UTF8"`
	LeaveDate            *string `parquet:"name=leave_date, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PointGivingUnitPrice *string `parquet:"name=point_giving_unit_price, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PointGivingUnit      *string `parquet:"name=point_giving_unit, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PinCode              *string `parquet:"name=pin_code, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	PassportNo           *string `parquet:"name=passport_no, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Nationality          *string `parquet:"name=nationality, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	AlphabetName         *string `parquet:"name=alphabet_name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	MailReceiveFlag      string  `parquet:"name=mail_receive_flag, type=BYTE_ARRAY, convertedtype=UTF8"`
	Note                 *string `parquet:"name=note, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Note2                *string `parquet:"name=note2, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	FavoriteList         *string `parquet:"name=favorite_list, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	BrowsingList         *string `parquet:"name=browsing_list, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Status               string  `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8"`
	StoreID              *int64  `json:",string" parquet:"name=store_id, type=INT64, repetitiontype=OPTIONAL"`
	InsDateTime          string  `parquet:"name=ins_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
	UpdDateTime          string  `parquet:"name=upd_date_time, type=BYTE_ARRAY, convertedtype=UTF8"`
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
