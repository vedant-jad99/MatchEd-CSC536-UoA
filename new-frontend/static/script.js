import { fetchAllCourses, fetchAllPreferences, upsertPreference,
         deleteUser, fetchAllCoursesFormatted, fetchAllPreferencesFormatted, 
         fetchAllUsers, fetchLatestMatchings, fetchLatestMatchingsFormatted, 
         upsertUser, fetchUserById, fetchPreferenceById, triggerMatchingEngine,
          upsertCourseAndSemester, removeCourseSemester, 
          fetchCourseSemester,
          updateCourseSemester} from "./api.js";

let editRow = null;
let editMode = ''; 
let facultyChanges = [];
let courseChanges = [];
let preferenceChanges = [];
let editPreferenceRow = null;

function updateButtonIndecies(tableBody){
  const rows = tableBody.querySelectorAll('tr');
  rows.forEach((row, i) => {
    const button = row.querySelector('button');
    if (button) {
      button.dataset.index = i; // Reassign data-index based on the current row's position
    }
  });
}

function editFaculty(id) {
  const row = document.querySelector(`#facultyTableBody tr[data-id="${id}"]`);
  editRow = row;
  if (row) {
    editMode = 'faculty';
    document.getElementById("editfacultyName").value = row.children[0].textContent;
    openFacultySidebar();
  } else {
    console.log('Row with the specified ID not found');
  }
}

// Remove Faculty by ID
function removeFaculty(id) {
  const row = document.querySelector(`#facultyTableBody tr[data-id="${id}"]`); // Find the row with the matching data-id
  
  if (row) {
    deleteUser(id);
    const name = row.children[0].textContent;  // Get the name from the first column
    row.remove();  // Remove the row from the table
    facultyChanges.push({ type: "remove", name: name });
    updateFacultyPushState();
  } else {
    console.log('Row with the specified ID not found');
  }
}

// Add Faculty
async function addFaculty() {
  const name = document.getElementById('facultyName').value.trim();
  if (!name) return;

  const userData = {
    id: 0,
    name: name,
  };
  console.log(userData)

  const user = await upsertUser(userData)
  .then(response => console.log('User data upserted:', response))
  .catch(error => console.error('Error upserting user data:', error));

  // todo, use user.id to lookup facultyTable row
  // instead of refreshing
  loadFaculty();
  document.getElementById('facultyName').value = '';  // Clear input field
}

// Edit Course Entry
function editCourse(id) {
  const row = document.querySelector(`#coursesTableBody tr[data-id="${id}"]`);
  editRow = row;

  if (row) {
    editMode = 'course';
    document.getElementById("editName").value = editRow.children[0].textContent;
    const section = editRow.children[1].textContent.trim();
    document.getElementById("editSection").value = section;
    openSidebar();
  } else {
    console.log('Row with the specified ID not found');
  }
  
  openSidebar();
}

function removeCourse(id) {
  const row = document.querySelector(`#coursesTableBody tr[data-id="${id}"]`); // Find the row with the matching data-id
  
  if (row) {
    removeCourseSemester(id);
    const name = row.children[0].textContent;  // Get the name from the first column
    courseChanges.push({ type: "remove", name: name });
    updateCoursePushState();
    row.remove();
  } else {
    console.log('Row with the specified ID not found');
  }
}

// add course
async function addCourse() {
  console.log("add course button clicked");
  const name = document.getElementById('courseName').value.trim();
  console.log(name)
  const section = document.getElementById('sectionCount').value.trim();

  if (!name) return;

  const courseData = {
    id: 0,
    name: name,
    number: name,
  };
  console.log(courseData)

  const courseSemester = await upsertCourseAndSemester(courseData)
  .then(response => console.log('Course data upserted:', response))
  .catch(error => console.error('Error upserting course data:', error));

  console.log(courseSemester)
  // todo, id to lookup row
  // instead of refreshing
  loadCourses();

  document.getElementById('facultyName').value = '';  // Clear input field
  document.getElementById('courseName').value = '';  // Clear input field
  document.getElementById('sectionCount').value = '1';  // Reset section count
}


