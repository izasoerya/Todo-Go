function ShowNotification() {
    if (Notification.permission === "granted") {
        new Notification("Todo Created!");
    } else if (Notification.permission !== "denied") {
        Notification.requestPermission().then(function(permission) {
            if (permission === "granted") {
                new Notification("Todo Created!");
            } 
        });
    }
}

function handleKeyDown(event) {
  if (event.keyCode === 13 || event.key === 'Enter') {
    createTodo();
  }
}

function createTodo() {
  fetch('/app', {
    method: 'POST',
  })
  ShowNotification()
    .then(response => {
      if (!response.ok) {
        throw new Error('Error:', response.status);
      }
    })
    .catch(error => {
      console.error('Error:', error);
      // Handle the error here
    });
}



  
