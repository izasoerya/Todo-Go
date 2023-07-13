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
	//DO SOMETHING
	}
}

function submitForm(iterHandler) {
	const titleInput = document.getElementById('inputTodos');		//Get input from input HTML
	const titleValue = titleInput.value;		
	
	const createData = {
		title: titleValue		//Make json {"title":titleValue}
	};
	const deleteData = 2
	const showData = 2
	const editData = 2
	
	switch(iterHandler) {
		case "Create":
			createHandler(createData)
			break;
		case "Delete":
			createHandler(deleteData)
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
	fetch('/api/Todos/CreateTodoPage', {
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


  
