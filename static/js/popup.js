function handleKeyDown(event) {
  if (event.keyCode === 13 || event.key === 'Enter') {
    createTodo();
  }
}

function createTodo() {
  fetch('localhost:3000/app', {
    method: 'POST',
  })
  .then(response => response.json())
  .then(data => {
    console.log(data)
  })
  new Notification("Todo Created!")
  .catch(error => {
    console.error('Error:', error);
  });
}



  
