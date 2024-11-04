document.addEventListener('DOMContentLoaded', () => {
    // Initial fetch of school list when page loads
    fetchSchools();
    
    // Event listeners for popup buttons
    document.getElementById('addSchoolBtn').addEventListener('click', openPopup);
    document.getElementById('cancelBtn').addEventListener('click', closePopup);
    document.getElementById('saveBtn').addEventListener('click', addSchool);
});

async function fetchSchools() {

    fetch(`${hostURL}/api/schools/list`)
        .then(response => response.json())
        .then(data => {
            console.log("Schools: ",data);
            if (data.status === "success" && Array.isArray(data.schools_info)) {
                displaySchools(data.schools_info);
            } else {
                console.error('Unexpected API response format:', data);
            }
        })
        .catch(error => console.error('Error fetching schools list:', error));
}


// Function to display the schools in the school list section
function displayStudents(schools) {
    tableBody.innerHTML = ""; // Clear the existing table rows

    schools.forEach(student => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${schools.school}</td>
        `;
        tableBody.appendChild(row);
    });
}

// Function to open the popup
function openPopup() {
    document.getElementById('popup').style.display = 'block';
}

// Function to close the popup
function closePopup() {
    document.getElementById('popup').style.display = 'none';
}

// Function to add a new school
async function addSchool() {
    const schoolName = document.getElementById('schoolName').value;
    if (!schoolName) {
        alert('Please enter a school name');
        return;
    }
    try {
        await fetch('/api/schools/add', {  // Replace with actual API endpoint
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name: schoolName })
        });
        closePopup();
        fetchSchools(); // Reload the school list after adding a new school
    } catch (error) {
        console.error('Error adding school:', error);
    }
}