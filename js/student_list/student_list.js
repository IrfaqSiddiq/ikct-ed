// student_list.js

const tableBody = document.getElementById('table-body');
const searchBox = document.getElementById('search-box');
const religionFilter = document.getElementById('religion-filter');
const schoolFilter = document.getElementById('school-filter');
console.log("jsfile host url", hostURL);

// Function to populate filters with data from API
function populateFilters() {
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

    // Populate School Dropdown
    fetch(`${hostURL}/api/schools`)
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

// Function to fetch and display all students
function fetchStudents() {
    fetch(`${hostURL}/api/students/list`)
        .then(response => response.json())
        .then(data => {
            if (data.status === "success" && Array.isArray(data.students_info)) {
                displayStudents(data.students_info); // Initial display of all students
            } else {
                console.error('Unexpected API response format:', data);
            }
        })
        .catch(error => console.error('Error fetching student list:', error));
}

// Function to display students in the table
function displayStudents(students) {
    tableBody.innerHTML = ""; // Clear the existing table rows

    students.forEach(student => {
        const row = document.createElement('tr');

        row.innerHTML = `
            <td>${student.sno}</td>
            <td>${student.name}</td>
            <td>${student.assistance}</td>
            <td>${student.religion}</td>
            <td>${student.nrc}</td>
            <td>${student.contact}</td>
            <td>${student.school}</td>
            <td><a href="#" class="details-button" data-url="/v1/student/detail/${student.id}">Details</a></td>
        `;

        tableBody.appendChild(row);
    });

    // Add event listeners for Details buttons
    document.querySelectorAll('.details-button').forEach(button => {
        button.addEventListener('click', function (event) {
            event.preventDefault();
            const url = this.getAttribute('data-url');
            window.location.href = url;
        });
    });
}

// Function to filter students based on input and dropdowns
function filterStudents() {
    const searchQuery = searchBox.value.toLowerCase();
    const selectedReligion = religionFilter.value;
    const selectedSchool = schoolFilter.value;

    fetch(`${hostURL}/api/students/list`)
        .then(response => response.json())
        .then(data => {
            if (data.status === "success" && Array.isArray(data.students_info)) {
                const filteredStudents = data.students_info.filter(student => {
                    const matchesSearch = student.name.toLowerCase().includes(searchQuery);
                    const matchesReligion = selectedReligion ? student.religion === selectedReligion : true;
                    const matchesSchool = selectedSchool ? student.school === selectedSchool : true;
                    return matchesSearch && matchesReligion && matchesSchool;
                });

                displayStudents(filteredStudents); // Update table with filtered students
            }
        })
        .catch(error => console.error('Error filtering students:', error));
}

// Event listeners and initialization
document.addEventListener('DOMContentLoaded', () => {
    populateFilters(); // Populate religion and school filters on page load
    fetchStudents(); // Fetch and display all students on page load

    religionFilter.addEventListener('change', filterStudents); // Filter students on religion change
    schoolFilter.addEventListener('change', filterStudents); // Filter students on school change
    searchBox.addEventListener('input', filterStudents); // Filter students as the user types in search box

    // Existing upload functionality
    const uploadButton = document.getElementById('upload-button');
    const fileInput = document.getElementById('file-input');

    uploadButton.addEventListener('click', () => {
        fileInput.click();
    });

    fileInput.addEventListener('change', (event) => {
        const file = event.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = (e) => {
                const contents = e.target.result;
                console.log("File contents:", contents);
                // Process the CSV contents here
            };
            reader.readAsText(file);
        }
    });

    // Existing logout functionality
    const logoutButton = document.getElementById('logout-button');
    const logoutModal = document.getElementById('logoutModal');
    const confirmLogoutBtn = document.getElementById('confirmLogoutBtn');
    const cancelLogoutBtn = document.getElementById('cancelLogoutBtn');

    logoutButton.addEventListener('click', (event) => {
        event.preventDefault();
        logoutModal.style.display = 'block';
    });

    cancelLogoutBtn.addEventListener('click', () => {
        logoutModal.style.display = 'none';
    });

    confirmLogoutBtn.addEventListener('click', () => {
        fetch('/api/students/logout', { method: 'POST' })
            .then(response => response.json())
            .then(data => {
                if (data.status === 'success') {
                    window.location.reload();
                }
            })
            .catch(error => console.error('Error logging out:', error));
    });

    window.addEventListener('click', (event) => {
        if (event.target === logoutModal) {
            logoutModal.style.display = 'none';
        }
    });
});
