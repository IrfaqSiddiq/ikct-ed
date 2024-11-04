document.addEventListener('DOMContentLoaded', () => {
    // Initial fetch of school list when page loads
    fetchSchools();
    
    // Event listeners for popup buttons
    document.getElementById('add-school-button').addEventListener('click', openPopup);
    document.getElementById('cancelButton').addEventListener('click', closePopup);
    document.getElementById('saveSchoolButton').addEventListener('click', addSchool);
});

async function fetchSchools() {

    fetch(`${hostURL}/api/schools/list`)
        .then(response => response.json())
        .then(data => {
            console.log("Schools: ",data);
            if (data.status === "success" && Array.isArray(data.schools)) {
                //console.log("Qafri: ", data.status);
                displaySchools(data.schools);
            } else {
                console.error('Unexpected API response format:', data);
            }
        })
        .catch(error => console.error('Error fetching schools list:', error));
}


// Function to display the schools in the school list section
function displaySchools(schools) {
    console.log("Schools val", schools);
    const tableBody = document.getElementById('table-body');
    tableBody.innerHTML = ""; // Clear the existing table rows

    schools.forEach(school => {
        const row = document.createElement('tr');
        //console.log("SNo", schools.sno);
        row.innerHTML = `
            <td>${school.sno}</td>
            <td>${school.school}</td>
        `;
        tableBody.appendChild(row);
    });
}

// Function to open the popup
function openPopup() {
    document.getElementById('schoolModal').style.display = 'block';
}

// Function to close the popup
function closePopup() {
    document.getElementById('schoolModal').style.display = 'none';
}

// Function to add a new school
async function addSchool() {
    const school = document.getElementById('school').value;
    if (!school) {
        alert('Please enter a school name');
        return;
    }
    const formData = new FormData();
    formData.append('school', school);  

    try {
        const response = await fetch('/api/schools/add', {  
            method: 'POST',
            body: formData  
        });

        if (!response.ok) {
            throw new Error(`Failed to add school: ${response.statusText}`);
        }

        closePopup();  
        fetchSchools();  
    } catch (error) {
        console.error('Error adding school:', error);
    }
}