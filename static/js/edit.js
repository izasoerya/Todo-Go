function inputHandler() {
    fetch('/app/edit', {
		method: 'PUT',
		headers: {
		  'Content-Type': 'application/json'
		},
	})
	.then(response => response.json())
	.then(jsonData => {
		// Handle the JSON data here
		console.log(jsonData);
		alert(jsonData);
	})
	.catch(error => {
		// Handle any errors that occurred during the request
		console.error('Error:', error);
	});
	  
	
	const completeEdit      = document.getElementById('completeEdit');		//Get input from input HTML
	const editCompleteValue = completeEdit.value;	
    const titleEdit         = document.getElementById('titleEdit');		//Get input from input HTML
	const editTitleValue    = titleEdit.value;		
    
    const editData = {
		title: editTitleValue,			    //* Make json {"title":titleValue}
        completed : editCompleteValue,      //* Make json {"completed":completedValue}
    }; 
    
}


