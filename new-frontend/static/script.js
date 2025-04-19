let editRow = null;
let editMode = ''; 
let facultyChanges = [];
let courseChanges = [];

// Add Faculty
function addFaculty() {
  const name = document.getElementById('facultyName').value.trim();
  if (!name) return;
  if (name) {
    const table = document.getElementById('facultyTableBody');
    
    const existing = [...table.rows].some(row => row.cells[0].textContent.trim().toLowerCase() === name.toLowerCase());
    if (existing) {
      alert("Faculty already exists!"); 
      return;
    }
    
    const row = document.createElement('tr');
    row.innerHTML = `
      <td>${name}</td>
      <td>
        <button onclick="editFaculty(this)">Edit</button>
        <button onclick="removeFaculty(this)">Remove</button>
      </td>
    `;
    table.appendChild(row);
    document.getElementById('facultyName').value = '';
    facultyChanges.push({ type: "add", name });
    updateFacultyPushState();
  }
}


function removeFaculty(btn) {
  const row = btn.closest('tr');
  row.remove();
  const name = row.children[0].textContent;
  facultyChanges.push({ type: "remove", name });
  updateFacultyPushState();
}

// Add Course
function addCourse() {
  const name = document.getElementById('courseName').value.trim();
  const section = document.getElementById('sectionCount').value;
  if (!name) return;

  if (name) {
    const table = document.getElementById('courseTableBody');
    
    const existing = [...table.rows].some(row => {
    const existingName = row.cells[0].textContent.trim().toLowerCase();
    const existingSection = row.cells[1].textContent.trim();
    return existingName === name.toLowerCase() && existingSection === section;
    });

    if (existing) {
      alert("Course already exists!");
      return;
    }
    
    const row = document.createElement('tr');
    row.innerHTML = `
      <td>${name}</td>
      <td>${section}</td>
      <td>
        <button onclick="editCourse(this)">Edit</button>
        <button onclick="removeCourse(this)">Remove</button>
      </td>
    `;
    table.appendChild(row);
    document.getElementById('courseName').value = '';
    document.getElementById('sectionCount').value = '1';
    courseChanges.push({ type: "add", name, section });
    updateCoursePushState();
  }
}

function removeCourse(btn) {
  const row = btn.closest('tr');
  row.remove();
  const name = row.children[0].textContent;
  courseChanges.push({ type: "remove", name });
  updateCoursePushState();
}


// Edit Faculty Entry
function editFaculty(btn) {
  editRow = btn.closest('tr');
  editMode = 'faculty';
  document.getElementById("editfacultyName").value = editRow.children[0].textContent;
  openFacultySidebar();
}

