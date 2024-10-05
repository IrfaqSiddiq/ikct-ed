package models

import (
	"database/sql"
	"fmt"
	"ikct-ed/config"
	"ikct-ed/utility"
	"log"
	"os"
	"os/exec"
	"time"
)

type StudentsFinancialInfo struct {
	SerialNumber                    int64   `json:"sno"`
	Id                              int64   `json:"id"`
	Name                            string  `json:"name"`
	Assistance                      string  `json:"assistance"`
	Religion                        string  `json:"religion"`
	NRC                             string  `json:"nrc"`
	Contact                         string  `json:"contact"`
	School                          string  `json:"string"`
	Course                          string  `json:"course"`
	ProgramDuration                 int64   `json:"program_duration"`
	CurrentYear                     int64   `json:"currennt_year"`
	SemesterTerm                    string  `json:"semester_term"`
	TotalCourseCost                 float64 `json:"total_course_cost"`
	EstimatedFeesYear1              float64 `json:"estimated_fees_year_1"`
	EstimatedFeesYear2              float64 `json:"estimated_fees_year_2"`
	EstimatedFeesYear3              float64 `json:"estimated_fees_year_3"`
	EstimatedFeesYear4              float64 `json:"estimated_fees_year_4"`
	EstimatedFeesYear5              float64 `json:"estimated_fees_year_5"`
	PaymentDateSem1Year1            string  `json:"payment_date_sem1_year1"`
	PaymentAmountSem1Year1          float64 `json:"payment_amount_sem1_year1"`
	PaymentDateSem1Year2            string  `json:"payment_date_sem1_year2"`
	PaymentAmountSem1Year2          float64 `json:"payment_amount_sem1_year2"`
	ETC                             string  `json:"etc"`
	OtherFeesPaymentDate1           string  `json:"other_fees_payment_date1"`
	OtherFeesDetails1               string  `json:"other_fees_details1"`
	OtherFeesAmount1                float64 `json:"other_fees_amount1"`
	OtherFeesPaymentDate2           string  `json:"other_fees_payment_date2"`
	OtherFeeDetails2                string  `json:"other_fees_details2"`
	OtherFeesAmount2                float64 `json:"other_fees_amount2"`
	OtherFeesPaymentDate3           string  `json:"other_fees_payment_date3"`
	OtherFeesDetails3               string  `json:"other_fees_details3"`
	OtherFeesAmount3                float64 `json:"other_fees_amount3"`
	ProjectedTotalFeesCurrentYear   float64 `json:"projected_total_fees_current_year"`
	RemainingTuitionFeesCurrentYear float64 `json:"remaining_tuition_fees_current_year"`
	TuitionFeesPaidBy               string  `json:"tuition_fees_paid_by"`
	RentPaymentDate1                string  `json:"rent_payment_date1"`
	RentPaidMonths1                 string  `json:"rent_paid_months1"`
	RentAmount1                     float64 `json:"rent_amount1"`
	RentPaymentDate2                string  `json:"rent_payment_date2"`
	RentPaidMonths2                 string  `json:"rent_paid_months2"`
	RentAmount2                     float64 `json:"rent_amount2"`
	RentPaymentDate3                string  `json:"rent_payment_date3"`
	RentPaidMonths3                 string  `json:"rent_paid_months3"`
	RentAmount3                     float64 `json:"rent_amount3"`
	RentPaymentDate4                string  `json:"rent_payment_date4"`
	RentPaidMonths4                 string  `json:"rent_paid_months4"`
	RentAmount4                     float64 `json:"rent_amount4"`
	UpkeepPaymentDate1              string  `json:"upkeep_payment_date1"`
	UpkeepPaidMonths1               string  `json:"upkeep_paid_months1"`
	UpkeepAmount1                   float64 `json:"upkeep_amount1"`
	UpkeepPaymentDate2              string  `json:"upkeep_payment_date2"`
	UpkeepPaidMonths2               string  `json:"upkeep_paid_months2"`
	UpkeepAmount2                   float64 `json:"upkeep_amount2"`
	UpkeepPaymentDate3              string  `json:"upkeep_payment_date3"`
	UpkeepPaidMonths3               string  `json:"upkeep_paid_months3"`
	UpkeepAmount3                   float64 `json:"upkeep_amount3"`
	UpkeepPaymentDate4              string  `json:"upkeep_payment_date4"`
	UpkeepPaidMonths4               string  `json:"upkeep_paid_months4"`
	UpkeepAmount4                   float64 `json:"upkeep_amount4"`
}

