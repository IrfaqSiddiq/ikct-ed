// Call the API when the page loads
document.addEventListener('DOMContentLoaded', function () {
    populateReligion();
    populateSchool();
    document.getElementById('profile-btn').addEventListener('click', saveChanges);
    const logoutButton = document.getElementById('logout-button');
    const logoutModal = document.getElementById('logoutModal');
    const confirmLogoutBtn = document.getElementById('confirmLogoutBtn');
    const cancelLogoutBtn = document.getElementById('cancelLogoutBtn');
    var modal = document.getElementById("myModal");
    var icon = document.getElementById("icon");
    var span = document.getElementsByClassName("close")[0];
    var studentPhoto = document.getElementById("student-photo");
    var viewPhotoSection = document.getElementById("view-photo");
    const defaultPhoto = document.getElementById("default-photo");
    var changePhotoBtn = document.getElementById("changePhotoBtn");
    var uploadPhotoInput = document.getElementById("uploadPhotoInput");

async function saveChanges() {

    const apiUrl = `${hostURL}/api/student/insert`; // Replace with your actual API URL

    const selectedAssistance = Array.from(document.querySelectorAll('.assistance-checkbox:checked'))
        .map(checkbox => checkbox.value)
        .join(',');
    const data = {
        name: document.getElementById('name').value,
        assistance: selectedAssistance, 
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
        console.log("insert api data",JSON.stringify(data))
        const response = await fetch(apiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const responseData = await response.json();
        alert('Added New Student!'); // Show alert after a successful response
        console.log(responseData); // Log the response data if needed

        window.location.href="/v1/student/list";

    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        alert('Error adding new student. Please try again.');
    }
}

    // When the user clicks on the icon, open the modal and show the student image
    icon.onclick = function () {
        // Reset the photo display styles each time modal opens
        studentPhoto.style.display = "none";
        defaultPhoto.style.display = "none";
        
        // Display student's current photo
        modal.style.display = "block";
        var photoUrl = "/api/student/image/";
        studentPhoto.src = photoUrl;

        studentPhoto.onload = function() {
            // Show the photo if it loads successfully
            studentPhoto.style.display = "block";
            defaultPhoto.style.display = "none"; // Hide the default icon
        };

        studentPhoto.onerror = function() {
            // Show default photo if image fails to load
            studentPhoto.style.display = "none";
            defaultPhoto.style.display = "block";
        };
    };

    // Close the modal and reset display styles
    span.onclick = function () {
        closeModal();
    };
    window.onclick = function (event) {
        if (event.target == modal) {
            closeModal();
        }
    };

    // Close and reset the modal function
    function closeModal() {
        modal.style.display = "none";
        studentPhoto.style.display = "none";
        defaultPhoto.style.display = "none";
        studentPhoto.src = ""; // Clear src to force reload next time
    }



    // Trigger file input click when "Change Photo" is clicked
    changePhotoBtn.onclick = function () {
        uploadPhotoInput.click(); // Trigger the file input
    };


        // Handle the file input change event
            uploadPhotoInput.onchange = function(event) {
            var selectedFile = event.target.files[0];
            if (selectedFile) {
                console.log("File selected:", selectedFile.name);
                
                // Create a FormData object to send the file
                var formData = new FormData();
                formData.append("profile_pic", selectedFile); // Use the same field name as expected by the Go API
                
                // Fetch API to send the file to the server
                fetch(`/api/student/upload/img`, { // Assuming studentId is defined and holds the student's ID
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
                    window.location.reload();
                })
                .catch(error => {
                    console.error("Error:", error);
                });
            
        }
    };

    deletePhotoBtn.onclick = function () {
            // Assuming studentId is defined and holds the student's ID
            fetch(`/api/student/delete/img/`, { 
                method: 'DELETE' // Use the DELETE method
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Failed to delete image");
                }
                return response.json();
            })
            .then(data => {
                console.log("Photo deleted successfully:", data);
                window.location.reload();
            })
            .catch(error => {
                console.error("Error:", error);
            });
    };
    
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

    // Populate School Dropdown
    fetch(`${hostURL}/api/school/list`)
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


function toggleDropdown() { 
    document.querySelector('.custom-dropdown').classList.toggle('active');
    updateDropdownText(); // Call the function to update button text when toggling
}

function updateDropdownText() {
    const checkboxes = document.querySelectorAll('.assistance-checkbox');
    const dropdownBtn = document.getElementById('Assistance');
    const selected = Array.from(checkboxes)
        .filter(checkbox => checkbox.checked)
        .map(checkbox => checkbox.value);

    // Update button text based on selected checkboxes
    if (selected.length > 0) {
        dropdownBtn.textContent = selected.join(', ');
    } else {
        dropdownBtn.textContent = 'Select Assistance'; // Default text
    }
}

// Close dropdown if clicked outside
window.addEventListener("click", function(event) {
    if (!event.target.closest('.custom-dropdown')) {
        document.querySelector('.custom-dropdown').classList.remove('active');
    }
});

// Add event listeners to checkboxes to update text when selected
document.querySelectorAll('.assistance-checkbox').forEach(checkbox => {
    checkbox.addEventListener('change', updateDropdownText);
});

function goBack() {

    // Set your dynamic URL here, for example
    const dynamicUrl = '/v1/student/list'; 

    // Redirect to the dynamic URL
    window.location.href = dynamicUrl;
}
