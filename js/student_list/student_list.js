// student_list.js

const tableBody = document.getElementById('table-body');
console.log("jsfile host url",hostURL)

fetch(`${hostURL}/api/students/list`) // Replace with your API URL
    .then(response => response.json())
    .then(data => {
        // Ensure the API response contains the students_info array
        if (data.status === "success" && Array.isArray(data.students_info)) {
            data.students_info.forEach(item => {
                const row = document.createElement('tr');
                
                // Populate table row with data from API
                row.innerHTML = `
                    <td>${item.sno}</td>
                    <td>${item.name}</td>
                    <td>${item.assistance}</td>
                    <td>${item.religion}</td>
                    <td>${item.nrc}</td>
                    <td>${item.contact}</td>
                    <td>${item.string}</td>
                    <td><a href="#" class="details-button" data-url="/v1/student/detail/${item.id}">Details</a></td> <!-- Details button -->
                `;

                tableBody.appendChild(row);
            });

            // Add event listener for Details buttons
            document.querySelectorAll('.details-button').forEach(button => {
                button.addEventListener('click', function(event) {
                    event.preventDefault(); // Prevent the default link behavior
                    const url = this.getAttribute('data-url'); // Get URL from data attribute
                    window.location.href = url; // Change the page URL
                });
            });
        } else {
            console.error('Unexpected API response format:', data);
        }
    })
    .catch(error => console.error('Error fetching data:', error));

    document.addEventListener('DOMContentLoaded', () => {
        const uploadButton = document.getElementById('upload-button');
        const fileInput = document.getElementById('file-input');
        const logoutButton = document.getElementById('logout-button');
        const logoutModal = document.getElementById('logoutModal');
        const confirmLogoutBtn = document.getElementById('confirmLogoutBtn');
        const cancelLogoutBtn = document.getElementById('cancelLogoutBtn');
        
        uploadButton.addEventListener('click', () => {
            console.log("Upload button clicked"); // Debug log
            fileInput.click(); // Trigger the hidden file input
        });
        
        fileInput.addEventListener('change', (event) => {
            console.log("File input changed"); // Debug log
            const file = event.target.files[0];
            if (file) {
                console.log("File selected:", file.name); // Debug log
                const reader = new FileReader();
                reader.onload = (e) => {
                    const contents = e.target.result;
                    console.log("File contents:", contents); // Log the file contents
                    // Process the CSV contents here
                };
                reader.readAsText(file); // Read the file as text
            } else {
                console.log("No file selected"); // Debug log
            }
        });

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