func GetStudentsList(page int64) ([]StudentsFinancialInfo, error) {
	db, err := config.GetDB2()
	if err != nil {
		log.Println("GetStudentsList: Failed while connecting with database with error: ", err)
		return []StudentsFinancialInfo{}, err
	}
	defer db.Close()
	studentInfo := []StudentsFinancialInfo{}
	limit := 10
	offset := (page - 1) * int64(limit)

	sno := (page-1)*10 + 1
	query := ` SELECT 
				count(*) OVER(),
				id, 
				name, 
				assistance, 
				religion, 
				nrc, 
				contact, 
				school,
				course,
    			program_duration,
				current_year,
				semester_term
			FROM
				student_financial_info
			LIMIT $1
			OFFSET $2
			`

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		log.Println("GetStudentsList: Failed while executing the query with error: ", err)
		return []StudentsFinancialInfo{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			totalCount      int64
			id              int64
			name            sql.NullString
			assistance      sql.NullString
			religion        sql.NullString
			nrc             sql.NullString
			contact         sql.NullString
			school          sql.NullString
			course          sql.NullString
			programDuration sql.NullInt64
			currentYear     sql.NullInt64
			semesterTerm    sql.NullString
		)

		err = rows.Scan(&totalCount,
			&id,
			&name,
			&assistance,
			&religion,
			&nrc,
			&contact,
			&school,
			&course,
			&programDuration,
			&currentYear,
			&semesterTerm,
		)
		if err != nil {
			log.Println("GetStudentsList: Failed while scanning the query results with error: ", err)
			continue
		}

		studentInfo = append(studentInfo, StudentsFinancialInfo{
			SerialNumber:    sno,
			Id:              id,
			Name:            utility.SQLNullStringToString(name),
			Assistance:      utility.SQLNullStringToString(assistance),
			Religion:        utility.SQLNullStringToString(religion),
			NRC:             utility.SQLNullStringToString(nrc),
			Contact:         utility.SQLNullStringToString(contact),
			School:          utility.SQLNullStringToString(school),
			Course:          utility.SQLNullStringToString(course),
			ProgramDuration: utility.SQLNullIntToInt(programDuration),
			CurrentYear:     utility.SQLNullIntToInt(currentYear),
			SemesterTerm:    utility.SQLNullStringToString(semesterTerm),
		})
		sno++
	}
	return studentInfo, nil
}

