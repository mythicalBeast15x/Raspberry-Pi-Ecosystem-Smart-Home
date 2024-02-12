// Function to add a new light connection
function addLightConnection() {
    // Retrieve the values from input fields
    var name = document.getElementById('lightName').value;
    var location = document.getElementById('location').value;

    // Check if both fields are filled
    if(name && location) {
        // Find the list in the document
        var list = document.getElementById('lightConnectionsList');
        // Create a new list item
        var entry = document.createElement('li');
        // Set the text of the list item to include the name and location
        entry.textContent = `${name} in ${location}`;
        // Add the new list item to the list
        list.appendChild(entry);

        // Clear the input fields for the next entry
        document.getElementById('lightName').value = '';
        document.getElementById('location').value = '';
    } else {
        // Alert the user if any field is missing
        alert('Please fill in all fields.');
    }
}
