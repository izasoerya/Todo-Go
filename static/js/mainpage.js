function buttonCreate() {submitForm("Create")}
function buttonDelete() {submitForm("Delete")}
function buttonEdit() 	{submitForm("Edit")	 }
function buttonShow() 	{submitForm("Show")	 }
function buttonBack() 	{submitForm("Back")  }

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
	const rawEdit 	 = inputValue									//* Edit data	
	const editData	 = rawEdit.toString()							//* Edit data
	
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
			showHandler()
			break;
		case "Back":
			window.location.assign("http://localhost:3000/app");
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
	window.location.assign("http://localhost:3000/app/edit/" + editData)
}

function showHandler() {
	window.location.assign("http://localhost:3000/app/show");
}

  