// insertCSVIntoDB runs the \copy command to insert CSV data into PostgreSQL
func InsertCSVIntoDB(filePath string) error {

	query := fmt.Sprintf(`\copy temp_student_financial_info(
		id,
		name,
		assistance,
		religion,
		nrc,
		contact,
		school,
		course,
		program_duration,
		current_year,
		semester_term,
		total_course_cost,
		estimated_fees_year_1,
		estimated_fees_year_2,
		estimated_fees_year_3,
		estimated_fees_year_4,
		estimated_fees_year_5,
		payment_date_sem1_year1,
		payment_amount_sem1_year1,
		payment_date_sem1_year2,
		payment_amount_sem1_year2,
		etc,
		other_fees_payment_date1,
		other_fees_details1,
		other_fees_amount1,
		other_fees_payment_date2,
		other_fees_details2,
		other_fees_amount2,
		other_fees_payment_date3,
		other_fees_details3,
		other_fees_amount3,
		projected_total_fees_current_year,
		remaining_tuition_fees_current_year,
		tuition_fees_paid_by,
		rent_payment_date1,
		rent_paid_months1,
		rent_amount1,
		rent_payment_date2,
		rent_paid_months2,
		rent_amount2,
		rent_payment_date3,
		rent_paid_months3,
		rent_amount3,
		rent_payment_date4,
		rent_paid_months4,
		rent_amount4,
		upkeep_payment_date1,
		upkeep_paid_months1,
		upkeep_amount1,
		upkeep_payment_date2,
		upkeep_paid_months2,
		upkeep_amount2,
		upkeep_payment_date3,
		upkeep_paid_months3,
		upkeep_amount3,
		upkeep_payment_date4,
		upkeep_paid_months4,
		upkeep_amount4
		) FROM '%s' WITH CSV HEADER`, filePath)

	fmt.Println("temp query", query)

	os.Setenv("PGPASSWORD", os.Getenv("DBPASS"))
	// Construct the \copy command
	cmd := exec.Command("psql", "-U", os.Getenv("DBUSER"), "-d", os.Getenv("DBNAME"), "-c",
		query)

	// Run the command and capture any errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("InsertCSVIntoDB: Error running copy command: %s\nOutput: %s\n", err, string(output))
		return err
	}

	fmt.Printf("InsertCSVIntoDB: CSV successfully inserted into database. Output: %s\n", string(output))
	err = InsertUniqueRecord()
	if err != nil {
		log.Println("InsertCSVIntoDB: Error inserting records in FINAl table", err)
		return err
	}
	return nil
}

func InsertUniqueRecord() error {
	db, err := config.GetDB2()
	if err != nil {
		log.Println("InsertUniqueRecord: Failed while connecting with database with error: ", err)
		return err
	}
	defer db.Close()
	query := `INSERT INTO student_financial_info(
	id,
	name,
	assistance,
	religion,
	nrc,
	contact,
	school,
	course,
	program_duration,
	current_year,
	semester_term,
	total_course_cost,
	estimated_fees_year_1,
	estimated_fees_year_2,
	estimated_fees_year_3,
	estimated_fees_year_4,
	estimated_fees_year_5,
	payment_date_sem1_year1,
	payment_amount_sem1_year1,
	payment_date_sem1_year2,
	payment_amount_sem1_year2,
	etc,
	other_fees_payment_date1,
	other_fees_details1,
	other_fees_amount1,
	other_fees_payment_date2,
	other_fees_details2,
	other_fees_amount2,
	other_fees_payment_date3,
	other_fees_details3,
	other_fees_amount3,
	projected_total_fees_current_year,
	remaining_tuition_fees_current_year,
	tuition_fees_paid_by,
	rent_payment_date1,
	rent_paid_months1,
	rent_amount1,
	rent_payment_date2,
	rent_paid_months2,
	rent_amount2,
	rent_payment_date3,
	rent_paid_months3,
	rent_amount3,
	rent_payment_date4,
	rent_paid_months4,
	rent_amount4,
	upkeep_payment_date1,
	upkeep_paid_months1,
	upkeep_amount1,
	upkeep_payment_date2,
	upkeep_paid_months2,
	upkeep_amount2,
	upkeep_payment_date3,
	upkeep_paid_months3,
	upkeep_amount3,
	upkeep_payment_date4,
	upkeep_paid_months4,
	upkeep_amount4
)
SELECT 
	id,
	name,
	assistance,
	religion,
	nrc,
	contact,
	school,
	course,
	program_duration,
	current_year,
	semester_term,
	total_course_cost,
	estimated_fees_year_1,
	estimated_fees_year_2,
	estimated_fees_year_3,
	estimated_fees_year_4,
	estimated_fees_year_5,
	TO_DATE(payment_date_sem1_year1, 'DD/MM/YYYY'),
	payment_amount_sem1_year1,
	TO_DATE(payment_date_sem1_year2, 'DD/MM/YYYY'),
	payment_amount_sem1_year2,
	etc,
	TO_DATE(other_fees_payment_date1, 'DD/MM/YYYY'),
	other_fees_details1,
	other_fees_amount1,
	TO_DATE(other_fees_payment_date2, 'DD/MM/YYYY'),
	other_fees_details2,
	other_fees_amount2,
	TO_DATE(other_fees_payment_date3, 'DD/MM/YYYY'),
	other_fees_details3,
	other_fees_amount3,
	projected_total_fees_current_year,
	remaining_tuition_fees_current_year,
	tuition_fees_paid_by,
	TO_DATE(rent_payment_date1, 'DD/MM/YYYY'),
	rent_paid_months1,
	rent_amount1,
	TO_DATE(rent_payment_date2, 'DD/MM/YYYY'),
	rent_paid_months2,
	rent_amount2,
	TO_DATE(rent_payment_date3, 'DD/MM/YYYY'),
	rent_paid_months3,
	rent_amount3,
	TO_DATE(rent_payment_date4, 'DD/MM/YYYY'),
	rent_paid_months4,
	rent_amount4,
	TO_DATE(upkeep_payment_date1, 'DD/MM/YYYY'),
	upkeep_paid_months1,
	upkeep_amount1,
	TO_DATE(upkeep_payment_date2, 'DD/MM/YYYY'),
	upkeep_paid_months2,
	upkeep_amount2,
	TO_DATE(upkeep_payment_date3, 'DD/MM/YYYY'),
	upkeep_paid_months3,
	upkeep_amount3,
	TO_DATE(upkeep_payment_date4, 'DD/MM/YYYY'),
	upkeep_paid_months4,
	upkeep_amount4
 FROM temp_student_financial_info
ON CONFLICT (nrc) DO NOTHING;
`

	_, err = db.Exec(query)
	if err != nil {
		log.Println("InsertUniqueRecord: Failed while executing the query with error: ", err)
		return err
	}
	return nil
}

