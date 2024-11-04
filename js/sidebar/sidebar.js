document.addEventListener('DOMContentLoaded', function() {

    logoutButton.addEventListener('click', (event) => {
        event.preventDefault(); // Prevent default behavior
        logoutModal.style.display = 'block'; // Show modal
    });
    // Your code to manipulate DOM goes here
    console.log('DOM fully loaded and parsed');
});