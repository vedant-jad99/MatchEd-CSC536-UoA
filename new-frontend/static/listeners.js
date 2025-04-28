export function bindTableListeners() {
    const tableBody = document.getElementById('preferencesTableBody');
  
    // Use event delegation to handle dynamic buttons
    tableBody.addEventListener('click', function(event) {
      const button = event.target.closest('button');
      if (!button) return;  // Ensure it's a button that was clicked
  
      const index = button.dataset.index; // Get the index from the button's data attribute
  
      // Handle button clicks based on button text or class
      if (button.classList.contains('editButton')) {
        editPreference(index);  // Call the edit function
      } else if (button.classList.contains('removeButton')) {
        removePreference(index);  // Call the remove function
      }
    });
  }