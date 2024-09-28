// Function to call the API with the student ID
function callApiWithId(id) {
    console.log("ID: ",id)
    const apiUrl = `http://localhost:8778/api/students/detail/${id}`; // Replace with your actual API URL
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            console.log(data); // Handle the API response data
            // You can update the DOM or do something with the data here
            document.getElementById('name').innerText = `${data.student_info.name}`;
            document.getElementById('assistance').value=data.student_info.assistance || '';
            document.getElementById('Religion').value=data.student_info.religion || '';
            document.getElementById('nrc').value=data.student_info.nrc || '';
            document.getElementById('contact').value=data.student_info.contact || '';
            document.getElementById('school').value=data.student_info.school || '';
            document.getElementById('course').value=data.student_info.course || '';
            document.getElementById('program_duration').value=data.student_info.program_duration || '';
            document.getElementById('current_year').value=data.student_info.current_year || '';
            document.getElementById('semester_term').value=data.student_info.semester_term || '';
            document.getElementById('total_course_cost').value=data.student_info.total_course_cost || '';
            document.getElementById('estimated_fees_year_1').value=data.student_info.estimated_fees_year_1 || '';
            document.getElementById('estimated_fees_year_2').value=data.student_info.estimated_fees_year_2 || '';
            document.getElementById('estimated_fees_year_3').value=data.student_info.estimated_fees_year_3 || '';
            document.getElementById('estimated_fees_year_4').value=data.student_info.estimated_fees_year_4 || '';
            document.getElementById('estimated_fees_year_5').value=data.student_info.estimated_fees_year_5 || '';
            document.getElementById('payment_date_sem1_year1').value=data.student_info.payment_date_sem1_year1 || '';
            document.getElementById('payment_amount_sem1_year1').value=data.student_info.payment_amount_sem1_year1 || '';
            document.getElementById('payment_date_sem1_year2').value=data.student_info.payment_date_sem1_year2 || '';
            document.getElementById('payment_amount_sem1_year2').value=data.student_info.payment_amount_sem1_year2 || '';
            document.getElementById('other_fees_payment_date1').value=data.student_info.other_fees_payment_date1 || '';
            document.getElementById('other_fees_details1').value=data.student_info.other_fees_details1 || '';
            document.getElementById('other_fees_amount1').value=data.student_info.other_fees_amount1 || '';
            document.getElementById('other_fees_payment_date2').value=data.student_info.other_fees_payment_date2 || '';
            document.getElementById('other_fees_details2').value=data.student_info.other_fees_details2 || '';
            document.getElementById('other_fees_amount2').value=data.student_info.other_fees_amount2 || '';
            document.getElementById('other_fees_payment_date3').value=data.student_info.other_fees_payment_date3 || '';
            document.getElementById('other_fees_details3').value=data.student_info.other_fees_details3 || '';
            document.getElementById('other_fees_amount3').value=data.student_info.other_fees_amount3 || '';
            document.getElementById('projected_total_fees_curr_year').value=data.student_info.projected_total_fees_curr_year || '';
            document.getElementById('remaining_tuition_fees_curr_year').value=data.student_info.remaining_tuition_fees_curr_year || '';
            document.getElementById('tuition_fees_paid_by').value=data.student_info.tuition_fees_paid_by || '';
            document.getElementById('rent_payment_date1').value=data.student_info.rent_payment_date1 || '';
            document.getElementById('rent_paid_month1').value=data.student_info.rent_paid_month1 || '';
            document.getElementById('rent_amount1').value=data.student_info.rent_amount1 || '';
            document.getElementById('rent_payment_date2').value=data.student_info.rent_payment_date2 || '';
            document.getElementById('rent_paid_month2').value=data.student_info.rent_paid_month2 || '';
            document.getElementById('rent_amount2').value=data.student_info.rent_amount2 || '';
            document.getElementById('rent_payment_date3').value=data.student_info.rent_payment_date3 || '';
            document.getElementById('rent_paid_month3').value=data.student_info.rent_paid_month3 || '';
            document.getElementById('rent_amount3').value=data.student_info.rent_amount3 || '';
            document.getElementById('rent_payment_date4').value=data.student_info.rent_payment_date4 || '';
            document.getElementById('rent_paid_month4').value=data.student_info.rent_paid_month4 || '';
            document.getElementById('rent_amount4').value=data.student_info.rent_amount4 || '';
            document.getElementById('upkeep_payment_date1').value=data.student_info.upkeep_payment_date1 || '';
            document.getElementById('upkeep_paid_months1').value=data.student_info.upkeep_paid_months1 || '';
            document.getElementById('upkeep_amount1').value=data.student_info.upkeep_amount1 || '';
            document.getElementById('upkeep_payment_date2').value=data.student_info.upkeep_payment_date2 || '';
            document.getElementById('upkeep_paid_months2').value=data.student_info.upkeep_paid_months2 || '';
            document.getElementById('upkeep_amount2').value=data.student_info.upkeep_amount2 || '';
            document.getElementById('upkeep_payment_date3').value=data.student_info.upkeep_payment_date3 || '';
            document.getElementById('upkeep_paid_months3').value=data.student_info.upkeep_paid_months3 || '';
            document.getElementById('upkeep_amount3').value=data.student_info.upkeep_amount3 || '';
            document.getElementById('upkeep_payment_date4').value=data.student_info.upkeep_payment_date4 || '';
            document.getElementById('upkeep_paid_months4').value=data.student_info.upkeep_paid_months4 || '';
            document.getElementById('upkeep_amount4').value=data.student_info.upkeep_amount4 || '';
        })
        .catch(error => console.error('Error:', error));
}

// Call the API when the page loads
document.addEventListener('DOMContentLoaded', function () {
    callApiWithId(studentId); // Call the API with the student ID
});