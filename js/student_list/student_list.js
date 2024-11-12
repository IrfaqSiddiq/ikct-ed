const tableBody = document.getElementById('table-body');
const searchBox = document.getElementById('search-box');
const religionFilter = document.getElementById('religion-filter');
const schoolFilter = document.getElementById('school-search');
const schoolResults = document.getElementById('school-results');
const assistanceFilter = document.getElementById('assistance-filter');
const paginationContainer = document.getElementById('pagination-container');
const goToPageInput = document.getElementById('go-to-page-input');
const goToPageButton = document.getElementById('go-to-page-button');
let currentPage = 1;
let totalPages = 1;
console.log("jsfile host url", hostURL);

function toggleDropdown() {
    const dropdownContent = document.getElementById('dropdown-content');
    dropdownContent.style.display = dropdownContent.style.display === 'block' ? 'none' : 'block';
  }

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

        schoolFilter.addEventListener("keyup", function () {
            // Clear previous suggestions
            schoolResults.innerHTML = '';
        
            // Get the current input value
            const query = schoolFilter.value.toLowerCase().trim();
        
            // Only make an API call if the input is not empty
            if (query.length === 0) {
                schoolResults.style.display = 'none'; // Hide the suggestion box if input is empty
                return;
            }
        
            // Call the API to fetch the school list dynamically
            fetch(`${hostURL}/api/schools/list?school=${query}&limit=10`)
                .then(response => response.json())
                .then(data => {
                    console.log("School Data:", data); // Debug log
        
                    // Check if response format is as expected
                    if (data.status === "success" && Array.isArray(data.schools)) {
                        const filteredSchools = data.schools.map(schoolObj => schoolObj.school); // Extract school names
        
                        // Populate the <ul> with <li> elements for each matching school
                        filteredSchools.forEach(school => {
                            const listItem = document.createElement('li');
                            listItem.textContent = school;
                            listItem.classList.add("suggestion-item"); // Add CSS class for styling
        
                            // Event to fill the input with the selected school name when an item is clicked
                            listItem.addEventListener("click", () => {
                                schoolFilter.value = school;
                                schoolResults.innerHTML = ''; // Clear suggestions after selection
                                schoolResults.style.display = 'none';
                                fetchStudents();
                            });
        
                            schoolResults.appendChild(listItem);
                        });
        
                        // Show or hide the suggestion container based on whether there are any matches
                        schoolResults.style.display = filteredSchools.length > 0 ? 'block' : 'none';
                    } else {
                        console.error('Unexpected format for schools:', data);
                    }
                })
                .catch(error => console.error('Error fetching schools:', error));
        });


    // // Populate School Dropdown
    // fetch(`${hostURL}/api/schools/list`)
    //     .then(response => response.json())
    //     .then(data => {
    //         console.log("School Data:", data); // Debug log
    //         if (data.status === "success" && Array.isArray(data.schools)) {
    //             data.schools.forEach(schoolObj => {
    //                 const option = document.createElement('option');
    //                 option.value = schoolObj.school;  // Use the 'school' property
    //                 option.textContent = schoolObj.school; // Display the 'school' name
    //                 schoolFilter.appendChild(option);
    //             });
    //         } else {
    //             console.error('Unexpected format for schools:', data);
    //         }
    //     })
    //     .catch(error => console.error('Error fetching schools:', error));

}

// Function to fetch and display all students
// Fetch students for the current page with optional filters
function fetchStudents(page = 1) {
    const searchQuery = searchBox.value.toLowerCase();
    const selectedReligion = religionFilter.value;
    const selectedSchool = schoolFilter.value;
    // Get selected checkbox options and join them as a comma-separated string
    const selectedOptions = Array.from(document.querySelectorAll('.dropdown-content input[type="checkbox"]:checked'))
        .map(checkbox => checkbox.value)
        .join(',');

    // Build query parameter string for filters and page number
    let queryParams = new URLSearchParams();
    if (searchQuery) queryParams.append('search', searchQuery);
    if (selectedReligion) queryParams.append('religion', selectedReligion);
    if (selectedSchool) queryParams.append('school', selectedSchool);
    if (selectedOptions) queryParams.append('assistance', selectedOptions);
    queryParams.append('page', page);

    fetch(`${hostURL}/api/students/list?${queryParams.toString()}`)
        .then(response => response.json())
        .then(data => {
            if (data.status === "success" && Array.isArray(data.students_info)) {
                displayStudents(data.students_info);
                totalPages = data.total_page; // Set total pages from API response
                renderPaginationButtons();
            } else {
                console.error('Unexpected API response format:', data);
            }
        })
        .catch(error => console.error('Error fetching student list:', error));

        document.getElementById('dropdown-content').style.display = 'none';
}

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

    document.querySelectorAll('.details-button').forEach(button => {
        button.addEventListener('click', function (event) {
            event.preventDefault();
            const url = this.getAttribute('data-url');
            window.location.href = url;
        });
    });
}

