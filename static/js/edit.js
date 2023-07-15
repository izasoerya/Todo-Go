function editTodoPage() {
	const url = window.location.href;					//* parse url for get id
	const urlParse = /\/app\/edit\/(.+)/;
	const match = url.match(urlParse);
	if (match) {
		idEdit = match[1].toString();
	}

    const titleEdit = document.getElementById('titleEdit');
    const editTitleValue = titleEdit.value;

    const completeEdit = document.getElementById('completeEdit');
    const editCompleteValue = completeEdit.checked;

    const editData = {
        title: editTitleValue,
        completed: editCompleteValue
    };

	fetch('/api/Todos/' + idEdit, {
		method: 'PUT',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(editData)		//Convert JavaScript object to JSON string
		})
	.then(Response => {
		if(Response.ok) {
			window.location.href = '/app';
		}
		else {
			console.error('Error:', Response);
		}
	})
	.catch(error => {
		console.error('Error:', error);
	});
}
