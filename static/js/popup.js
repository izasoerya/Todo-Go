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
	const titleInput = document.getElementById('inputTodos');		//Get input from input HTML
	const inputValue = titleInput.value;		
	
	const createData = {
		title: inputValue			//* Make json {"title":titleValue}
	};
	const deleteData = inputValue
	const showData = 2
	const editData = 2
	
	switch(iterHandler) {
		case "Create":
			createHandler(createData)
			break;
		case "Delete":
			deleteHandler(deleteData)
			break;
		case "Edit":
			createHandler(editData)
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
}

function editHandler(editData) {
	fetch('/api/Todos/' + editData,  {
		method: 'PUT'
	})
}

function showHandler(deleteData) {
	window.location.replace("http://localhost:3000/api/Todos/" + deleteData);
}

  
