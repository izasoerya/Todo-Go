document.addEventListener('DOMContentLoaded', function() {
  if (window.location.href.includes('localhost:3000/app/show')) {
    // Call your function here
    fetchDB();
  }
});

function fetchDB() {
    fetch('/api/todos')
      .then(response => {
        if (response.ok) {
          return response.json();
        } else {
          throw new Error('Request failed with status:', response.status);
        }
      })
      .then(data => {
        const todos = data.data.todo; // Access the 'todo' array from the response
				tableHandler(todos);
      })
      .catch(error => {
        console.error('Request failed:', error);
      });
}
  

function tableHandler(data) {
	const table = document.createElement('table');

  // Create table header row
  const headerRow = document.createElement('tr');
  const headers = Object.keys(data[0]);
  headers.forEach(header => {
    const cell = document.createElement('th');
    cell.textContent = header;
    headerRow.appendChild(cell);
  });
  table.appendChild(headerRow);

  // Create table rows
  data.forEach(item => {
    const row = document.createElement('tr');
    Object.values(item).forEach(value => {
      const cell = document.createElement('td');
      cell.textContent = value;
      row.appendChild(cell);
    });
    table.appendChild(row);
  });

  // Add the table to the document body
  document.body.appendChild(table);
}

function buttonBack() {
	window.location.assign("http://localhost:3000/app");
}