// Tab navigation functionality
document.querySelectorAll('.tab-button').forEach(button => {
  button.addEventListener('click', () => {
    document.querySelectorAll('.tab-button').forEach(btn => btn.classList.remove('active'));
    document.querySelectorAll('.tab-content').forEach(tab => tab.classList.remove('active'));

    button.classList.add('active');
    document.getElementById(button.dataset.tab).classList.add('active');
  });
});

// Open Sidebar
function openSidebar() {
    document.getElementById("editSidebar").classList.add("active");
    document.body.classList.add("sidebar-open");
  }

function openFacultySidebar() {
  document.getElementById("editFacultySidebar").classList.add("active");
  document.body.classList.add("sidebar-open");
}
  
function closeSidebar() {
    document.getElementById("editSidebar").classList.remove("active");
    document.body.classList.remove("sidebar-open");
  }

function closeFacultySidebar() {
  document.getElementById("editFacultySidebar").classList.remove("active");
  document.body.classList.remove("sidebar-open");
  document.getElementById("editfacultyName").value = '';
  document.getElementById("facultyName").value = '';
}
  
// Push buttons
function updateFacultyPushState() {
  const btn = document.getElementById("pushFacultyChangesBtn");
  btn.disabled = facultyChanges.length === 0;
  btn.classList.toggle("active", facultyChanges.length > 0);
}
  
function updateCoursePushState() {
  const btn = document.getElementById("pushCourseChangesBtn");
  btn.disabled = courseChanges.length === 0;
  btn.classList.toggle("active", courseChanges.length > 0);
}

function updatePreferencePushState() {
  const btn = document.getElementById("pushChangesBtn");
  btn.disabled = preferenceChanges.length === 0;
  btn.classList.toggle("active", preferenceChanges.length > 0);
}
  
document.addEventListener('click', function(event) {
    const sidebar = document.getElementById('editSidebar');
    if (
      sidebar.classList.contains('active') &&
      !sidebar.contains(event.target) &&
      !event.target.closest('button') // Prevents closing when clicking buttons
    ) {
      closeSidebar();
    }
});
  
  
// Handle courses Sidebar Form Submission
document.getElementById("editForm").addEventListener("submit", async function (e) {
  e.preventDefault();
  const name = document.getElementById("editName").value.trim();

  if (!name) return;
  if (editMode === 'faculty') {
    if (newName !== oldName) {

      facultyChanges.push({ type: "edit", name: oldName, newName: newName });
      updatePushState();
    }
  } else if (editMode === 'course') {
    const section = document.getElementById("editSection").value;

    const id = editRow.dataset.id;

    const courseSemester = await fetchCourseSemester(id);

    console.log(courseSemester);
  
    if (!courseSemester) {
      console.warn(`No courseSemester found with id ${id}`);
      return; // or handle fallback logic here
    }
    courseSemester.section = section;
    await updateCourseSemester(courseSemester);

    loadCourses();

    closeSidebar();
   
    const oldName = editRow.children[0];
    const sectionCell = editRow.children[1];
    const oldSection = sectionCell.textContent.trim();
    sectionCell.textContent = section;
    courseChanges.push({ type: "edit", name: oldName, newName: name });
    updateCoursePushState();
  }
});

// Open sidebar to edit preference
function editPreference(id) {
  const row = document.querySelector(`#preferencesTableBody tr[data-id="${id}"]`);
  if (!row){
    console.log('Row with the specified ID not found');
    return;
  }

  editRow = row;
  editMode = 'preferences';

  // TODO you can get preferences from the db instead
  // Extract the values from the row
  const course = editRow.children[0].textContent;  // First column
  const faculty = editRow.children[1].textContent;  // Second column
  const preference = editRow.children[2].textContent;  // Third column
  const weight = editRow.children[4].textContent;  // Fifth column

  console.log(weight)
  
  document.getElementById('editCourse').value = course;
  document.getElementById('editFaculty').value = faculty;
  document.getElementById('editPreference').value = preference;
  document.getElementById('editWeight').value = weight;
  document.getElementById('editWeightValue').textContent = weight;
  
  // todo weight
  console.log(weight);

  openPreferenceSidebar();
}