// Edit Course Entry
function editCourse(btn) {
  editRow = btn.closest('tr');
  editMode = 'course';

  // document.getElementById("editType").value = "Course";
  document.getElementById("editName").value = editRow.children[0].textContent;

  const section = editRow.children[1].textContent.trim();;
  document.getElementById("editSection").value = section;
  // document.getElementById("editSectionGroup").classList.remove("hidden");

  openSidebar();
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
  
  // function openSidebar(sidebarId) {
  //   document.getElementById(sidebarId).style.display = "block";
  // }
  
  function closeFacultySidebar() {
    document.getElementById("editFacultySidebar").classList.remove("active");
    document.body.classList.remove("sidebar-open");
    document.getElementById("editfacultyName").value = '';
    document.getElementById("facultyName").value = '';
  }
  
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
  
  
// Handle Sidebar Form Submission
document.getElementById("editForm").addEventListener("submit", function (e) {
  e.preventDefault();
  const newName = document.getElementById("editName").value.trim();
  const oldName = editRow.children[0].textContent.trim();

  if (!newName) return;

  editRow.children[0].textContent = newName;

  if (editMode === 'faculty') {
    if (newName !== oldName) {
      facultyChanges.push({ type: "edit", oldName, newName });
      updateFacultyPushState();
    }
  } else if (editMode === 'course') {
    const newSection = document.getElementById("editSection").value;
    const sectionCell = editRow.children[1];
    const oldSection = sectionCell.textContent.trim();
    sectionCell.textContent = newSection;


    courseChanges.push({ type: "edit", oldName, newName, oldSection, newSection });
    updateCoursePushState();
  }

  closeSidebar();
});

let editPreferenceRow = null;

// to be fetched from db
const preferencesData = [
  { course: 'CS101', faculty: 'Diazh', preference: 'green' },
  { course: 'CS102', faculty: 'Anson', preference: 'yellow' },
];

function loadPreferences() {
    const tableBody = document.getElementById('preferencesTableBody');
    tableBody.innerHTML = '';  // Clear any existing rows
    preferencesData.forEach((pref, index) => {
      const row = document.createElement('tr');
      const isEdited = editedRows.some(r => r.rowIndex === index);

      row.innerHTML = `
        <td>${pref.course}</td>
        <td>${pref.faculty}</td>
        <td><span class="preference ${pref.preference}">${pref.preference}</span></td>
        <td>
          <button onclick="editPreference(${index})">Edit</button>
          <button onclick="removePreference(${index})" style="display: none;">Remove</button>
        </td>
      `;
      if (isEdited) {
        row.style.backgroundColor = '#f0f8ff'; 
      }
      tableBody.appendChild(row);
    });
  }
  
  // Load preferences on page load
  document.addEventListener('DOMContentLoaded', loadPreferences);
  
  // Open sidebar to edit preference
  function editPreference(index) {
    editPreferenceRow = index;
    const preference = preferencesData[index];
  
    document.getElementById('editCourse').value = preference.course;
    document.getElementById('editFaculty').value = preference.faculty;
    document.getElementById('editPreference').value = preference.preference;
  
    openPreferenceSidebar();
  }
  
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
  
  // Save edited preference
  document.getElementById('editPreferenceForm').addEventListener('submit', function (e) {
    e.preventDefault();
    
    const updatedPreference = document.getElementById('editPreference').value;
    const currentpref = preferencesData[editPreferenceRow];
    const oldpref = currentpref.preference;
    preferencesData[editPreferenceRow].preference = updatedPreference;
    
    if(updatedPreference !== oldpref) {
      trackEdit(editPreferenceRow, 'Preference', oldpref, updatedPreference);
      console.log("Tracking edit:", { rowIndex: editPreferenceRow, column: 'Preference', oldValue: oldpref, newValue: updatedPreference });

    }
    loadPreferences();
    closePreferenceSidebar();
  });

  document.getElementById("editFacultyForm").addEventListener("submit", function (e) {
    e.preventDefault(); // Prevents page reload
  
    const newName = document.getElementById("editfacultyName").value.trim();
    if (!newName || !editRow) return;
  
    const oldName = editRow.children[0].textContent.trim();
    editRow.children[0].textContent = newName;
  
    if (newName !== oldName) {
      facultyChanges.push({ type: "edit", oldName, newName });
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
    
    // document.getElementById("confirmBtn").addEventListener("click", function () {
    //   alert("Changes confirmed!");
    //   document.getElementById("confirmationModal").classList.add("hidden");
    //   const pushBtn = document.getElementById("pushChangesBtn");
    //   pushBtn.classList.remove("active");
    //   pushBtn.disabled = true;
    document.getElementById("confirmBtn").addEventListener("click", function () {
      alert("Changes confirmed!");
      const modal = document.getElementById("confirmationModal");
      modal.classList.remove("active");
      modal.classList.add("hidden");
      
      editedRows = [];
      facultyChanges = [];
      courseChanges = [];
      updatePushButtonState();
      updateFacultyPushState();
      updateCoursePushState();

      document.getElementById("confirmationModal").classList.remove("active");
      document.getElementById("confirmationModal").classList.add("hidden");
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
  