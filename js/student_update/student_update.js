// Function to call the API with the student ID
function callApiWithId(id) {
    console.log("ID: ",id)
    const apiUrl = `http://localhost:8778/api/students/detail/${id}`; // Replace with your actual API URL
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            console.log(data); // Handle the API response data
            // You can update the DOM or do something with the data here
            document.getElementById('name').value=data.student_info.name || '';
           
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
    
    document.getElementById('profile-btn').addEventListener('click', saveChanges);

async function saveChanges() {

    const apiUrl = `http://localhost:8778/api/students/update/${studentId}`; // Replace with your actual API URL
    const data = {
        name: document.getElementById('name').value,
        assistance: document.getElementById('assistance').value,
        religion: document.getElementById('Religion').value,
        nrc: document.getElementById('nrc').value,
        contact: document.getElementById('contact').value,
        school: document.getElementById('school').value,
        course: document.getElementById('course').value,
        program_duration: Number(document.getElementById('program_duration').value),
        current_year: document.getElementById('current_year').value,
        semester_term: document.getElementById('semester_term').value,
        total_course_cost: parseFloat(document.getElementById('total_course_cost').value) || 0.0,
        estimated_fees_year_1: parseFloat(document.getElementById('estimated_fees_year_1').value) || 0.0,
        estimated_fees_year_2: parseFloat(document.getElementById('estimated_fees_year_2').value) || 0.0,
        estimated_fees_year_3: parseFloat(document.getElementById('estimated_fees_year_3').value) || 0.0,
        estimated_fees_year_4: parseFloat(document.getElementById('estimated_fees_year_4').value) || 0.0,
        estimated_fees_year_5: parseFloat(document.getElementById('estimated_fees_year_5').value) || 0.0,
        payment_date_sem1_year1: document.getElementById('payment_date_sem1_year1').value,
        payment_amount_sem1_year1: parseFloat(document.getElementById('payment_amount_sem1_year1').value) || 0.0,
        payment_date_sem1_year2: document.getElementById('payment_date_sem1_year2').value,
        payment_amount_sem1_year2: parseFloat(document.getElementById('payment_amount_sem1_year2').value) ||0.0,
        other_fees_payment_date1: document.getElementById('other_fees_payment_date1').value,
        other_fees_details1: document.getElementById('other_fees_details1').value,
        other_fees_amount1: parseFloat(document.getElementById('other_fees_amount1').value),
        other_fees_payment_date2: document.getElementById('other_fees_payment_date2').value,
        other_fees_details2: document.getElementById('other_fees_details2').value,
        other_fees_amount2: parseFloat(document.getElementById('other_fees_amount2').value),
        other_fees_payment_date3: document.getElementById('other_fees_payment_date3').value,
        other_fees_details3: document.getElementById('other_fees_details3').value,
        other_fees_amount3: parseFloat(document.getElementById('other_fees_amount3').value),
        projected_total_fees_curr_year: parseFloat(document.getElementById('projected_total_fees_curr_year').value),
        remaining_tuition_fees_curr_year: parseFloat(document.getElementById('remaining_tuition_fees_curr_year').value),
        tuition_fees_paid_by: document.getElementById('tuition_fees_paid_by').value,
        rent_payment_date1: document.getElementById('rent_payment_date1').value,
        rent_paid_month1: document.getElementById('rent_paid_month1').value,
        rent_amount1: parseFloat(document.getElementById('rent_amount1').value),
        rent_payment_date2: document.getElementById('rent_payment_date2').value,
        rent_paid_month2: document.getElementById('rent_paid_month2').value,
        rent_amount2: parseFloat(document.getElementById('rent_amount2').value),
        rent_payment_date3: document.getElementById('rent_payment_date3').value,
        rent_paid_month3: document.getElementById('rent_paid_month3').value,
        rent_amount3: parseFloat(document.getElementById('rent_amount3').value),
        rent_payment_date4: document.getElementById('rent_payment_date4').value,
        rent_paid_month4: document.getElementById('rent_paid_month4').value,
        rent_amount4: parseFloat(document.getElementById('rent_amount4').value),
        upkeep_payment_date1: document.getElementById('upkeep_payment_date1').value,
        upkeep_paid_months1: document.getElementById('upkeep_paid_months1').value,
        upkeep_amount1: parseFloat(document.getElementById('upkeep_amount1').value),
        upkeep_payment_date2: document.getElementById('upkeep_payment_date2').value,
        upkeep_paid_months2: document.getElementById('upkeep_paid_months2').value,
        upkeep_amount2: parseFloat(document.getElementById('upkeep_amount2').value),
        upkeep_payment_date3: document.getElementById('upkeep_payment_date3').value,
        upkeep_paid_months3: document.getElementById('upkeep_paid_months3').value,
        upkeep_amount3: parseFloat(document.getElementById('upkeep_amount3').value),
        upkeep_payment_date4: document.getElementById('upkeep_payment_date4').value,
        upkeep_paid_months4: document.getElementById('upkeep_paid_months4').value,
        upkeep_amount4: parseFloat(document.getElementById('upkeep_amount4').value),
    };

    try {
        console.log("update api data",JSON.stringify(data))
        const response = await fetch(apiUrl, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const responseData = await response.json();
        alert('Changes have been saved!'); // Show alert after a successful response
        console.log(responseData); // Log the response data if needed

        window.location.href="/v1/student/list";

    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        alert('Error saving changes. Please try again.');
    }
}

});

document.addEventListener('DOMContentLoaded', () => {
    const logoutButton = document.getElementById('logout-button');
    const logoutModal = document.getElementById('logoutModal');
    const confirmLogoutBtn = document.getElementById('confirmLogoutBtn');
    const cancelLogoutBtn = document.getElementById('cancelLogoutBtn');

    logoutButton.addEventListener('click', (event) => {
        event.preventDefault(); // Prevent default behavior
        logoutModal.style.display = 'block'; // Show modal
    });

    cancelLogoutBtn.addEventListener('click', () => {
        logoutModal.style.display = 'none'; // Hide modal
    });

    // Confirm logout and redirect
    confirmLogoutBtn.addEventListener('click', () => {
        window.location.href = `/login?user_id=${userId}`; // Redirect to logout page
    });

    // Close the modal if clicking outside the modal content
    window.addEventListener('click', (event) => {
        if (event.target == logoutModal) {
            logoutModal.style.display = 'none';
        }
    });
});