async function loadFaculty(){
  const response = await fetchAllUsers();
  const tableBody = document.getElementById('facultyTableBody');
  tableBody.innerHTML = '';  // Clear any existing rows

  response.forEach((user) => {
    const row = document.createElement('tr');
    row.setAttribute('data-id', user.id);

    row.innerHTML = `
      <td>${user.name}</td>
      <td>
        <button class="editButton" data-id="${user.id}">Edit</button>
        <button class="removeButton" data-id="${user.id}">Remove</button>
      </td>
    `;
    tableBody.appendChild(row);
  });
}

async function loadCourses(){
  const response = await fetchAllCoursesFormatted();
  const tableBody = document.getElementById('coursesTableBody');
  tableBody.innerHTML = '';  // Clear any existing rows

  response.forEach((course) => {
    const row = document.createElement('tr');
    row.setAttribute('data-id', course.id);

    row.innerHTML = `
      <td>${course.name}</td>
      <td>${course.sections}</td>
      <td>
        <button class="editButton" data-id="${course.id}">Edit</button>
        <button class="removeButton" data-id="${course.id}">Remove</button>
      </td>
    `;
    tableBody.appendChild(row);
  });
}

// loads preference table from the server
async function loadPreferences() {
  const preferencesData = await fetchAllPreferencesFormatted();

  const tableBody = document.getElementById('preferencesTableBody');
  tableBody.innerHTML = '';  // Clear any existing rows

  preferencesData.forEach((pref, index) => {
    const row = document.createElement('tr');
    row.setAttribute('data-id', pref.id);

    const isEdited = editedRows.some(r => r.rowIndex === index);
    
    row.innerHTML = `
      <td>${pref.course}</td>
      <td>${pref.faculty}</td>
      <td><span class="preference" style="background-color: ${pref.preference};">${pref.preference}</span></td>
      <td>
        <button class="editButton" data-id="${pref.id}">Edit</button>
        <button class="removeButton" data-id="${pref.id}" style="display: none;">Remove</button>
      </td>
      <td>${pref.weight}</td>
    `;
    if (isEdited) {
      row.style.backgroundColor = '#f0f8ff'; 
    }
    tableBody.appendChild(row);
  });
}

async function loadOutput(){
  const response = await fetchLatestMatchingsFormatted();
  const tableBody = document.getElementById('outputsTableBody');
  tableBody.innerHTML = '';  // Clear any existing rows

  response.forEach((matching, index) => {
    const row = document.createElement('tr');

    row.innerHTML = `
      <td>${matching.course}</td>
      <td>${matching.faculty}</td>
      <td><span class="score ${matching.score} data-id="${matching.id}">${matching.score}</span></td>
    `;
    tableBody.appendChild(row);
  });
}

// Bind event listeners for the faculty table
function bindFacultyTableListeners() {
  const facultyTableBody = document.getElementById('facultyTableBody');

  // Event delegation for edit and remove buttons in the faculty table
  facultyTableBody.addEventListener('click', function(event) {
    const button = event.target.closest('button');
    if (!button) return;

    const id = button.dataset.id;
    console.log("clicked: " + id)

    if (button.classList.contains('editButton')) {
      editFaculty(id);  // Edit faculty
    } else if (button.classList.contains('removeButton')) {
      removeFaculty(id);  // Remove faculty
      updateButtonIndecies(facultyTableBody);
    }
  });
}

// Bind event listeners for the course table
function bindCoursesTableListeners() {
  const coursesTableBody = document.getElementById('coursesTableBody');

  // Event delegation for edit and remove buttons in the course table
  coursesTableBody.addEventListener('click', function(event) {
    console.log(event.target)
    const button = event.target.closest('button');
    if (!button) return;

    const id = button.dataset.id;
    console.log("button id");
    console.log(id);

    if (button.classList.contains('editButton')) {
      editCourse(id);  // Edit course
    } else if (button.classList.contains('removeButton')) {
      removeCourse(id);  // Remove course
    }
  });
}

