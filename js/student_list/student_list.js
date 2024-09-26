// student_list.js

const tableBody = document.getElementById('table-body');

fetch('http://localhost:8778/api/students/list') // Replace with your API URL
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
                    <td><a href="#" class="details-button" data-url="/details/${item.id}">Details</a></td> <!-- Details button -->
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