// Render pagination buttons
function renderPaginationButtons() {
    paginationContainer.innerHTML = ""; // Clear existing buttons
    console.log("Total Pages:", totalPages);
    const MAX_VISIBLE_PAGES = 3; // Maximum visible pages around currentPage
    const ellipsis = document.createElement('span');
    ellipsis.innerText = "...";
    ellipsis.classList.add("ellipsis"); // Add a class for styling the ellipsis if needed

    if (totalPages <= MAX_VISIBLE_PAGES) {
        // Display all pages if they fit within MAX_VISIBLE_PAGES
        for (let i = 1; i <= totalPages; i++) {
            addPaginationButton(i);
        }
    } else {
        // Show first page
        addPaginationButton(1);

        // Add ellipsis before middle pages if currentPage is far from start
        //if (currentPage > 3) paginationContainer.appendChild(ellipsis.cloneNode(true));

        // Show pages around the current page
        let startPage = Math.max(2, currentPage - 1);
        let endPage = Math.min(totalPages - 1, currentPage + 1);

        for (let i = startPage; i <= endPage; i++) {
            addPaginationButton(i);
        }

        // Add ellipsis after middle pages if currentPage is far from end
        if (currentPage < totalPages - 2) paginationContainer.appendChild(ellipsis.cloneNode(true));

        // Show last page
        addPaginationButton(totalPages);
    }
}

function addPaginationButton(pageNumber) {
    const button = document.createElement('button');
    button.innerText = pageNumber;
    button.classList.add("pagination-button");
    if (pageNumber === currentPage) button.classList.add("active");

    button.addEventListener('click', () => {
        currentPage = pageNumber;
        fetchStudents(currentPage); // Fetch students for the selected page
    });

    paginationContainer.appendChild(button);
}
// Filter students based on search and filter inputs
function filterStudents() {
    currentPage = 1;
    fetchStudents(currentPage); // Fetch filtered students
}

goToPageButton.addEventListener('click', () => {
    const pageNumber = parseInt(goToPageInput.value, 10);

    // Validate the input for page number
    if (isNaN(pageNumber) || pageNumber < 1 || pageNumber > totalPages) {
        alert(`Please enter a valid page number between 1 and ${totalPages}`);
        goToPageInput.value = ''; // Clear input
        return;
    }

    // Update current page and fetch students for the selected page
    currentPage = pageNumber;
    fetchStudents(currentPage);
    goToPageInput.value = ''; // Clear input after navigation
});

// Event listeners and initialization
document.addEventListener('DOMContentLoaded', () => {
    populateFilters(); // Populate religion and school filters on page load
    fetchStudents(); // Fetch and display all students on page load

    religionFilter.addEventListener('change', filterStudents); // Filter students on religion change
    schoolFilter.addEventListener('change', filterStudents); // Filter students on school change
    searchBox.addEventListener('input', filterStudents); // Filter students as the user types in search box

    // Existing upload functionality
    const uploadButton = document.getElementById('upload-button');
    const addStudentButton = document.getElementById('add-student-button');
    const fileInput = document.getElementById('file-input');
    const errorModal = document.getElementById('errorModal');
    const errorMessageElement = document.getElementById('errorMessage');
    const closeModalBtn = document.getElementById('closeModalBtn');
     
    addStudentButton.addEventListener('click', function () {
        window.location.href = hostURL+'/v1/student/add';
    });

    uploadButton.addEventListener('click', () => {
        fileInput.click();
    });

    fileInput.addEventListener('change', (event) => {
        const selectedFile = event.target.files[0];
        // if (selectedFile) {
        //     const reader = new FileReader();
        //     reader.onload = (e) => {
        //         const contents = e.target.result;
        //         console.log("File contents:", contents);
        //         // Process the CSV contents here
        //     };
        //     reader.readAsText(selectedFile);
        // }
        if (selectedFile) {
            console.log("File selected:", selectedFile.name);
            
            // Create a FormData object to send the file
            var formData = new FormData();
            formData.append("file", selectedFile); // Use the same field name as expected by the Go API
            
            // Fetch API to send the file to the server
            fetch(`/api/students/add/csv`, { // Assuming studentId is defined and holds the student's ID
                method: 'POST',
                body: formData
            })
            .then(response => response.json())
            .then(data => {
                if (data.error) {  
                    errorMessageElement.textContent = data.error; // General error message
                    errorModal.style.display = 'flex';
                    throw new Error(data.error);
                }
                else {
                    // Redirect to success page
                    window.location.href = '/v1/student/list';
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });        
    }
    closeModalBtn.addEventListener('click', () => {
    errorModal.style.display = 'none'; // Hide the modal
});
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
        fetch('/api/students/logout', {
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

    window.addEventListener('click', (event) => {
        if (event.target === logoutModal) {
            logoutModal.style.display = 'none';
        }
    });
});