// Bind event listeners for the preferences table
function bindPreferencesTableListeners() {
  const preferencesTableBody = document.getElementById('preferencesTableBody');

  // Event delegation for edit and remove buttons in the preferences table
  preferencesTableBody.addEventListener('click', function(event) {
    const button = event.target.closest('button');
    if (!button) return;

    const id = button.dataset.id;

    if (button.classList.contains('editButton')) {
      editPreference(id);  // Edit preference
    } else if (button.classList.contains('removeButton')) {
      removePreference(id);  // Remove preference
    }
  });
}

// bind match button to matching engine trigger and reload
document.querySelector('#matchBtn').addEventListener('click', async () => {
  //await triggerMatchingEngine();
  loadOutput();
});

// below used to be inline
// Bind the addFaculty and addCourse buttons from index.html
document.addEventListener('DOMContentLoaded', function() {
  loadFaculty();
  loadCourses();
  loadPreferences();
  loadOutput();
  // Bind 'Add Faculty' button
  const addFacultyButton = document.querySelector('#faculty button');
  console.log(addFacultyButton);
  addFacultyButton.addEventListener('click', addFaculty);

  // Bind 'Add Course' button
  const addCourseButton = document.querySelector('#courses button');
  addCourseButton.addEventListener('click', addCourse);

  // Bind preference sidebar cancel button
  const closePreferenceButton = document.getElementById('closePreferenceButton');
  closePreferenceButton.addEventListener('click', closePreferenceSidebar)

 // Bind faculty sidebar cancel button
 const closeFacultyButton = document.getElementById('closeFacultyButton');
 closeFacultyButton.addEventListener('click', closeFacultySidebar)

 // bind general sidebar function 
 const closeSidebarButton = document.getElementById('closeSidebarButton');
 closeSidebarButton.addEventListener('click', closeSidebar)


  // Bind tables after loading content
  bindFacultyTableListeners();
  bindCoursesTableListeners();
  bindPreferencesTableListeners();
});


// Open sidebar
function openPreferenceSidebar() {
  document.getElementById('editPreferenceSidebar').classList.add('active');
  document.body.classList.add('sidebar-open');

}

// Close sidebar
function closePreferenceSidebar() {
  document.getElementById('editPreferenceSidebar').classList.remove('active');
  document.body.classList.remove('sidebar-open');
}

document.getElementById('editWeight').addEventListener('input', function () {
  document.getElementById('editWeightValue').textContent = this.value;
});

// Save edited preference
document.getElementById('editPreferenceForm').addEventListener('submit', async function (e) {
  e.preventDefault();

  let course, faculty, preference, weight;
  
  course = document.getElementById('editCourse').value;
  faculty = document.getElementById('editFaculty').value;
  preference = document.getElementById('editPreference').value;
  weight = parseInt(document.getElementById('editWeight').value, 10);

  const id = editRow.dataset.id;
  const pref = await fetchPreferenceById(id);
  if (!pref) {
    console.warn(`No preference found with id ${id}`);
    return; // or handle fallback logic here
  }

  console.log(pref)
  // Modify the correct snake_case fields
  pref.preference_level = preference;
  pref.preference_weight = weight; // TODO: replace with slider value later

  // update database
  let newPref = await upsertPreference(pref);
  if (!newPref) {
    console.warn(`failed to update preference`);
    return; // or handle fallback logic here
  }

  // update push button
  updatePreferencePushState();

  // add preference update to changelog
  preferenceChanges.push({ type: "edit", courseName: course, facultyName: faculty, newColor: preference});

  loadPreferences();
  closePreferenceSidebar();
});

document.getElementById("editFacultyForm").addEventListener("submit", async function (e) {
  e.preventDefault(); // Prevents page reload
  
  const newName = document.getElementById("editfacultyName").value.trim();
  if (!newName || !editRow) return;

  const id = editRow.dataset.id;
  const user = await fetchUserById(id);
  if (!user) {
    console.warn(`No user found with id ${id}`);
    return; // or handle fallback logic here
  }

  const oldName = editRow.children[0].textContent.trim();
  editRow.children[0].textContent = newName;

  if (newName !== oldName) {
    user.name = newName;
    upsertUser(user);

    facultyChanges.push({ type: "edit", name: oldName, newName: newName });
    updateFacultyPushState();
  }
  
  closeFacultySidebar();
});


