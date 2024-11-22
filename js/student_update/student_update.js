// Function to call the API with the student ID
function callApiWithId(id, hostURL) {
    console.log("ID: ",id)
    const apiUrl = `${hostURL}/api/student/detail/${id}`; // Replace with your actual API URL
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
            document.getElementById('payment_date_sem1_year1').value = formatDateToYMD(data.student_info.payment_date_sem1_year1) || '';
            document.getElementById('payment_amount_sem1_year1').value=data.student_info.payment_amount_sem1_year1 || '';
            document.getElementById('payment_date_sem1_year2').value=formatDateToYMD(data.student_info.payment_date_sem1_year2) || '';
            document.getElementById('payment_amount_sem1_year2').value=data.student_info.payment_amount_sem1_year2 || '';
            document.getElementById('other_fees_payment_date1').value=formatDateToYMD(data.student_info.other_fees_payment_date1) || '';
            document.getElementById('other_fees_details1').value=data.student_info.other_fees_details1 || '';
            document.getElementById('other_fees_amount1').value=data.student_info.other_fees_amount1 || '';
            document.getElementById('other_fees_payment_date2').value=formatDateToYMD(data.student_info.other_fees_payment_date2) || '';
            document.getElementById('other_fees_details2').value=data.student_info.other_fees_details2 || '';
            document.getElementById('other_fees_amount2').value=data.student_info.other_fees_amount2 || '';
            document.getElementById('other_fees_payment_date3').value=formatDateToYMD(data.student_info.other_fees_payment_date3) || '';
            document.getElementById('other_fees_details3').value=data.student_info.other_fees_details3 || '';
            document.getElementById('other_fees_amount3').value=data.student_info.other_fees_amount3 || '';
            document.getElementById('projected_total_fees_curr_year').value=data.student_info.projected_total_fees_current_year || '';
            document.getElementById('remaining_tuition_fees_curr_year').value=data.student_info.remaining_tuition_fees_current_year || '';
            document.getElementById('tuition_fees_paid_by').value=data.student_info.tuition_fees_paid_by || '';
            document.getElementById('rent_payment_date1').value=formatDateToYMD(data.student_info.rent_payment_date1) || '';
            document.getElementById('rent_paid_month1').value=data.student_info.rent_paid_months1 || '';
            document.getElementById('rent_amount1').value=data.student_info.rent_amount1 || '';
            document.getElementById('rent_payment_date2').value=formatDateToYMD(data.student_info.rent_payment_date2) || '';
            document.getElementById('rent_paid_month2').value=data.student_info.rent_paid_months2 || '';
            document.getElementById('rent_amount2').value=data.student_info.rent_amount2 || '';
            document.getElementById('rent_payment_date3').value=formatDateToYMD(data.student_info.rent_payment_date3) || '';
            document.getElementById('rent_paid_month3').value=data.student_info.rent_paid_months3 || '';
            document.getElementById('rent_amount3').value=data.student_info.rent_amount3 || '';
            document.getElementById('rent_payment_date4').value=formatDateToYMD(data.student_info.rent_payment_date4) || '';
            document.getElementById('rent_paid_month4').value=data.student_info.rent_paid_months4 || '';
            document.getElementById('rent_amount4').value=data.student_info.rent_amount4 || '';
            document.getElementById('upkeep_payment_date1').value=formatDateToYMD(data.student_info.upkeep_payment_date1) || '';
            document.getElementById('upkeep_paid_months1').value=data.student_info.upkeep_paid_months1 || '';
            document.getElementById('upkeep_amount1').value=data.student_info.upkeep_amount1 || '';
            document.getElementById('upkeep_payment_date2').value=formatDateToYMD(data.student_info.upkeep_payment_date2) || '';
            document.getElementById('upkeep_paid_months2').value=data.student_info.upkeep_paid_months2 || '';
            document.getElementById('upkeep_amount2').value=data.student_info.upkeep_amount2 || '';
            document.getElementById('upkeep_payment_date3').value=formatDateToYMD(data.student_info.upkeep_payment_date3) || '';
            document.getElementById('upkeep_paid_months3').value=data.student_info.upkeep_paid_months3 || '';
            document.getElementById('upkeep_amount3').value=data.student_info.upkeep_amount3 || '';
            document.getElementById('upkeep_payment_date4').value=formatDateToYMD(data.student_info.upkeep_payment_date4) || '';
            document.getElementById('upkeep_paid_months4').value=data.student_info.upkeep_paid_months4 || '';
            document.getElementById('upkeep_amount4').value=data.student_info.upkeep_amount4 || '';
        })
        .catch(error => console.error('Error:', error));
}



