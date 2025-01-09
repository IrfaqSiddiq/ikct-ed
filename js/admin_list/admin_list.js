// Global variable to store the selected role ID
let selectedRoleId = null;

document.addEventListener('DOMContentLoaded', function () {
  // Fetch roles from API when the page loads
  fetchRoles();

  // Show the popup when the "Add New Admin" button is clicked
  document
    .getElementById('add-admin-button')
    .addEventListener('click', function () {
      var popupOverlay = document.getElementById('popup-overlay');
      popupOverlay.style.display = 'flex'; // Show the popup
    });

  // Close popup when Cancel button is clicked
  document.getElementById('cancel-btn').addEventListener('click', function () {
    var popupOverlay = document.getElementById('popup-overlay');
    popupOverlay.style.display = 'none'; // Hide the popup
  });

  // Submit form
  document
    .getElementById('register-user-form')
    .addEventListener('submit', function (event) {
      event.preventDefault(); // Prevent default form submission

      var formData = new FormData(this);

      // Use the global role ID
      var roleId = parseInt(selectedRoleId, 10); // Convert to integer

      // Prepare data to be sent
      const userData = {
        name: formData.get('name'),
        email: formData.get('email'),
        password: formData.get('password'),
        role_id: roleId, // Send the selected role ID
      };

      // Send data to your API (for example, using fetch)
      saveUser(userData);
    });
});

// Fetch roles from API and populate the dropdown
function fetchRoles() {
  fetch(`${hostURL}/api/roles`) // Replace with your actual API endpoint
    .then((response) => response.json())
    .then((data) => {
      if (data.status === 'ok' && data.roles) {
        const roleSelect = document.getElementById('role');
        // Clear existing options before populating
        roleSelect.innerHTML = '<option value="">Select Role</option>';

        data.roles.forEach((role) => {
          const option = document.createElement('option');
          option.value = role.ID; // Role ID from API
          option.textContent = role.Role; // Role name from API
          roleSelect.appendChild(option);
        });

        // Listen for changes in role selection
        roleSelect.addEventListener('change', function () {
          selectedRoleId = roleSelect.value; // Set the global variable when a role is selected
        });
      } else {
        console.error('Error: Invalid response or no roles found');
      }
    })
    .catch((error) => {
      console.error('Error fetching roles:', error);
    });
}

// Save user data
function saveUser(userData) {
  console.log('body', JSON.stringify(userData));
  fetch(`${hostURL}/api/user/create`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(userData),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log('User saved:', data);
      // Handle success (e.g., close popup, show confirmation, etc.)
      document.getElementById('popup-overlay').style.display = 'none';
    })
    .catch((error) => {
      console.error('Error saving user:', error);
      // Handle error (e.g., show error message)
    });
}