func GetStudentDetail(studentID int64) (StudentsFinancialInfo, error) {
	var (
		id                                  int64
		name                                sql.NullString
		assistance                          sql.NullString
		religion                            sql.NullString
		nrc                                 sql.NullString
		contact                             sql.NullString
		school                              sql.NullString
		course                              sql.NullString
		program_duration                    sql.NullInt64
		current_year                        sql.NullInt64
		semester_term                       sql.NullString
		total_course_cost                   sql.NullFloat64
		estimated_fees_year_1               sql.NullFloat64
		estimated_fees_year_2               sql.NullFloat64
		estimated_fees_year_3               sql.NullFloat64
		estimated_fees_year_4               sql.NullFloat64
		estimated_fees_year_5               sql.NullFloat64
		payment_date_sem1_year1             sql.NullTime
		payment_amount_sem1_year1           sql.NullFloat64
		payment_date_sem1_year2             sql.NullTime
		payment_amount_sem1_year2           sql.NullFloat64
		etc                                 sql.NullString
		other_fees_payment_date1            sql.NullTime
		other_fees_details1                 sql.NullString
		other_fees_amount1                  sql.NullFloat64
		other_fees_payment_date2            sql.NullTime
		other_fees_details2                 sql.NullString
		other_fees_amount2                  sql.NullFloat64
		other_fees_payment_date3            sql.NullTime
		other_fees_details3                 sql.NullString
		other_fees_amount3                  sql.NullFloat64
		projected_total_fees_current_year   sql.NullFloat64
		remaining_tuition_fees_current_year sql.NullFloat64
		tuition_fees_paid_by                sql.NullString
		rent_payment_date1                  sql.NullTime
		rent_paid_months1                   sql.NullString
		rent_amount1                        sql.NullFloat64
		rent_payment_date2                  sql.NullTime
		rent_paid_months2                   sql.NullString
		rent_amount2                        sql.NullFloat64
		rent_payment_date3                  sql.NullTime
		rent_paid_months3                   sql.NullString
		rent_amount3                        sql.NullFloat64
		rent_payment_date4                  sql.NullTime
		rent_paid_months4                   sql.NullString
		rent_amount4                        sql.NullFloat64
		upkeep_payment_date1                sql.NullTime
		upkeep_paid_months1                 sql.NullString
		upkeep_amount1                      sql.NullFloat64
		upkeep_payment_date2                sql.NullTime
		upkeep_paid_months2                 sql.NullString
		upkeep_amount2                      sql.NullFloat64
		upkeep_payment_date3                sql.NullTime
		upkeep_paid_months3                 sql.NullString
		upkeep_amount3                      sql.NullFloat64
		upkeep_payment_date4                sql.NullTime
		upkeep_paid_months4                 sql.NullString
		upkeep_amount4                      sql.NullFloat64
	)
	db, err := config.GetDB2()
	if err != nil {
		log.Println("GetStudentsList: Failed while connecting with database with error: ", err)
		return StudentsFinancialInfo{}, err
	}
	defer db.Close()
	query := `SELECT 
				id,
    			name,
    			assistance,
    			religion,
    			nrc,
    			contact,
    			school,
    			course,
    			program_duration,
    			current_year,
    			semester_term,
    			total_course_cost,
    			estimated_fees_year_1,
    			estimated_fees_year_2,
    			estimated_fees_year_3,
    			estimated_fees_year_4,
    			estimated_fees_year_5,
    			payment_date_sem1_year1,
    			payment_amount_sem1_year1,
    			payment_date_sem1_year2,
    			payment_amount_sem1_year2,
    			etc,
    			other_fees_payment_date1,
    			other_fees_details1,
    			other_fees_amount1,
    			other_fees_payment_date2,
    			other_fees_details2,
    			other_fees_amount2,
    			other_fees_payment_date3,
    			other_fees_details3,
    			other_fees_amount3,
    			projected_total_fees_current_year,
    			remaining_tuition_fees_current_year,
    			tuition_fees_paid_by,
    			rent_payment_date1,
    			rent_paid_months1,
    			rent_amount1,
    			rent_payment_date2,
    			rent_paid_months2,
    			rent_amount2,
    			rent_payment_date3,
    			rent_paid_months3,
    			rent_amount3,
    			rent_payment_date4,
    			rent_paid_months4,
    			rent_amount4,
    			upkeep_payment_date1,
    			upkeep_paid_months1,
    			upkeep_amount1,
    			upkeep_payment_date2,
    			upkeep_paid_months2,
    			upkeep_amount2,
    			upkeep_payment_date3,
    			upkeep_paid_months3,
    			upkeep_amount3,
    			upkeep_payment_date4,
    			upkeep_paid_months4,
    			upkeep_amount4
			FROM
				student_financial_info
			WHERE 
				id = $1
			`
	err = db.QueryRow(query, studentID).Scan(
		&id,
		&name,
		&assistance,
		&religion,
		&nrc,
		&contact,
		&school,
		&course,
		&program_duration,
		&current_year,
		&semester_term,
		&total_course_cost,
		&estimated_fees_year_1,
		&estimated_fees_year_2,
		&estimated_fees_year_3,
		&estimated_fees_year_4,
		&estimated_fees_year_5,
		&payment_date_sem1_year1,
		&payment_amount_sem1_year1,
		&payment_date_sem1_year2,
		&payment_amount_sem1_year2,
		&etc,
		&other_fees_payment_date1,
		&other_fees_details1,
		&other_fees_amount1,
		&other_fees_payment_date2,
		&other_fees_details2,
		&other_fees_amount2,
		&other_fees_payment_date3,
		&other_fees_details3,
		&other_fees_amount3,
		&projected_total_fees_current_year,
		&remaining_tuition_fees_current_year,
		&tuition_fees_paid_by,
		&rent_payment_date1,
		&rent_paid_months1,
		&rent_amount1,
		&rent_payment_date2,
		&rent_paid_months2,
		&rent_amount2,
		&rent_payment_date3,
		&rent_paid_months3,
		&rent_amount3,
		&rent_payment_date4,
		&rent_paid_months4,
		&rent_amount4,
		&upkeep_payment_date1,
		&upkeep_paid_months1,
		&upkeep_amount1,
		&upkeep_payment_date2,
		&upkeep_paid_months2,
		&upkeep_amount2,
		&upkeep_payment_date3,
		&upkeep_paid_months3,
		&upkeep_amount3,
		&upkeep_payment_date4,
		&upkeep_paid_months4,
		&upkeep_amount4,
	)
	if err != nil {
		log.Println("GetStudentsList: Failed while executing the query with error: ", err)
		return StudentsFinancialInfo{}, err
	}

	fmt.Println("*******studentDetails*******paymentdatesem1year1", payment_date_sem1_year1.Time.Format("02/01/2006"))
	studentInfo := StudentsFinancialInfo{
		Id:                              id,
		Name:                            name.String,
		Assistance:                      assistance.String,
		Religion:                        religion.String,
		NRC:                             nrc.String,
		Contact:                         contact.String,
		School:                          school.String,
		Course:                          course.String,
		ProgramDuration:                 program_duration.Int64,
		CurrentYear:                     current_year.Int64,
		SemesterTerm:                    semester_term.String,
		TotalCourseCost:                 total_course_cost.Float64,
		EstimatedFeesYear1:              estimated_fees_year_1.Float64,
		EstimatedFeesYear2:              estimated_fees_year_2.Float64,
		EstimatedFeesYear3:              estimated_fees_year_3.Float64,
		EstimatedFeesYear4:              estimated_fees_year_4.Float64,
		EstimatedFeesYear5:              estimated_fees_year_5.Float64,
		PaymentDateSem1Year1:            FormatDateTime(payment_date_sem1_year1),
		PaymentAmountSem1Year1:          payment_amount_sem1_year1.Float64,
		PaymentDateSem1Year2:            FormatDateTime(payment_date_sem1_year2),
		PaymentAmountSem1Year2:          payment_amount_sem1_year2.Float64,
		ETC:                             etc.String,
		OtherFeesPaymentDate1:           FormatDateTime(other_fees_payment_date1),
		OtherFeesDetails1:               other_fees_details1.String,
		OtherFeesAmount1:                other_fees_amount1.Float64,
		OtherFeesPaymentDate2:           FormatDateTime(other_fees_payment_date2),
		OtherFeeDetails2:                other_fees_details2.String,
		OtherFeesAmount2:                other_fees_amount2.Float64,
		OtherFeesPaymentDate3:           FormatDateTime(other_fees_payment_date3),
		OtherFeesDetails3:               other_fees_details3.String,
		OtherFeesAmount3:                other_fees_amount3.Float64,
		ProjectedTotalFeesCurrentYear:   projected_total_fees_current_year.Float64,
		RemainingTuitionFeesCurrentYear: remaining_tuition_fees_current_year.Float64,
		TuitionFeesPaidBy:               tuition_fees_paid_by.String,
		RentPaymentDate1:                FormatDateTime(rent_payment_date1),
		RentPaidMonths1:                 rent_paid_months1.String,
		RentAmount1:                     rent_amount1.Float64,
		RentPaymentDate2:                FormatDateTime(rent_payment_date2),
		RentPaidMonths2:                 rent_paid_months2.String,
		RentAmount2:                     rent_amount2.Float64,
		RentPaymentDate3:                FormatDateTime(rent_payment_date3),
		RentPaidMonths3:                 rent_paid_months3.String,
		RentAmount3:                     rent_amount3.Float64,
		RentPaymentDate4:                FormatDateTime(rent_payment_date4),
		RentPaidMonths4:                 rent_paid_months4.String,
		RentAmount4:                     rent_amount4.Float64,
		UpkeepPaymentDate1:              FormatDateTime(upkeep_payment_date1),
		UpkeepPaidMonths1:               upkeep_paid_months1.String,
		UpkeepAmount1:                   upkeep_amount1.Float64,
		UpkeepPaymentDate2:              FormatDateTime(upkeep_payment_date2),
		UpkeepPaidMonths2:               upkeep_paid_months2.String,
		UpkeepAmount2:                   upkeep_amount2.Float64,
		UpkeepPaymentDate3:              FormatDateTime(upkeep_payment_date3),
		UpkeepPaidMonths3:               upkeep_paid_months3.String,
		UpkeepAmount3:                   upkeep_amount3.Float64,
		UpkeepPaymentDate4:              FormatDateTime(upkeep_payment_date4),
		UpkeepPaidMonths4:               upkeep_paid_months4.String,
		UpkeepAmount4:                   upkeep_amount4.Float64,
	}
	return studentInfo, nil
}