// Call the API when the page loads
document.addEventListener('DOMContentLoaded', function () {
    callApiWithId(studentId,hostURL); // Call the API with the student ID
    populateReligion();
    populateSchool();
    document.getElementById('profile-btn').addEventListener('click', saveChanges);
    document.getElementById('discard-btn').addEventListener('click', discardChanges);
    const logoutButton = document.getElementById('logout-button');
    const logoutModal = document.getElementById('logoutModal');
    const confirmLogoutBtn = document.getElementById('confirmLogoutBtn');
    const cancelLogoutBtn = document.getElementById('cancelLogoutBtn');

async function saveChanges() {

    const apiUrl = `${hostURL}/api/student/update/${studentId}`; // Replace with your actual API URL
    const data = {
        name: document.getElementById('name').value,
        assistance: document.getElementById('assistance').value,
        religion: document.getElementById('Religion').value,
        nrc: document.getElementById('nrc').value,
        contact: document.getElementById('contact').value,
        school: document.getElementById('school').value,
        course: document.getElementById('course').value,
        program_duration: Number(document.getElementById('program_duration').value),
        current_year: Number(document.getElementById('current_year').value),
        semester_term: document.getElementById('semester_term').value,
        total_course_cost: parseFloat(document.getElementById('total_course_cost').value) || 0.0,
        estimated_fees_year_1: parseFloat(document.getElementById('estimated_fees_year_1').value) || 0.0,
        estimated_fees_year_2: parseFloat(document.getElementById('estimated_fees_year_2').value) || 0.0,
        estimated_fees_year_3: parseFloat(document.getElementById('estimated_fees_year_3').value) || 0.0,
        estimated_fees_year_4: parseFloat(document.getElementById('estimated_fees_year_4').value) || 0.0,
        estimated_fees_year_5: parseFloat(document.getElementById('estimated_fees_year_5').value) || 0.0,
        payment_date_sem1_year1: formatDate(document.getElementById('payment_date_sem1_year1').value),
        payment_amount_sem1_year1: parseFloat(document.getElementById('payment_amount_sem1_year1').value) || 0.0,
        payment_date_sem1_year2: formatDate(document.getElementById('payment_date_sem1_year2').value),
        payment_amount_sem1_year2: parseFloat(document.getElementById('payment_amount_sem1_year2').value) ||0.0,
        other_fees_payment_date1: formatDate(document.getElementById('other_fees_payment_date1').value) ,
        other_fees_details1: document.getElementById('other_fees_details1').value,
        other_fees_amount1: parseFloat(document.getElementById('other_fees_amount1').value),
        other_fees_payment_date2: formatDate(document.getElementById('other_fees_payment_date2').value) || '',
        other_fees_details2: document.getElementById('other_fees_details2').value,
        other_fees_amount2: parseFloat(document.getElementById('other_fees_amount2').value),
        other_fees_payment_date3: formatDate(document.getElementById('other_fees_payment_date3').value),
        other_fees_details3: document.getElementById('other_fees_details3').value,
        other_fees_amount3: parseFloat(document.getElementById('other_fees_amount3').value),
        projected_total_fees_current_year: parseFloat(document.getElementById('projected_total_fees_curr_year').value),
        remaining_tuition_fees_current_year: parseFloat(document.getElementById('remaining_tuition_fees_curr_year').value),
        tuition_fees_paid_by: document.getElementById('tuition_fees_paid_by').value,
        rent_payment_date1: formatDate(document.getElementById('rent_payment_date1').value),
        rent_paid_months1: document.getElementById('rent_paid_month1').value,
        rent_amount1: parseFloat(document.getElementById('rent_amount1').value),
        rent_payment_date2: formatDate(document.getElementById('rent_payment_date2').value),
        rent_paid_months2: document.getElementById('rent_paid_month2').value,
        rent_amount2: parseFloat(document.getElementById('rent_amount2').value),
        rent_payment_date3: formatDate(document.getElementById('rent_payment_date3').value) || '',
        rent_paid_months3: document.getElementById('rent_paid_month3').value,
        rent_amount3: parseFloat(document.getElementById('rent_amount3').value),
        rent_payment_date4: formatDate(document.getElementById('rent_payment_date4').value),
        rent_paid_months4: document.getElementById('rent_paid_month4').value,
        rent_amount4: parseFloat(document.getElementById('rent_amount4').value),
        upkeep_payment_date1: formatDate(document.getElementById('upkeep_payment_date1').value),
        upkeep_paid_months1: document.getElementById('upkeep_paid_months1').value,
        upkeep_amount1: parseFloat(document.getElementById('upkeep_amount1').value),
        upkeep_payment_date2: formatDate(document.getElementById('upkeep_payment_date2').value),
        upkeep_paid_months2: document.getElementById('upkeep_paid_months2').value,
        upkeep_amount2: parseFloat(document.getElementById('upkeep_amount2').value),
        upkeep_payment_date3: formatDate(document.getElementById('upkeep_payment_date3').value),
        upkeep_paid_months3: document.getElementById('upkeep_paid_months3').value,
        upkeep_amount3: parseFloat(document.getElementById('upkeep_amount3').value),
        upkeep_payment_date4: formatDate(document.getElementById('upkeep_payment_date4').value),
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

function discardChanges() {
    // Logic to discard changes, such as reloading the page
    if (confirm('Are you sure you want to discard changes?')) {
        location.reload();
    }
}
    
    logoutButton.addEventListener('click', (event) => {
        event.preventDefault(); // Prevent default behavior
        logoutModal.style.display = 'block'; // Show modal
    });

    cancelLogoutBtn.addEventListener('click', () => {
        logoutModal.style.display = 'none'; // Hide modal
    });

    // Confirm logout and redirect
    confirmLogoutBtn.addEventListener('click', () => {
        fetch('/api/student/logout', {
            method: 'POST', // Adjust the method if needed (e.g., 'GET')
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Something went wrong!!!');
            }
            console.log(response.json()); // If the API returns JSON, handle it
        })
        .then(data => {
            // If successful, reload the page
            console.log('Logout successful:', data);
            window.location.reload(); // Reloads the current page
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
    });

    // Close the modal if clicking outside the modal content
    window.addEventListener('click', (event) => {
        if (event.target == logoutModal) {
            logoutModal.style.display = 'none';
        }
    });
});
 

function formatDate(dateString) {
    if(dateString != '') {
        const date = new Date(dateString);
        const day = String(date.getDate()).padStart(2, '0'); // Pad single digit days
        const month = String(date.getMonth() + 1).padStart(2, '0'); // Pad single digit months
        const year = date.getFullYear();
        const formattedDate = `${day}/${month}/${year}`;   
        console.log('Formatted Date:', formattedDate); // Log the formatted date to the console
        return formattedDate;
    }
}

function formatDateToYMD(dateString) {
    const parts = dateString.split('/');
    if (parts.length === 3) {
        // Rearrange date parts to "yyyy-MM-dd"
        return `${parts[2]}-${parts[1].padStart(2, '0')}-${parts[0].padStart(2, '0')}`;
    }
    return dateString; // Return the original string if it's not in "dd/MM/yyyy" format
}

function populateReligion() {
    const religionFilter = document.getElementById('Religion'); // Get the select element
    
    // Populate Religion Dropdown
    fetch(`${hostURL}/api/religion`)
        .then(response => response.json())
        .then(data => {
            console.log("Religion Data:", data); // Debug log
            if (data.status === "success" && Array.isArray(data.religion)) {
                data.religion.forEach(religionObj => {
                    const option = document.createElement('option');
                    option.value = religionObj.religion;  // Use the 'religion' property
                    option.textContent = religionObj.religion; // Display the 'religion' name
                    religionFilter.appendChild(option);
                });
            } else {
                console.error('Unexpected format for religions:', data);
            }
        })
        .catch(error => console.error('Error fetching religions:', error));
}

function populateSchool() {
    const schoolFilter = document.getElementById('school'); // Get the select element

    // Populate Religion Dropdown
    fetch(`${hostURL}/api/schools/list`)
            .then(response => response.json())
            .then(data => {
                console.log("School Data:", data); // Debug log
                if (data.status === "success" && Array.isArray(data.schools)) {
                    data.schools.forEach(schoolObj => {
                        const option = document.createElement('option');
                        option.value = schoolObj.school;  // Use the 'school' property
                        option.textContent = schoolObj.school; // Display the 'school' name
                        schoolFilter.appendChild(option);
                    });
                } else {
                    console.error('Unexpected format for schools:', data);
                }
            })
            .catch(error => console.error('Error fetching schools:', error));
}

function goBack() {

    // Set your dynamic URL here, for example
    const dynamicUrl = '/v1/student/detail/'+studentId; 

    // Redirect to the dynamic URL
    window.location.href = dynamicUrl;
}