// Remove preference
function removePreference(index) {
  preferencesData.splice(index, 1);
  loadPreferences();
}

let editedRows = [];


function trackEdit(rowIndex, column, oldValue, newValue) {
  const existingIndex = editedRows.findIndex(
    row => row.rowIndex === rowIndex && row.column === column
  );

  if (existingIndex !== -1) {
    if (oldValue === newValue) {
      editedRows.splice(existingIndex, 1);
    } else {
      editedRows[existingIndex].newValue = newValue;
    }
  } else if (oldValue !== newValue) {
    editedRows.push({ rowIndex, column, oldValue, newValue });
  }
  console.log("Edited Rows:", JSON.stringify(editedRows, null, 2));
  updatePushButtonState();
}

function updatePushButtonState() {
  const btn = document.getElementById("pushChangesBtn");
  if (editedRows.length > 0) {
    btn.disabled = false;
    btn.classList.add("active");
  } else {
    btn.disabled = true;
    btn.classList.remove("active");
  }
}

function setupCellListeners() {
  const table = document.getElementById("preferencesTable");
  console.log(table);
  const headers = [...table.querySelectorAll("thead th")];

  table.querySelectorAll("tbody tr").forEach((row, rowIndex) => {
    row.querySelectorAll("td").forEach((cell, colIndex) => {
      if (cell.getAttribute("contenteditable") === "true") {
        let oldValue = cell.innerText.trim();

        cell.addEventListener("blur", () => {
          const newValue = cell.innerText.trim();
          const columnName = headers[colIndex].innerText;
          trackEdit(rowIndex, columnName, oldValue, newValue);
          oldValue = newValue; // update after blur
        });
      }
    });
  });
}

// Confirmation Modal
function showModal() {
  const modal = document.getElementById("confirmationModal");
  const summary = document.getElementById("changeSummary");
  let content = [];

// Preferences changes
if (editedRows.length > 0) {
  const preferenceEdits = editedRows.map(edit => {
    const pref = preferencesData[edit.rowIndex];
    const label = pref ? `(${pref.course} - ${pref.faculty})` : `Row ${edit.rowIndex + 1}`;
    return `<div>${label} — <strong>${edit.column}</strong>: "<span style="color: red">${edit.oldValue}</span>" → "<span style="color: green">${edit.newValue}</span>"</div>`;
  });
  content.push(...preferenceEdits);
}

// Faculty changes
if (facultyChanges.length > 0) {
  const facultyEdits = facultyChanges.map(change => {
    if (change.type === "edit") {
      return `<div><strong>Faculty</strong> "${change.oldName}" → "<span style="color: green">${change.newName}</span>"</div>`;
    } else if (change.type === "remove") {
      return `<div style="color: red">Deleted faculty: <strong>${change.name}</strong></div>`;
    } else if (change.type === "add") {
      return `<div style="color: green">Added faculty: <strong>${change.name}</strong></div>`;
    }
  });
  content.push(...facultyEdits);
}

// Course changes
if (courseChanges.length > 0) {
  const courseEdits = courseChanges.map(change => {
    if (change.type === "edit") {
      const nameChange = change.oldName !== change.newName 
        ? `Name: "<span style="color: red">${change.oldName}</span>" → "<span style="color: green">${change.newName}</span>"`
        : '';
      const sectionChange = change.oldSection !== change.newSection 
        ? `Sections: "<span style="color: red">${change.oldSection}</span>" → "<span style="color: green">${change.newSection}</span>"`
        : '';
      return `<div><strong>Course</strong> ${nameChange} ${nameChange && sectionChange ? '<br>' : ''} ${sectionChange}</div>`;
    } else if (change.type === "remove") {
      return `<div style="color: red">Deleted course: <strong>${change.name}</strong></div>`;
    } else if (change.type === "add") {
      return `<div style="color: green">Added course: <strong>${change.name}</strong> (Sections: ${change.section})</div>`;
    }
  });
  content.push(...courseEdits);
}

// Default message
summary.innerHTML = content.length > 0 ? content.join('') : "<em>No changes found.</em>";
modal.classList.remove("hidden");

  // if (editedRows.length === 0) {
  //   summary.innerHTML = "<em>No changes found.</em>";
  // } else {
  //   summary.innerHTML = editedRows.map(edit => {
  //     const pref = preferencesData[edit.rowIndex];
  //     const label = pref ? `(${pref.course} - ${pref.faculty})` : `Row ${edit.rowIndex + 1}`;
  //     return `<div>${label} — <strong>${edit.column}</strong>: "<span style="color: red">${edit.oldValue}</span>" → "<span style="color: green">${edit.newValue}</span>"</div>`;
  //         }).join("");
  // }

  // modal.classList.add("active");
}