func FormatDateTime(dateTime sql.NullTime) string {
	paymentDate := ""

	if dateTime.Valid && !dateTime.Time.IsZero() && !dateTime.Time.Equal(time.Time{}) {
		paymentDate = dateTime.Time.Format("02/01/2006")
	} else {
		fmt.Println("*****date is not valid")
	}
	return paymentDate
}

func UpdateStudentDetail(studentInfo StudentsFinancialInfo) error {

	db, err := config.GetDB2()
	if err != nil {
		log.Println("GetStudentsList: Failed while connecting with database with error: ", err)
		return err
	}
	defer db.Close()

	query := `UPDATE
				student_financial_info
			SET 
    			name = $2,
    			assistance = $3,
    			religion = $4,
    			nrc = $5,
    			contact = $6,
    			school = $7,
    			course = $8,
    			program_duration = $9,
    			current_year = $10,
    			semester_term = $11,
    			total_course_cost = $12,
    			estimated_fees_year_1 = $13,
    			estimated_fees_year_2 = $14,
    			estimated_fees_year_3 = $15,
    			estimated_fees_year_4 = $16,
    			estimated_fees_year_5 = $17,
    			payment_date_sem1_year1 = CASE WHEN $18 != '' THEN TO_DATE($18, 'DD/MM/YYYY') ELSE NULL END,
    			payment_amount_sem1_year1 = $19,
    			payment_date_sem1_year2 = CASE WHEN $20 != '' THEN TO_DATE($20, 'DD/MM/YYYY') ELSE NULL END,
    			payment_amount_sem1_year2 = $21,
    			etc = $22,
    			other_fees_payment_date1 = CASE WHEN $23 != '' THEN TO_DATE($23, 'DD/MM/YYYY') ELSE NULL END,
    			other_fees_details1 = $24,
    			other_fees_amount1 = $25, 
    			other_fees_payment_date2 = CASE WHEN $26 != '' THEN TO_DATE($26, 'DD/MM/YYYY') ELSE NULL END,
    			other_fees_details2 = $27,
    			other_fees_amount2 = $28,
    			other_fees_payment_date3 = CASE WHEN $29 != '' THEN TO_DATE($29, 'DD/MM/YYYY') ELSE NULL END,
    			other_fees_details3 = $30,
    			other_fees_amount3 = $31,
    			projected_total_fees_current_year = $32,
    			remaining_tuition_fees_current_year = $33,
    			tuition_fees_paid_by = $34,
    			rent_payment_date1 = CASE WHEN $35 != '' THEN TO_DATE($35, 'DD/MM/YYYY') ELSE NULL END,
    			rent_paid_months1 = $36,
    			rent_amount1 = $37,
    			rent_payment_date2 = CASE WHEN $38 != '' THEN TO_DATE($38, 'DD/MM/YYYY') ELSE NULL END,
    			rent_paid_months2 = $39,
    			rent_amount2 = $40,
    			rent_payment_date3 = CASE WHEN $41 != '' THEN TO_DATE($41, 'DD/MM/YYYY') ELSE NULL END,
    			rent_paid_months3 = $42,
    			rent_amount3 = $43,
    			rent_payment_date4 = CASE WHEN $44 != '' THEN TO_DATE($44, 'DD/MM/YYYY') ELSE NULL END,
    			rent_paid_months4 = $45,
    			rent_amount4 = $46,
    			upkeep_payment_date1 = CASE WHEN $47 != '' THEN TO_DATE($47, 'DD/MM/YYYY') ELSE NULL END,
    			upkeep_paid_months1 = $48,
    			upkeep_amount1 = $49,
    			upkeep_payment_date2 = CASE WHEN $50 != '' THEN TO_DATE($50, 'DD/MM/YYYY') ELSE NULL END,
    			upkeep_paid_months2 = $51,
    			upkeep_amount2 = $52,
    			upkeep_payment_date3 = CASE WHEN $53 != '' THEN TO_DATE($53, 'DD/MM/YYYY') ELSE NULL END,
    			upkeep_paid_months3 = $54,
    			upkeep_amount3 = $55,
    			upkeep_payment_date4 = CASE WHEN $56 != '' THEN TO_DATE($56, 'DD/MM/YYYY') ELSE NULL END,
    			upkeep_paid_months4 = $57,
    			upkeep_amount4 = $58
			WHERE 
				id = $1
			`

	_, err = db.Exec(query,
		studentInfo.Id,
		studentInfo.Name,
		studentInfo.Assistance,
		studentInfo.Religion,
		studentInfo.NRC,
		studentInfo.Contact,
		studentInfo.School,
		studentInfo.Course,
		studentInfo.ProgramDuration,
		studentInfo.CurrentYear,
		studentInfo.SemesterTerm,
		studentInfo.TotalCourseCost,
		studentInfo.EstimatedFeesYear1,
		studentInfo.EstimatedFeesYear2,
		studentInfo.EstimatedFeesYear3,
		studentInfo.EstimatedFeesYear4,
		studentInfo.EstimatedFeesYear5,
		studentInfo.PaymentDateSem1Year1,
		studentInfo.PaymentAmountSem1Year1,
		studentInfo.PaymentDateSem1Year2,
		studentInfo.PaymentAmountSem1Year2,
		studentInfo.ETC,
		studentInfo.OtherFeesPaymentDate1,
		studentInfo.OtherFeesDetails1,
		studentInfo.OtherFeesAmount1,
		studentInfo.OtherFeesPaymentDate2,
		studentInfo.OtherFeeDetails2,
		studentInfo.OtherFeesAmount2,
		studentInfo.OtherFeesPaymentDate3,
		studentInfo.OtherFeesDetails3,
		studentInfo.OtherFeesAmount3,
		studentInfo.ProjectedTotalFeesCurrentYear,
		studentInfo.RemainingTuitionFeesCurrentYear,
		studentInfo.TuitionFeesPaidBy,
		studentInfo.RentPaymentDate1,
		studentInfo.RentPaidMonths1,
		studentInfo.RentAmount1,
		studentInfo.RentPaymentDate2,
		studentInfo.RentPaidMonths2,
		studentInfo.RentAmount2,
		studentInfo.RentPaymentDate3,
		studentInfo.RentPaidMonths3,
		studentInfo.RentAmount3,
		studentInfo.RentPaymentDate4,
		studentInfo.RentPaidMonths4,
		studentInfo.RentAmount4,
		studentInfo.UpkeepPaymentDate1,
		studentInfo.UpkeepPaidMonths1,
		studentInfo.UpkeepAmount1,
		studentInfo.UpkeepPaymentDate2,
		studentInfo.UpkeepPaidMonths2,
		studentInfo.UpkeepAmount2,
		studentInfo.UpkeepPaymentDate3,
		studentInfo.UpkeepPaidMonths3,
		studentInfo.UpkeepAmount3,
		studentInfo.UpkeepPaymentDate4,
		studentInfo.UpkeepPaidMonths4,
		studentInfo.UpkeepAmount4,
	)
	if err != nil {
		log.Println("UpdateStudentDetails: failed to update student details with error: ", err)
		return err
	}
	return nil
}

func UploadImageofStudent(imageData []byte, id int64) error {
	db, err := config.GetDB2()
	if err != nil {
		log.Println("UploadImageofStudent: Failed while connecting with database with error: ", err)
		return err
	}
	defer db.Close()

	query := ` UPDATE student_financial_info SET profile_pic = $1 WHERE id=$2`

	_, err = db.Exec(query, imageData, id)
	if err != nil {
		log.Println("UploadImageofStudent: Failed while executing the query with error: ", err)
		return err
	}
	return nil
}
