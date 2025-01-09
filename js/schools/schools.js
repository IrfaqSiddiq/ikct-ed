const searchBox = document.getElementById("search-box");

document.addEventListener("DOMContentLoaded", () => {
  let currentPage = 1; // Set initial page to 1

  // Initial fetch of school list when page loads
  fetchSchools(currentPage);

  // Event listeners for popup buttons
  document
    .getElementById("add-school-button")
    .addEventListener("click", openPopup);
  document.getElementById("cancelButton").addEventListener("click", closePopup);
  document
    .getElementById("saveSchoolButton")
    .addEventListener("click", addSchool);
  const logoutButton = document.getElementById("logout-button");
  const logoutModal = document.getElementById("logoutModal");

  console.log("Search", searchBox);
  const confirmLogoutBtn = document.getElementById("confirmLogoutBtn");
  const cancelLogoutBtn = document.getElementById("cancelLogoutBtn");
  searchBox.addEventListener("input", () => fetchSchools(currentPage));

  logoutButton.addEventListener("click", (event) => {
    event.preventDefault(); // Prevent default behavior
    logoutModal.style.display = "block"; // Show modal
  });

  cancelLogoutBtn.addEventListener("click", () => {
    logoutModal.style.display = "none"; // Hide modal
  });

  // Confirm logout and redirect
  confirmLogoutBtn.addEventListener("click", () => {
    fetch("/api/student/logout", {
      method: "POST", // Adjust the method if needed (e.g., 'GET')
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Something went wrong!!!");
        }
        console.log(response.json()); // If the API returns JSON, handle it
      })
      .then((data) => {
        // If successful, reload the page
        console.log("Logout successful:", data);
        window.location.reload(); // Reloads the current page
      })
      .catch((error) => {
        console.error("There was a problem with the fetch operation:", error);
      });
  });

  // Close the modal if clicking outside the modal content
  window.addEventListener("click", (event) => {
    if (event.target == logoutModal) {
      logoutModal.style.display = "none";
    }
  });
  // Listen for pagination button clicks
  document
    .getElementById("pagination-container")
    .addEventListener("click", (event) => {
      if (event.target && event.target.matches(".pagination-btn")) {
        currentPage = parseInt(event.target.getAttribute("data-page"));
        fetchSchools(currentPage);
      }
    });
});

// Function to fetch schools list with pagination support
async function fetchSchools(page = 1) {
  const searchQuery = searchBox.value.toLowerCase();
  const limit = 10; // Adjust the number of schools per page if needed
  let queryParams = new URLSearchParams();
  if (searchQuery) queryParams.append("school", searchQuery);
  queryParams.append("page", Number(page));

  const url = `${hostURL}/api/school/list?${queryParams.toString()}&limit=10`;

  try {
    const response = await fetch(url);
    const data = await response.json();

    if (data.status === "success" && Array.isArray(data.schools)) {
      // Logic to check the condition and disable the button

      if (data.permissions.Role2permission.create === false) {
        // Replace with your actual condition
        const button = document.getElementById("add-school-button");

        button.disabled = true;
        // Apply custom CSS styles
        button.style.backgroundColor = "#ccc"; // Example: gray background
        button.style.cursor = "not-allowed"; // Example: disabled cursor
        button.style.opacity = "0.6"; // Example: reduced opacity
      }

      displaySchools(data.schools);
      createPagination(data.total_page, page);
    } else {
      console.error("Unexpected API response format:", data);
    }
  } catch (error) {
    console.error("Error fetching schools list:", error);
  }
}

// Function to display the schools in the school list section
function displaySchools(schools) {
  const tableBody = document.getElementById("table-body");
  tableBody.innerHTML = ""; // Clear the existing table rows

  schools.forEach((school) => {
    const row = document.createElement("tr");
    row.innerHTML = `
            <td>${school.sno}</td>
            <td>${school.school}</td>
            <td>
                <button class="delete-btn" onclick="deleteSchool(${school.id})">
                    🗑️
                </button>
            </td>
        `;
    tableBody.appendChild(row);
  });
}

// Function to handle delete action
function deleteSchool(id) {
  const confirmation = confirm("Are you sure you want to delete this school?");
  if (confirmation) {
    // Call the API to delete the school
    fetch(`${hostURL}/api/school/delete/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => {
        if (response.ok) {
          alert("School deleted successfully!");
          // Re-fetch the schools to update the table
          fetchSchools(); // Pass the current page if necessary
        } else {
          alert("Failed to delete the school. Try again.");
        }
      })
      .catch((error) => {
        console.error("Error deleting school:", error);
        alert("An error occurred. Please try again.");
      });
  }
}

// Function to create pagination buttons
function createPagination(totalPages = 1, currentPage) {
  console.log(totalPages, currentPage);
  const paginationContainer = document.getElementById("pagination-container");
  paginationContainer.innerHTML = ""; // Clear previous pagination
  console.log("Total Pages: ", totalPages);
  for (let i = 1; i <= totalPages; i++) {
    const button = document.createElement("button");
    button.textContent = i;
    button.classList.add("pagination-btn");
    button.setAttribute("data-page", i);
    if (i === currentPage) {
      button.classList.add("active"); // Highlight the current page
    }
    paginationContainer.appendChild(button);
  }
}

// Function to open the popup
function openPopup() {
  document.getElementById("schoolModal").style.display = "block";
}

// Function to close the popup
function closePopup() {
  document.getElementById("schoolModal").style.display = "none";
}

// Function to add a new school
async function addSchool() {
  const school = document.getElementById("school").value;
  if (!school) {
    alert("Please enter a school name");
    return;
  }
  const formData = new FormData();
  formData.append("school", school);

  try {
    const response = await fetch("/api/school/add", {
      method: "POST",
      body: formData,
    });

    if (!response.ok) {
      throw new Error(`Failed to add school: ${response.statusText}`);
    }

    closePopup();
    fetchSchools(1);
  } catch (error) {
    console.error("Error adding school:", error);
  }
}