// pass the three changelog arrays to the database
// resets the changelogs if successful
async function pushBatchChanges() {
  const changes = {
    facultyChanges: facultyChanges,
    courseChanges: courseChanges,
    preferenceChanges: preferenceChanges
  };

  // Call the batch handler API
  const response = await fetch('/api/batch', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(changes)
  });

  const result = await response.json();
  
  if (response.ok) {
    // If the batch update is successful, show success and reset state
    alert("Changes confirmed!");

    // Hide the confirmation modal
    const modal = document.getElementById("confirmationModal");
    modal.classList.remove("active");
    modal.classList.add("hidden");

    // Reset the changelog and button states
    editedRows = [];
    facultyChanges = [];
    courseChanges = [];
    preferenceChanges = [];

    updatePushButtonState(); // Reset the push button state
    updateFacultyPushState(); // Reset faculty push button state
    updateCoursePushState();  // Reset course push button state
    updatePreferencePushState(); // Reset preference push button
  } else {
    // Handle failure if there was an error
    alert("Error: " + result.error);
  }
}

// Attach the function to the confirm button in the modal
document.getElementById("confirmBtn").addEventListener("click", function () {
  confirmBatchChanges();  // Call the function to process the changes
});



document.addEventListener("DOMContentLoaded", () => {
  setupCellListeners();

  document.getElementById("pushChangesBtn").addEventListener("click", function () {
    // alert("Button clicked");
    // showModal();
    const modal = document.getElementById("confirmationModal");
    const summary = document.getElementById("changeSummary");
    if (editedRows.length === 0) {
      summary.innerHTML = "<em>No changes found.</em>";
    } else {
      summary.innerHTML = editedRows.map(edit => {
        const pref = preferencesData[edit.rowIndex];
        const label = pref ? `(${pref.course} - ${pref.faculty})` : `Row ${edit.rowIndex + 1}`;
        return `<div>${label} — <strong>${edit.column}</strong>: "<span style="color: red">${edit.oldValue}</span>" → "<span style="color: green">${edit.newValue}</span>"</div>`;
      }).join("");
    }
    
    modal.classList.remove("hidden");
    modal.classList.add("active");
  });
  
  // Push changes
  document.getElementById("confirmBtn").addEventListener("click", function () {
    pushBatchChanges();
  });
  
  document.getElementById("cancelBtn").addEventListener("click", function () {
    const modal = document.getElementById("confirmationModal");
    modal.classList.remove("active");
    modal.classList.add("hidden");
  });    
});

  document.getElementById("pushFacultyChangesBtn").addEventListener("click", function () {
    const summary = document.getElementById("changeSummary");
    if (facultyChanges.length === 0) {
      summary.innerHTML = "<em>No faculty changes found.</em>";
    } else {
      summary.innerHTML = facultyChanges.map(change => {
        if (change.type === 'add') {
          return `<div>Added Faculty: <strong>${change.name}</strong></div>`;
        } else if (change.type === 'edit') {
          return `<div>Edited Faculty: <span style="color: red">${change.oldName}</span> → <span style="color: green">${change.newName}</span></div>`;
        }
        else if (change.type === 'remove') {
          return `<div style="color: red">Deleted Faculty: <strong>${change.name}</strong></div>`;
        }
      }).join('');
    }
    document.getElementById("confirmationModal").classList.remove("hidden");
    document.getElementById("confirmationModal").classList.add("active");
  });
  
  document.getElementById("pushCourseChangesBtn").addEventListener("click", function () {
    const summary = document.getElementById("changeSummary");
    if (courseChanges.length === 0) {
      summary.innerHTML = "<em>No course changes found.</em>";
    } else {
      summary.innerHTML = courseChanges.map(change => {
        if (change.type === 'add') {
          return `<div>Added Course: <strong>${change.name}</strong></div>`;
        } else if (change.type === 'edit') {
          return `<div>Edited Course: <span style="color: red">${change.oldName}</span> → <span style="color: green">${change.newName}</span> (Section: ${change.oldSection} → ${change.newSection})</div>`;
        } else if (change.type === 'remove') {
          return `<div style="color: red">Deleted Course: <strong>${change.name}</strong></div>`;
        }
      }).join('');
    }
    document.getElementById("confirmationModal").classList.remove("hidden");
    document.getElementById("confirmationModal").classList.add("active");
  });
  
  document.addEventListener('click', function(event) {
    const prefSidebar = document.getElementById('editPreferenceSidebar');
    const courseSidebar = document.getElementById('editSidebar');
    const facultySidebar = document.getElementById('editFacultySidebar');
  
    const isButton = event.target.closest('button');
  
    // Only close if sidebar is open, click is outside the sidebar, and not on a button
    if (!isButton) {
      if (prefSidebar.classList.contains('active') && !prefSidebar.contains(event.target)) {
        closePreferenceSidebar();
      }
  
      if (courseSidebar.classList.contains('active') && !courseSidebar.contains(event.target)) {
        closeSidebar();
      }
  
      if (facultySidebar.classList.contains('active') && !facultySidebar.contains(event.target)) {
        closeFacultySidebar();
      }
    }
  });
  

