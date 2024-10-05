// Function to call the API with the student ID
function callApiWithId(id) {
    console.log("ID: ",id)
    console.log("hostURL: ",hostURL)
    const apiUrl = `${hostURL}/api/students/detail/${id}`; // Replace with your actual API URL
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
    callApiWithId(studentId, hostURL); // Call the API with the student ID

    var modal = document.getElementById("myModal");
    var icon = document.getElementById("icon");
    var span = document.getElementsByClassName("close")[0];
    var studentPhoto = document.getElementById("student-photo");
    var viewPhotoSection = document.getElementById("view-photo");
    var changePhotoBtn = document.getElementById("changePhotoBtn");
    var uploadPhotoInput = document.getElementById("uploadPhotoInput");
    const logoutButton = document.getElementById('logout-button');
    const logoutModal = document.getElementById('logoutModal');
    const confirmLogoutBtn = document.getElementById('confirmLogoutBtn');
    const cancelLogoutBtn = document.getElementById('cancelLogoutBtn');

    // When the user clicks on the icon, open the modal and show the student image
    icon.onclick = function () {
        modal.style.display = "block";

        // Display student's current photo
        var photoUrl = "/api/students/image/" + studentId;
        studentPhoto.src = photoUrl;
        viewPhotoSection.style.display = "block";
    };

    // Close modal when clicking the 'x' or outside the modal
    span.onclick = function () {
        modal.style.display = "none";
        viewPhotoSection.style.display = "none";
    };
    window.onclick = function (event) {
        if (event.target == modal) {
            modal.style.display = "none";
            viewPhotoSection.style.display = "none";
        }
    };

    // Trigger file input click when "Change Photo" is clicked
    changePhotoBtn.onclick = function () {
        uploadPhotoInput.click(); // Trigger the file input
    };

    // Handle file input change and upload the selected image
    uploadPhotoInput.onchange = function (event) {
        var selectedFile = event.target.files[0];
        if (selectedFile) {
            console.log("File selected:", selectedFile.name);

            // Create a FormData object to send the file
            var formData = new FormData();
            formData.append("profile_pic", selectedFile); // Adjust according to your API

        // Handle the file input change event
            uploadPhotoInput.onchange = function(event) {
            var selectedFile = event.target.files[0];
            if (selectedFile) {
                console.log("File selected:", selectedFile.name);
                
                // Create a FormData object to send the file
                var formData = new FormData();
                formData.append("profile_pic", selectedFile); // Use the same field name as expected by the Go API
                
                // Fetch API to send the file to the server
                fetch(`/api/students/upload/img/${studentId}`, { // Assuming studentId is defined and holds the student's ID
                    method: 'POST',
                    body: formData
                })
                .then(response => {
                    if (!response.ok) {
                        throw new Error("Failed to upload image");
                    }
                    return response.json();
                })
                .then(data => {
                    console.log("Success:", data);
                    // Optionally refresh the student's photo in the modal
                    studentPhoto.src = "/api/students/image/" + studentId + "?" + new Date().getTime(); // Force browser to load updated image
                })
                .catch(error => {
                    console.error("Error:", error);
                });
            
        }
    };
}}

    
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


        
        
    document.getElementById('profile-btn').addEventListener('click', saveChanges);

    async function saveChanges() {
    window.location.href=`/v1/student/update/${studentId}`
    }
});
