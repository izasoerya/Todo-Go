function buttonCreate() {
	submitForm("Create")
}
function buttonDelete() {
	submitForm("Delete")
}
function buttonEdit() {
	submitForm("Edit")
}
function buttonShow() {
	submitForm("Show")
}

function handleKeyDown(event) {
	if (event.keyCode === 13 || event.key === 'Enter') {
		//! Do something special
	}
}

function submitForm(iterHandler) {
	const titleInput = document.getElementById('inputTodos');		//* GET input
	const inputValue = titleInput.value;		

	const createData = {											//* Create data				
		title: inputValue			
	};
	const deleteData = inputValue									//* Delete data
	const editData   = inputValue									//* Edit data
	const showData	 = 2											//* Show data
	
	switch(iterHandler) {
		case "Create":
			createHandler(createData)
			break;
		case "Delete":
			deleteHandler(deleteData)
			break;
		case "Edit":
			editHandler(editData)
			break;
		case "Show":
			createHandler(showData)
			break;
	}
}

function createHandler(createData) {
	fetch('/api/Todos/', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(createData)		//Convert JavaScript object to JSON string
		})
	.catch(error => {
		console.error('Error:', error);
	});
}

function deleteHandler(deleteData) {
	fetch('/api/Todos/' + deleteData,  {
		method: 'DELETE'
	})
	.catch(error => {
		console.error('Error:', error);
	});
}

function editHandler(editData) {
	fetch('/app/edit',  {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(editData)		
	})
	.catch(error => {
		console.error('Error:', error);
	})
	window.location.replace("http://localhost:3000/app/edit/")
}

function showHandler(deleteData) {
	window.location.replace("http://localhost:3000/api/Todos/" + deleteData);
}

  