//submit faculty form
function submitForm() {
  const facultyName = document.getElementById('facultyName').value;
  const facultyEmail = document.getElementById('facultyEmail').value;
  const courseNames = document.querySelectorAll('.courseName');
  const coursePreferences = document.querySelectorAll('.coursePreference');

  let allValid = true;

  if (!facultyName || !facultyEmail) {
    alert('Please fill in the Faculty Name and Email!');
    allValid = false;
  }

  // Check if all course fields are filled
  courseNames.forEach((course, index) => {
    if (!course.value || !coursePreferences[index].value) {
      allValid = false;
      alert('Please fill in all course details!');
    }
  });

  if (allValid) {
    document.getElementById('confirmation').style.display = 'block';
  }
}


function addFormCourse() {
  const coursesContainer = document.getElementById('coursesContainer');
  
  // Create new course input fields
  const newCourseGroup = document.createElement('div');
  newCourseGroup.classList.add('course-group');
  
  const courseNameLabel = document.createElement('label');
  courseNameLabel.textContent = 'Course Name:';
  const courseNameInput = document.createElement('input');
  courseNameInput.type = 'text';
  courseNameInput.classList.add('courseName');
  courseNameInput.placeholder = 'Enter Course Name';
  
  const preferenceLabel = document.createElement('label');
  preferenceLabel.textContent = 'Course Preference Level:';
  const preferenceSelect = document.createElement('select');
  preferenceSelect.classList.add('coursePreference');
  preferenceSelect.innerHTML = `
    <option value="green">Green</option>
    <option value="yellow">Yellow</option>
    <option value="red">Red</option>
  `;
  
  // Append new elements to the new course group
  newCourseGroup.appendChild(courseNameLabel);
  newCourseGroup.appendChild(courseNameInput);
  newCourseGroup.appendChild(preferenceLabel);
  newCourseGroup.appendChild(preferenceSelect);
  
  // Append new course group to courses container
  coursesContainer.appendChild(newCourseGroup);
}